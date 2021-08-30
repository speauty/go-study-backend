package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createTmpEntry(t *testing.T) Entry {
	tmpAccount := createTmpAccount(t)
	arg := CreateEntryParams{
		AccountID: tmpAccount.ID,
		Amount: int64(10),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.NotZero(t, entry.ID)
	require.Equal(t, entry.AccountID, tmpAccount.ID)
	require.NotZero(t, entry.CreateAt)

	return entry
}

func TestQueries_CreateEntry(t *testing.T) {
	createTmpEntry(t)
}

func TestQueries_DeleteEntryById(t *testing.T) {
	tmpEntry := createTmpEntry(t)
	err := testQueries.DeleteEntryById(context.Background(), tmpEntry.ID)
	require.NoError(t, err)

	getEntry, err := testQueries.GetEntryById(context.Background(), tmpEntry.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, getEntry)
}

func TestQueries_GetEntryById(t *testing.T) {
	tmpEntry := createTmpEntry(t)

	getEntry, err := testQueries.GetEntryById(context.Background(), tmpEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getEntry)
	require.Equal(t, tmpEntry.ID, getEntry.ID)
	require.Equal(t, tmpEntry.AccountID, getEntry.AccountID)
	require.Equal(t, tmpEntry.Amount, getEntry.Amount)
	require.WithinDuration(t, tmpEntry.CreateAt, getEntry.CreateAt, time.Second)
}

func TestQueries_ListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTmpEntry(t)
	}

	arg := ListEntriesParams{
		Limit: 5,
		Offset: 5,
	}

	entries,err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
