package service

import (
	"context"
	"errors"
	"kings-comp/internal/entity"
	"kings-comp/internal/repository"
	"time"
)

const (
	DefaultState = "home"
)

type AccountService struct {
	accounts repository.AccountRepository
}

func NewAccountService(rep repository.AccountRepository) *AccountService {
	return &AccountService{accounts: rep}
}

// CreateOrUpdate creates a new user in the data store or updates the existing user
func (a *AccountService) CreateOrUpdate(ctx context.Context,
	account entity.Account) (entity.Account, bool, error) {
	savedAccount, err := a.accounts.Get(ctx, account.EntityID())
	// user exists in the database
	if err == nil {
		if savedAccount.Username != account.Username || savedAccount.FirstName != account.FirstName {
			savedAccount.Username = account.Username
			savedAccount.FirstName = account.FirstName
			return savedAccount, false, a.accounts.Save(ctx, savedAccount)
		}
		return savedAccount, false, nil
	}

	if errors.Is(err, repository.ErrNotFound) {
		account.JoinedAt = time.Now()
		account.State = DefaultState
		return account, true, a.accounts.Save(ctx, account)
	}

	return entity.Account{}, false, err
}

func (a *AccountService) Update(ctx context.Context, account entity.Account) error {
	return a.accounts.Save(ctx, account)
}
