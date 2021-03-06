package transaction

//
// import (
// 	"encoding/hex"
// 	"fmt"
// 	"github.com/MinterTeam/minter-go-node/coreV2/code"
// 	"github.com/MinterTeam/minter-go-node/coreV2/state"
// 	"github.com/MinterTeam/minter-go-node/coreV2/state/commission"
// 	"github.com/MinterTeam/minter-go-node/coreV2/types"
// 	abcTypes "github.com/tendermint/tendermint/abci/types"
// 	"math/big"
// )
//
// type MoveStakeData struct {
// 	From, To types.Pubkey
// 	Coin     types.CoinID
// 	Stake    *big.Int
// }
//
// func (data MoveStakeData) Gas() int64 {
// 	return gasMoveStake
// }
// func (data MoveStakeData) TxType() TxType {
// 	return TypeMoveStake
// }
//
// func (data MoveStakeData) basicCheck(tx *Transaction, context *state.CheckState) *Response {
// 	if !context.Coins().Exists(data.Coin) {
// 		return &Response{
// 			Code: code.CoinNotExists,
// 			Log:  fmt.Sprintf("Coin %s not exists", data.Coin),
// 			Info: EncodeError(code.NewCoinNotExists("", data.Coin.String())),
// 		}
// 	}
//
// 	if !context.Candidates().Exists(data.From) {
// 		return &Response{
// 			Code: code.CandidateNotFound,
// 			Log:  fmt.Sprintf("Candidate with %s public key not found", data.From),
// 			Info: EncodeError(code.NewCandidateNotFound(data.From.String())),
// 		}
// 	}
// 	if !context.Candidates().Exists(data.To) {
// 		return &Response{
// 			Code: code.CandidateNotFound,
// 			Log:  fmt.Sprintf("Candidate with %s public key not found", data.To),
// 			Info: EncodeError(code.NewCandidateNotFound(data.To.String())),
// 		}
// 	}
//
// 	sender, _ := tx.Sender()
//
// 	if waitlist := context.WaitList().Get(sender, data.From, data.Coin); waitlist != nil {
// 		if data.Stake.Cmp(waitlist.Value) == 1 {
// 			return &Response{
// 				Code: code.InsufficientWaitList,
// 				Log:  "Insufficient amount at waitlist for sender account",
// 				Info: EncodeError(code.NewInsufficientWaitList(waitlist.Value.String(), data.Stake.String())),
// 			}
// 		}
// 	} else {
// 		stake := context.Candidates().GetStakeValueOfAddress(data.From, sender, data.Coin)
//
// 		if stake == nil {
// 			return &Response{
// 				Code: code.StakeNotFound,
// 				Log:  "Stake of current user not found",
// 				Info: EncodeError(code.NewStakeNotFound(data.From.String(), sender.String(), data.Coin.String(), context.Coins().GetCoin(data.Coin).GetFullSymbol())),
// 			}
// 		}
//
// 		if stake.Cmp(data.Stake) == -1 {
// 			return &Response{
// 				Code: code.InsufficientStake,
// 				Log:  "Insufficient stake for sender account",
// 				Info: EncodeError(code.NewInsufficientStake(data.From.String(), sender.String(), data.Coin.String(), context.Coins().GetCoin(data.Coin).GetFullSymbol(), stake.String(), data.Stake.String())),
// 			}
// 		}
// 	}
//
// 	return nil
// }
//
// func (data MoveStakeData) String() string {
// 	return fmt.Sprintf("MOVE STAKE")
// }
//
// func (data MoveStakeData) CommissionData(price *commission.Price) *big.Int {
// 	return price.MoveStake
// }
//
// func (data MoveStakeData) Run(tx *Transaction, context state.Interface, rewardPool *big.Int, currentBlock uint64, price *big.Int) Response {
// 	sender, _ := tx.Sender()
//
// 	var checkState *state.CheckState
// 	var isCheck bool
// 	if checkState, isCheck = context.(*state.CheckState); !isCheck {
// 		checkState = state.NewCheckState(context.(*state.State))
// 	}
//
// 	response := data.basicCheck(tx, checkState)
// 	if response != nil {
// 		return *response
// 	}
//
// 	commissionInBaseCoin := tx.Commission(price)
// 	commissionPoolSwapper := checkState.Swap().GetSwapper(tx.GasCoin, types.GetBaseCoinID())
// 	gasCoin := checkState.Coins().GetCoin(tx.GasCoin)
// 	commission, isGasCommissionFromPoolSwap, errResp := CalculateCommission(checkState, commissionPoolSwapper, gasCoin, commissionInBaseCoin)
// 	if errResp != nil {
// 		return *errResp
// 	}
//
// 	if checkState.Accounts().GetBalance(sender, tx.GasCoin).Cmp(commission) < 0 {
// 		return Response{
// 			Code: code.InsufficientFunds,
// 			Log:  fmt.Sprintf("Insufficient funds for sender account: %s. Wanted %s %s", sender.String(), commission.String(), gasCoin.GetFullSymbol()),
// 			Info: EncodeError(code.NewInsufficientFunds(sender.String(), commission.String(), gasCoin.GetFullSymbol(), gasCoin.ID().String())),
// 		}
// 	}
// 	var tags []abcTypes.EventAttribute
// 	if deliverState, ok := context.(*state.State); ok {
// 		if isGasCommissionFromPoolSwap {
// 			commission, commissionInBaseCoin, _ = deliverState.Swap.PairSell(tx.GasCoin, types.GetBaseCoinID(), commission, commissionInBaseCoin)
// 		} else if !tx.GasCoin.IsBaseCoin() {
// 			deliverState.Coins.SubVolume(tx.GasCoin, commission)
// 			deliverState.Coins.SubReserve(tx.GasCoin, commissionInBaseCoin)
// 		}
// 		deliverState.Accounts.SubBalance(sender, tx.GasCoin, commission)
// 		rewardPool.Add(rewardPool, commissionInBaseCoin)
//
// 		if waitList := deliverState.Waitlist.Get(sender, data.From, data.Coin); waitList != nil {
// 			diffValue := big.NewInt(0).Sub(data.Stake, waitList.Value)
// 			deliverState.Waitlist.Delete(sender, data.From, data.Coin)
// 			if diffValue.Sign() == -1 {
// 				deliverState.Waitlist.AddWaitList(sender, data.From, data.Coin, big.NewInt(0).Neg(diffValue))
// 			}
// 		} else {
// 			deliverState.Candidates.SubStake(sender, data.From, data.Coin, data.Stake)
// 		}
//
// 		moveToCandidateId := deliverState.Candidates.ID(data.To)
// 		deliverState.FrozenFunds.AddFund(currentBlock+types.GetUnbondPeriod(), sender, data.From, deliverState.Candidates.ID(data.From), data.Coin, data.Stake, &moveToCandidateId)
//
// 		deliverState.Accounts.SetNonce(sender, tx.Nonce)
//
// 		tags = []abcTypes.EventAttribute{
// 			{Key: []byte("tx.commission_in_base_coin"), Value: []byte(commissionInBaseCoin.String())},
// 			{Key: []byte("tx.commission_conversion"), Value: []byte(isGasCommissionFromPoolSwap.String()), Index: true},
// 			{Key: []byte("tx.commission_amount"), Value: []byte(commission.String())},
// 			{Key: []byte("tx.from"), Value: []byte(hex.EncodeToString(sender[:]))},
//          {Key: []byte("tx.public_key_old"), Value: []byte(hex.EncodeToString(data.From[:])), Index: true},
//          {Key: []byte("tx.public_key_new"), Value: []byte(hex.EncodeToString(data.To[:])), Index: true},
// 		}
// 	}
//
// 	return Response{
// 		Code: code.OK,
// 		Tags: tags,
// 	}
// }
