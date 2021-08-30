package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore_TransferTx(t *testing.T) {
	store := NewStore(testDB)

	fromAccount := createTmpAccount(t)
	toAccount := createTmpAccount(t)

	n := 2
	amount := int64(10)
	errs := make(chan error)
	results := make(chan TransferTxResult)
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountId: fromAccount.ID,
				ToAccountId:   toAccount.ID,
				Amount:        amount,
			})

			errs <-err
			results <-result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs

		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, fromAccount.ID, transfer.FromAccountID)
		require.Equal(t, toAccount.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreateAt)

		_,err = store.GetTransferById(context.Background(), transfer.ID)
		require.NoError(t, err)

		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, fromAccount.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreateAt)

		_,err = store.GetEntryById(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, toAccount.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreateAt)

		_,err = store.GetEntryById(context.Background(), toEntry.ID)
		require.NoError(t, err)

		fromAccountResult := result.FromAccount
		require.NotEmpty(t, fromAccountResult)
		require.Equal(t, fromAccountResult.ID, fromAccount.ID)

		toAccountResult := result.ToAccount
		require.NotEmpty(t, toAccountResult)
		require.Equal(t, toAccountResult.ID, toAccount.ID)

		diff1 := fromAccount.Balance-fromAccountResult.Balance
		diff2 := toAccountResult.Balance-toAccount.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1>0)
		require.True(t, diff1%amount==0)

		k := int(diff1/amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	fromAccountUpdated,err := testQueries.GetAccountById(context.Background(), fromAccount.ID)
	require.NoError(t, err)

	toAccountUpdated,err := testQueries.GetAccountById(context.Background(), toAccount.ID)
	require.NoError(t, err)

	require.Equal(t, fromAccount.Balance-int64(n)*amount, fromAccountUpdated.Balance)
	require.Equal(t, toAccount.Balance+int64(n)*amount, toAccountUpdated.Balance)


}

func TestTransferTxDeadLock(t *testing.T) {
	store := NewStore(testDB)

	fromAccount := createTmpAccount(t)
	toAccount := createTmpAccount(t)

	n := 10
	amount := int64(10)
	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountId := fromAccount.ID
		toAccountId := toAccount.ID

		if i%2 == 1 {
			fromAccountId = toAccount.ID
			toAccountId = fromAccount.ID
		}
		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountId: fromAccountId,
				ToAccountId:   toAccountId,
				Amount:        amount,
			})

			errs <-err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	fromAccountUpdated,err := testQueries.GetAccountById(context.Background(), fromAccount.ID)
	require.NoError(t, err)

	toAccountUpdated,err := testQueries.GetAccountById(context.Background(), toAccount.ID)
	require.NoError(t, err)

	require.Equal(t, fromAccount.Balance, fromAccountUpdated.Balance)
	require.Equal(t, toAccount.Balance, toAccountUpdated.Balance)


}