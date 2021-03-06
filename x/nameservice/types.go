package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Whois struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}

var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

func NewWhois() Whois {
	return Whois{
		Price: MinNamePrice,
	}
}
