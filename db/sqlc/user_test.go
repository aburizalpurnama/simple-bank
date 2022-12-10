package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

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

	require.NotZero(t, user.PasswordChangedAt)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestTransactionalFailed(t *testing.T) {

	dsn := "host=localhost user=rizal password=qwerty123 dbname=bpr port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	// R&D
	// dsn := "host=localhost user=ukabima password=PwdDB3bpr6940D!ukbmNew dbname=bpr port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		}})

	require.NoError(t, err)

	err = db.Transaction(func(tx *gorm.DB) error {

		err := insertNegara(tx, "Wakanda")
		if err != nil {
			return err
		}

		err = insertPendidikan(tx, "Pergibahan")
		if err != nil {
			return err
		}

		return nil
	})

	require.Error(t, err)

	fmt.Println(err)

}

func TestTransactionalSuccess(t *testing.T) {

	dsn := "host=localhost user=rizal password=qwerty123 dbname=bpr port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	// R&D
	// dsn := "host=localhost user=ukabima password=PwdDB3bpr6940D!ukbmNew dbname=bpr port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		}})

	require.NoError(t, err)

	err = db.Transaction(func(tx *gorm.DB) error {

		err := insertNegara(tx, "Wakanda")
		if err != nil {
			return err
		}

		err = insertPendidikan(tx, "Pergibahan")
		if err != nil {
			return err
		}

		return nil
	})

	require.NoError(t, err)

}

func insertNegara(tx *gorm.DB, name string) error {

	sql := "insert into m_negara (nm_negara) values (?)"

	err := tx.Exec(sql, name).Error
	if err != nil {
		return err
	}

	return nil
}

func insertPendidikan(tx *gorm.DB, name string) error {

	sql := "insert into m_pendidikan (id, id_entry, d_entry, pendidikan) values (?,?,?,?)"

	err := tx.Exec(sql, 370, "sistem", time.Now(), name).Error
	if err != nil {
		return err
	}

	return nil
}