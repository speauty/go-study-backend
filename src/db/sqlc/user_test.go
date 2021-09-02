package db

import (
	"context"
	util "github.com/speauty/backend/src/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createTmpUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreateAt)

	return user
}

func TestQueries_CreateUser(t *testing.T) {
	createTmpUser(t)
}

func TestQueries_GetUserByUsername(t *testing.T) {
	tmpUser := createTmpUser(t)

	getUser, err := testQueries.GetUserByUsername(context.Background(), tmpUser.Username)

	require.NoError(t, err)
	require.NotEmpty(t, getUser)

	require.Equal(t, tmpUser.Username, getUser.Username)
	require.Equal(t, tmpUser.HashedPassword, getUser.HashedPassword)
	require.Equal(t, tmpUser.FullName, getUser.FullName)
	require.Equal(t, tmpUser.Email, getUser.Email)
	require.WithinDuration(t, tmpUser.PasswordChangedAt, getUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, tmpUser.CreateAt, getUser.CreateAt, time.Second)
}
