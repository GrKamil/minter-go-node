package accounts

import (
	"github.com/MinterTeam/minter-go-node/coreV2/types"
	"math/big"
)

type Bus struct {
	accounts *Accounts
}

func NewBus(accounts *Accounts) *Bus {
	return &Bus{accounts: accounts}
}

func (b *Bus) AddBalance(address types.Address, coin types.CoinID, value *big.Int) {
	b.accounts.AddBalance(address, coin, value)
}
