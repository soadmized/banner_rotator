// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	context "context"

	slot "github.com/soadmized/banners_rotator/internal/slot"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, id, desc
func (_m *Repository) Create(ctx context.Context, id slot.ID, desc string) error {
	ret := _m.Called(ctx, id, desc)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, slot.ID, string) error); ok {
		r0 = rf(ctx, id, desc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *Repository) Get(ctx context.Context, id slot.ID) (*slot.Slot, error) {
	ret := _m.Called(ctx, id)

	var r0 *slot.Slot
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, slot.ID) (*slot.Slot, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, slot.ID) *slot.Slot); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*slot.Slot)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, slot.ID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
