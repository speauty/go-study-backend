package db

import (
	"context"
	"database/sql"
	util2 "github.com/speauty/backend/src/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createTmpAccount(t *testing.T) Account {
	user := createTmpUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util2.RandomMoney(),
		Currency: util2.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreateAt)

	return account
}

func TestQueries_CreateAccount(t *testing.T) {
	createTmpAccount(t)
}

func TestQueries_DeleteAccountById(t *testing.T) {
	tmpAccount := createTmpAccount(t)
	err := testQueries.DeleteAccountById(context.Background(), tmpAccount.ID)
	require.NoError(t, err)

	getAccount, err := testQueries.GetAccountById(context.Background(), tmpAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, getAccount)
}

func TestQueries_GetAccountById(t *testing.T) {
	tmpAccount := createTmpAccount(t)

	getAccount, err := testQueries.GetAccountById(context.Background(), tmpAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, getAccount)

	require.Equal(t, tmpAccount.ID, getAccount.ID)
	require.Equal(t, tmpAccount.Owner, getAccount.Owner)
	require.Equal(t, tmpAccount.Balance, getAccount.Balance)
	require.Equal(t, tmpAccount.Currency, getAccount.Currency)
	require.WithinDuration(t, tmpAccount.CreateAt, getAccount.CreateAt, time.Second)
}

func TestQueries_ListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createTmpAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}

func TestQueries_SetAccountDec(t *testing.T) {
	tmpAccount := createTmpAccount(t)
	amount := int64(10)
	arg := SetAccountDecParams{
		ID:      tmpAccount.ID,
		Balance: amount,
	}

	SetAccount, err := testQueries.SetAccountDec(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, SetAccount)
	require.Equal(t, tmpAccount.Balance-amount, SetAccount.Balance)
}

func TestQueries_SetAccountInc(t *testing.T) {
	tmpAccount := createTmpAccount(t)
	amount := int64(10)
	arg := SetAccountIncParams{
		ID:      tmpAccount.ID,
		Balance: amount,
	}

	SetAccount, err := testQueries.SetAccountInc(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, SetAccount)
	require.Equal(t, tmpAccount.Balance+amount, SetAccount.Balance)
}

func TestQueries_UpdateAccountBalance(t *testing.T) {
	tmpAccount := createTmpAccount(t)

	arg := UpdateAccountBalanceParams{
		ID:      tmpAccount.ID,
		Balance: util2.RandomMoney(),
	}

	updateAccount, err := testQueries.UpdateAccountBalance(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updateAccount)

	require.Equal(t, tmpAccount.ID, updateAccount.ID)
	require.Equal(t, tmpAccount.Owner, updateAccount.Owner)
	require.Equal(t, arg.Balance, updateAccount.Balance)
	require.Equal(t, tmpAccount.Currency, updateAccount.Currency)
	require.WithinDuration(t, tmpAccount.CreateAt, updateAccount.CreateAt, time.Second)
}
