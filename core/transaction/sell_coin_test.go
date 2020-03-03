package transaction

import (
	"github.com/MinterTeam/minter-go-node/core/types"
	"github.com/MinterTeam/minter-go-node/crypto"
	"github.com/MinterTeam/minter-go-node/helpers"
	"github.com/MinterTeam/minter-go-node/rlp"
	"math/big"
	"sync"
	"testing"
)

func TestSellCoinTx(t *testing.T) {
	cState := getState()

	createTestCoin(cState)

	privateKey, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(privateKey.PublicKey)
	coin := types.GetBaseCoin()

	cState.Accounts.AddBalance(addr, coin, helpers.BipToPip(big.NewInt(1000000)))

	minValToBuy, _ := big.NewInt(0).SetString("957658277688702625", 10)
	data := SellCoinData{
		CoinToSell:        coin,
		ValueToSell:       helpers.BipToPip(big.NewInt(10)),
		CoinToBuy:         getTestCoinSymbol(),
		MinimumValueToBuy: minValToBuy,
	}

	encodedData, err := rlp.EncodeToBytes(data)

	if err != nil {
		t.Fatal(err)
	}

	tx := Transaction{
		Nonce:         1,
		GasPrice:      1,
		ChainID:       types.CurrentChainID,
		GasCoin:       coin,
		Type:          TypeSellCoin,
		Data:          encodedData,
		SignatureType: SigTypeSingle,
	}

	if err := tx.Sign(privateKey); err != nil {
		t.Fatal(err)
	}

	encodedTx, err := rlp.EncodeToBytes(tx)

	if err != nil {
		t.Fatal(err)
	}

	response := RunTx(cState, false, encodedTx, big.NewInt(0), 0, &sync.Map{}, 0)

	if response.Code != 0 {
		t.Fatalf("Response code is not 0. Error: %s", response.Log)
	}

	targetBalance, _ := big.NewInt(0).SetString("999989900000000000000000", 10)
	balance := cState.Accounts.GetBalance(addr, coin)
	if balance.Cmp(targetBalance) != 0 {
		t.Fatalf("Target %s balance is not correct. Expected %s, got %s", coin, targetBalance, balance)
	}

	targetTestBalance, _ := big.NewInt(0).SetString("999955002849793446", 10)
	testBalance := cState.Accounts.GetBalance(addr, getTestCoinSymbol())
	if testBalance.Cmp(targetTestBalance) != 0 {
		t.Fatalf("Target %s balance is not correct. Expected %s, got %s", getTestCoinSymbol(), targetTestBalance, testBalance)
	}
}

func TestSellCoinTxBaseToCustomBaseCommission(t *testing.T) {
	// sell_coin: MNT
	// buy_coin: TEST
	// gas_coin: MNT
}

func TestSellCoinTxCustomToBaseBaseCommission(t *testing.T) {
	// sell_coin: TEST
	// buy_coin: MNT
	// gas_coin: MNT
}

func TestSellCoinTxCustomToCustomBaseCommission(t *testing.T) {
	// sell_coin: TEST1
	// buy_coin: TEST2
	// gas_coin: MNT
}

func TestSellCoinTxBaseToCustomCustomCommission(t *testing.T) {
	// sell_coin: MNT
	// buy_coin: TEST
	// gas_coin: TEST
}

func TestSellCoinTxCustomToBaseCustomCommission(t *testing.T) {
	// sell_coin: TEST
	// buy_coin: MNT
	// gas_coin: TEST
}

func TestSellCoinTxCustomToCustomCustom1Commission(t *testing.T) {
	// sell_coin: TEST1
	// buy_coin: TEST2
	// gas_coin: TEST1
}

func TestSellCoinTxCustomToCustomCustom2Commission(t *testing.T) {
	// sell_coin: TEST1
	// buy_coin: TEST2
	// gas_coin: TEST2
}
