package mock

import (
	"backend/model/entity"
	"context"
	"github.com/stretchr/testify/mock"
)

type AccountRepoMock struct {
	mock.Mock
}

func (a *AccountRepoMock) Insert(ctx context.Context, input *entity.Account) (*entity.Account, error) {
	args := a.Mock.Called(ctx, input)

	account := args.Get(0)
	if account == nil {
		return nil, args.Error(1)
	}

	return account.(*entity.Account), args.Error(1)
}

func (a *AccountRepoMock) GetAll(ctx context.Context) ([]entity.Account, error) {
	args := a.Mock.Called(ctx)

	accounts := args.Get(0)
	if accounts == nil {
		return nil, args.Error(1)
	}

	return accounts.([]entity.Account), args.Error(1)
}
