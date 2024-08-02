package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"simplebank/util"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	createAccount := createRandomAccount(t)
	findAccount, err := testQueries.GetAccount(context.Background(), createAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, findAccount)

	require.Equal(t, createAccount.ID, findAccount.ID)
	require.Equal(t, createAccount.Owner, findAccount.Owner)
	require.Equal(t, createAccount.Balance, findAccount.Balance)
	require.Equal(t, createAccount.Currency, findAccount.Currency)
	require.WithinDuration(t, createAccount.CreatedAt, findAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	createAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID: createAccount.ID,
		Balance: util.RandomMoney(),
	}

	updateAccount, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updateAccount)

	require.Equal(t, createAccount.ID, updateAccount.ID)
	require.Equal(t, createAccount.Owner, updateAccount.Owner)
	require.Equal(t, arg.Balance, updateAccount.Balance)
	require.Equal(t, createAccount.Currency, updateAccount.Currency)
	require.WithinDuration(t, createAccount.CreatedAt, updateAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	createAccount := createRandomAccount(t)
	
	err := testQueries.DeleteAccount(context.Background(), createAccount.ID)

	require.NoError(t, err)

	findAccount, err := testQueries.GetAccount(context.Background(), createAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, findAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}