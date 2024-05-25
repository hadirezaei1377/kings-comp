// Code generated by mockery v2.40.2. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "kings-comp/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// AccountRepository is an autogenerated mock type for the AccountRepository type
type AccountRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, id
func (_m *AccountRepository) Get(ctx context.Context, id entity.ID) (entity.Account, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.ID) (entity.Account, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.ID) entity.Account); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.ID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MGet provides a mock function with given fields: ctx, ids
func (_m *AccountRepository) MGet(ctx context.Context, ids ...entity.ID) ([]entity.Account, error) {
	_va := make([]interface{}, len(ids))
	for _i := range ids {
		_va[_i] = ids[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for MGet")
	}

	var r0 []entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...entity.ID) ([]entity.Account, error)); ok {
		return rf(ctx, ids...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...entity.ID) []entity.Account); ok {
		r0 = rf(ctx, ids...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...entity.ID) error); ok {
		r1 = rf(ctx, ids...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, ent
func (_m *AccountRepository) Save(ctx context.Context, ent entity.Account) error {
	ret := _m.Called(ctx, ent)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Account) error); ok {
		r0 = rf(ctx, ent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAccountRepository creates a new instance of AccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountRepository {
	mock := &AccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}