package repository

import (
	"backend/model/entity"
	repository "backend/repository/interface"
	"context"
	"database/sql"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type AccountRepo struct {
	DB *sql.DB
}

// function provider
func NewAccountRepo(db *sql.DB) repository.IAccountRepo {
	return &AccountRepo{db}
}

// method get all data accounts
func (a *AccountRepo) GetAll(ctx context.Context) ([]entity.Account, error) {
	// cretate new context tracing
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "repository GetAll accounts")
	defer span.Finish()

	// create statement
	statement, err := a.DB.PrepareContext(ctxTracing, "SELECT email, username, password FROM account")
	defer statement.Close()
	if err != nil {
		return nil, err
	}

	// execute statement
	rows, err := statement.QueryContext(ctxTracing)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var accounts []entity.Account
	for rows.Next() {
		// scan result to object
		var account entity.Account
		if err := rows.Scan(&account.Email, &account.Username, &account.Password); err != nil {
			return nil, err
		}

		// append to response object
		accounts = append(accounts, account)
	}

	// check if data not found
	if len(accounts) == 0 {
		return nil, errors.New("record not found")
	}

	// success get all data accounts
	span.LogFields(
		log.Int("total-data", len(accounts)),
	)
	return accounts, nil
}

// method insert data account
func (a *AccountRepo) Insert(ctx context.Context, input *entity.Account) (*entity.Account, error) {
	//TODO implement me
	panic("implement me")
}
