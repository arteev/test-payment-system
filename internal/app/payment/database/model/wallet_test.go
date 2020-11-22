package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHashWalletOperation(t *testing.T) {
	type arg struct {
		WalletID uint
		Unit     Unit
		Amount   float64
		Args     []interface{}
	}
	cases := []struct {
		Arg1  arg
		Arg2  arg
		Equal bool
	}{
		{
			Arg1: arg{
				WalletID: 1,
				Unit:     UnitNameDeposit,
				Amount:   15.12000,
			},
			Arg2: arg{
				WalletID: 1,
				Unit:     UnitNameDeposit,
				Amount:   15.12,
			},
			Equal: true,
		},

		{
			Arg1: arg{
				WalletID: 1,
				Unit:     UnitNameDeposit,
				Amount:   15.12000,
			},
			Arg2: arg{
				WalletID: 2,
				Unit:     UnitNameDeposit,
				Amount:   15.12,
			},
			Equal: false,
		},

		{
			Arg1: arg{
				WalletID: 1,
				Unit:     UnitNameTransfer,
				Amount:   15.12,
				Args:     []interface{}{1},
			},
			Arg2: arg{
				WalletID: 1,
				Unit:     UnitNameTransfer,
				Amount:   15.12,
				Args:     []interface{}{2},
			},
			Equal: false,
		},

		{
			Arg1: arg{
				WalletID: 1,
				Unit:     UnitNameTransfer,
				Amount:   15.12,
				Args:     []interface{}{1},
			},
			Arg2: arg{
				WalletID: 1,
				Unit:     UnitNameTransfer,
				Amount:   15.12,
				Args:     []interface{}{1},
			},
			Equal: true,
		},

		{
			Arg1: arg{
				WalletID: 1,
				Unit:     UnitNameTransfer,
				Amount:   15.11,
				Args:     []interface{}{2},
			},
			Arg2: arg{
				WalletID: 1,
				Unit:     UnitNameTransfer,
				Amount:   15.12,
				Args:     []interface{}{2},
			},
			Equal: false,
		},
	}

	for _, test := range cases {
		journal1 := WalletOperJournal{
			WalletID: test.Arg1.WalletID,
			Amount:   test.Arg1.Amount,
			Unit:     test.Arg1.Unit,
		}
		journal2 := WalletOperJournal{
			WalletID: test.Arg2.WalletID,
			Amount:   test.Arg2.Amount,
			Unit:     test.Arg2.Unit,
		}
		hash1 := journal1.GetHashWalletOperation(test.Arg1.Args...)
		hash2 := journal2.GetHashWalletOperation(test.Arg2.Args...)
		if test.Equal {
			assert.Equal(t, hash1, hash2)
		} else {
			assert.NotEqual(t, hash1, hash2)
		}
	}
}
