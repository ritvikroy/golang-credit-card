package banksdb_test

import (
	"testing"

	"github.com/bitcyber/golang-credit-card"
	"github.com/stretchr/testify/require"
)

func TestFindBankEmptyCreditCard(t *testing.T) {
	bank := banksdb.FindBank("")
	require.Nil(t, bank)
}

func TestFindBankInvalidCreditCard(t *testing.T) {
	bank := banksdb.FindBank("no-digits")
	require.Nil(t, bank)
}
