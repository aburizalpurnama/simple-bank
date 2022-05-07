package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	dtb *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		dtb:     db,
		Queries: New(db),
	}
}

var txKey = struct{}{}

func (s *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := s.execTx(ctx, func(q *Queries) error {

		var err error
		txName := ctx.Value(txKey)

		fmt.Println(txName, "Create Transfer")
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "Create Entry 1")
		result.FromEntry, err = q.CreateEntrie(ctx, CreateEntrieParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "Create Entry 2")
		result.ToEntry, err = q.CreateEntrie(ctx, CreateEntrieParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// udpdate account's balance
		fmt.Println(txName, "GetAccount 1")
		account1, err := q.GetAccountForUpdate(context.Background(), arg.FromAccountID)
		if err != nil {
			return err
		}

		fmt.Println(txName, "UpdateAccount 1")
		result.FromAccount, err = q.UpdateAccount(context.Background(), UpdateAccountParams{
			ID:      account1.ID,
			Balance: account1.Balance - arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "GetAccount 2")
		account2, err := q.GetAccountForUpdate(context.Background(), arg.ToAccountID)
		if err != nil {
			return err
		}

		fmt.Println(txName, "UpdateAccount 2")
		result.ToAccount, err = q.UpdateAccount(context.Background(), UpdateAccountParams{
			ID:      account2.ID,
			Balance: account2.Balance + arg.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.dtb.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
