package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func TestCreateAccount(t *testing.T) {
	user := createRandomUser(t)
	args := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	result, err := testQueries.ListAccounts(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, result, 5)

	for _, account := range result {
		require.NotEmpty(t, account)
	}
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	newBalance := util.RandomMoney()
	account2, err := testQueries.UpdateAccount(context.Background(), UpdateAccountParams{ID: account1.ID, Balance: newBalance})

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, newBalance, account2.Balance)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)

	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)
	require.EqualError(t, sql.ErrNoRows, err.Error())
	require.Empty(t, account2)
}

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	args := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestAddingDate(t *testing.T) {
	now := time.Now()

	fmt.Println("now :", now)
	fmt.Println("adding:", now.AddDate(0, 0, 60))

	a := time.Date(2022, time.April, 10, 0, 0, 0, 0, time.Local)
	b := time.Date(2022, time.April, 11, 0, 0, 0, 0, time.Local)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a == b)
}

func TestErrorCheck(t *testing.T) {
	var x []int
	
	fmt.Println("before:", len(x))

	for i := 0; i < 10; i++ {
		x = append(x, i)
	}

	fmt.Println("after:", len(x))
}
