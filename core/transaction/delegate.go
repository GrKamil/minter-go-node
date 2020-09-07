package transaction

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"

	"github.com/MinterTeam/minter-go-node/core/code"
	"github.com/MinterTeam/minter-go-node/core/commissions"
	"github.com/MinterTeam/minter-go-node/core/state"
	"github.com/MinterTeam/minter-go-node/core/types"
	"github.com/MinterTeam/minter-go-node/formula"
	"github.com/MinterTeam/minter-go-node/hexutil"
	"github.com/tendermint/tendermint/libs/kv"
)

type DelegateData struct {
	PubKey types.Pubkey
	Coin   types.CoinID
	Value  *big.Int
}

func (data DelegateData) BasicCheck(tx *Transaction, context *state.CheckState) *Response {
	if data.Value == nil {
		return &Response{
			Code: code.DecodeError,
			Log:  "Incorrect tx data",
			Info: EncodeError(map[string]string{
				"code": strconv.Itoa(int(code.DecodeError)),
			})}
	}

	if !context.Coins().Exists(tx.GasCoin) {
		return &Response{
			Code: code.CoinNotExists,
			Log:  fmt.Sprintf("Coin %s not exists", tx.GasCoin),
			Info: EncodeError(map[string]string{
				"code":    strconv.Itoa(int(code.CoinNotExists)),
				"coin_id": fmt.Sprintf("%s", tx.GasCoin.String()),
			})}
	}

	if !context.Coins().Exists(data.Coin) {
		return &Response{
			Code: code.CoinNotExists,
			Log:  fmt.Sprintf("Coin %s not exists", data.Coin),
			Info: EncodeError(map[string]string{
				"code":    strconv.Itoa(int(code.CoinNotExists)),
				"coin_id": fmt.Sprintf("%s", data.Coin.String()),
			})}
	}

	if data.Value.Cmp(types.Big0) < 1 {
		return &Response{
			Code: code.StakeShouldBePositive,
			Log:  fmt.Sprintf("Stake should be positive"),
			Info: EncodeError(map[string]string{
				"code": strconv.Itoa(int(code.StakeShouldBePositive)),
			})}
	}

	if !context.Candidates().Exists(data.PubKey) {
		return &Response{
			Code: code.CandidateNotFound,
			Log:  fmt.Sprintf("Candidate with such public key not found"),
			Info: EncodeError(map[string]string{
				"code":    strconv.Itoa(int(code.CandidateNotFound)),
				"pub_key": data.PubKey.String(),
			}),
		}
	}

	sender, _ := tx.Sender()
	if !context.Candidates().IsDelegatorStakeSufficient(sender, data.PubKey, data.Coin, data.Value) {
		return &Response{
			Code: code.TooLowStake,
			Log:  fmt.Sprintf("Stake is too low"),
			Info: EncodeError(map[string]string{
				"code":    strconv.Itoa(int(code.TooLowStake)),
				"pub_key": data.PubKey.String(),
			})}
	}

	return nil
}

func (data DelegateData) String() string {
	return fmt.Sprintf("DELEGATE pubkey:%s ",
		hexutil.Encode(data.PubKey[:]))
}

func (data DelegateData) Gas() int64 {
	return commissions.DelegateTx
}

