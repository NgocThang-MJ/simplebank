package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/liquiddev99/simplebank/util"
)

func createRandomUser(t *testing.T) User {
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
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	newUser := createRandomUser(t)

	user, err := testQueries.GetUser(context.Background(), newUser.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, newUser.Username, user.Username)
	require.Equal(t, newUser.FullName, user.FullName)
	require.Equal(t, newUser.Email, user.Email)
	require.Equal(t, newUser.HashedPassword, user.HashedPassword)
	require.WithinDuration(t, newUser.PasswordChangedAt, user.PasswordChangedAt, time.Second)
	require.WithinDuration(t, newUser.CreatedAt, user.CreatedAt, time.Second)
}
