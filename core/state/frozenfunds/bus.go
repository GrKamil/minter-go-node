package frozenfunds

import (
	"github.com/MinterTeam/minter-go-node/core/types"
	"math/big"
)

type Bus struct {
	frozenfunds *FrozenFunds
}

func (b *Bus) AddFrozenFund(height uint64, address types.Address, pubkey types.Pubkey, coin types.CoinID, value *big.Int) {
	b.frozenfunds.AddFund(height, address, pubkey, coin, value)
}

func NewBus(frozenfunds *FrozenFunds) *Bus {
	return &Bus{frozenfunds: frozenfunds}
}