func (data DelegateData) Run(tx *Transaction, context state.Interface, rewardPool *big.Int, currentBlock uint64) Response {
	sender, _ := tx.Sender()

	var checkState *state.CheckState
	var isCheck bool
	if checkState, isCheck = context.(*state.CheckState); !isCheck {
		checkState = state.NewCheckState(context.(*state.State))
	}

	response := data.BasicCheck(tx, checkState)
	if response != nil {
		return *response
	}

	commissionInBaseCoin := tx.CommissionInBaseCoin()
	commission := big.NewInt(0).Set(commissionInBaseCoin)

	gasCoin := checkState.Coins().GetCoin(tx.GasCoin)
	coin := checkState.Coins().GetCoin(data.Coin)

	if !tx.GasCoin.IsBaseCoin() {
		errResp := CheckReserveUnderflow(gasCoin, commissionInBaseCoin)
		if errResp != nil {
			return *errResp
		}

		commission = formula.CalculateSaleAmount(gasCoin.Volume(), gasCoin.Reserve(), gasCoin.Crr(), commissionInBaseCoin)
	}

	if checkState.Accounts().GetBalance(sender, tx.GasCoin).Cmp(commission) < 0 {
		return Response{
			Code: code.InsufficientFunds,
			Log:  fmt.Sprintf("Insufficient funds for sender account: %s. Wanted %s %s", sender.String(), commission, gasCoin.GetFullSymbol()),
			Info: EncodeError(map[string]string{
				"code":         strconv.Itoa(int(code.InsufficientFunds)),
				"sender":       sender.String(),
				"needed_value": commission.String(),
				"coin_symbol":  gasCoin.GetFullSymbol(),
			}),
		}
	}

	if checkState.Accounts().GetBalance(sender, data.Coin).Cmp(data.Value) < 0 {
		return Response{
			Code: code.InsufficientFunds,
			Log:  fmt.Sprintf("Insufficient funds for sender account: %s. Wanted %s %s", sender.String(), data.Value, coin.GetFullSymbol()),
			Info: EncodeError(map[string]string{
				"code":         strconv.Itoa(int(code.InsufficientFunds)),
				"sender":       sender.String(),
				"needed_value": data.Value.String(),
				"coin_symbol":  coin.GetFullSymbol(),
			}),
		}
	}

	if data.Coin == tx.GasCoin {
		totalTxCost := big.NewInt(0)
		totalTxCost.Add(totalTxCost, data.Value)
		totalTxCost.Add(totalTxCost, commission)

		if checkState.Accounts().GetBalance(sender, tx.GasCoin).Cmp(totalTxCost) < 0 {
			return Response{
				Code: code.InsufficientFunds,
				Log:  fmt.Sprintf("Insufficient funds for sender account: %s. Wanted %s %s", sender.String(), totalTxCost.String(), gasCoin.GetFullSymbol()),
				Info: EncodeError(map[string]string{
					"code":         strconv.Itoa(int(code.InsufficientFunds)),
					"sender":       sender.String(),
					"needed_value": totalTxCost.String(),
					"coin_symbol":  gasCoin.GetFullSymbol(),
				}),
			}
		}
	}

	if deliverState, ok := context.(*state.State); ok {
		rewardPool.Add(rewardPool, commissionInBaseCoin)

		deliverState.Coins.SubReserve(tx.GasCoin, commissionInBaseCoin)
		deliverState.Coins.SubVolume(tx.GasCoin, commission)

		deliverState.Accounts.SubBalance(sender, tx.GasCoin, commission)
		deliverState.Accounts.SubBalance(sender, data.Coin, data.Value)

		value := big.NewInt(0).Set(data.Value)
		if watchList := deliverState.Waitlist.Get(sender, data.PubKey, data.Coin); watchList != nil {
			value.Add(value, watchList.Value)
			deliverState.Waitlist.Delete(sender, data.PubKey, data.Coin)
		}

		deliverState.Candidates.Delegate(sender, data.PubKey, data.Coin, value, big.NewInt(0))
		deliverState.Accounts.SetNonce(sender, tx.Nonce)
	}

	tags := kv.Pairs{
		kv.Pair{Key: []byte("tx.type"), Value: []byte(hex.EncodeToString([]byte{byte(TypeDelegate)}))},
		kv.Pair{Key: []byte("tx.from"), Value: []byte(hex.EncodeToString(sender[:]))},
	}

	return Response{
		Code:      code.OK,
		GasUsed:   tx.Gas(),
		GasWanted: tx.Gas(),
		Tags:      tags,
	}
}
