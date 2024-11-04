package db

import (
	"context"
	"database/sql"
	"github.com/document/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createAccountRandom(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Balance, account.Balance)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createAccountRandom(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createAccountRandom(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.Equal(t, account1, account2)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createAccountRandom(t)
	}

	arg := ListAccountParams{
		Limit:  5,
		Offset: 5,
	}

	account, err := testQueries.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, account, 5)

	for _, account := range account {
		require.NotEmpty(t, account)
	}
}

func TestUpdateAccount(t *testing.T) {
	account1 := createAccountRandom(t)
	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomInt(0, 1000),
	}
	err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEqual(t, account1, account2)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createAccountRandom(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}
