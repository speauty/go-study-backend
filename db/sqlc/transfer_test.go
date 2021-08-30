package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createTmpTransfer(t *testing.T) Transfer {
	fromTmpAccount := createTmpAccount(t)
	toTmpAccount := createTmpAccount(t)
	amount := int64(10)

	arg := CreateTransferParams{
		FromAccountID: fromTmpAccount.ID,
		ToAccountID: toTmpAccount.ID,
		Amount: amount,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, fromTmpAccount.ID)
	require.Equal(t, transfer.ToAccountID, toTmpAccount.ID)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreateAt)

	return transfer
}

func TestQueries_CreateTransfer(t *testing.T) {
	createTmpTransfer(t)
}

func TestQueries_GetTransferById(t *testing.T) {
	tmpTransfer := createTmpTransfer(t)
	getTransfer, err := testQueries.GetTransferById(context.Background(), tmpTransfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getTransfer)
	require.Equal(t, tmpTransfer.ID, getTransfer.ID)
	require.Equal(t, tmpTransfer.FromAccountID, getTransfer.FromAccountID)
	require.Equal(t, tmpTransfer.ToAccountID, getTransfer.ToAccountID)
	require.Equal(t, tmpTransfer.Amount, getTransfer.Amount)
	require.WithinDuration(t, tmpTransfer.CreateAt, getTransfer.CreateAt, time.Second)
}

func TestQueries_DeleteTransferById(t *testing.T) {
	tmpTransfer := createTmpTransfer(t)
	err := testQueries.DeleteTransferById(context.Background(), tmpTransfer.ID)
	require.NoError(t, err)

	getTransfer, err := testQueries.GetTransferById(context.Background(), tmpTransfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, getTransfer)
}

func TestQueries_ListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTmpTransfer(t)
	}

	arg := ListTransfersParams{
		Limit: 5,
		Offset: 5,
	}

	transfers,err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
