package repository

import (
	"backend/model/entity"
	repository "backend/repository/interface"
	"context"
	"database/sql"
)

type AccountRepo struct {
	DB *sql.DB
}

// function provider
func NewAccountRepo(db *sql.DB) repository.IAccountRepo {
	return &AccountRepo{db}
}

func (a *AccountRepo) Insert(ctx context.Context, input *entity.Account) (*entity.Account, error) {
	//TODO implement me
	panic("implement me")
}
