// Code generated by mockery v2.14.0. DO NOT EDIT.

//go:build !codeanalysis
// +build !codeanalysis

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "example/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// InsertReservation provides a mock function with given fields: ctx, reservation
func (_m *Client) InsertReservation(ctx context.Context, reservation types.Reservation) (types.Reservation, error) {
	ret := _m.Called(ctx, reservation)

	var r0 types.Reservation
	if rf, ok := ret.Get(0).(func(context.Context, types.Reservation) types.Reservation); ok {
		r0 = rf(ctx, reservation)
	} else {
		r0 = ret.Get(0).(types.Reservation)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.Reservation) error); ok {
		r1 = rf(ctx, reservation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectReservationByID provides a mock function with given fields: ctx, id
func (_m *Client) SelectReservationByID(ctx context.Context, id int64) (types.Reservation, error) {
	ret := _m.Called(ctx, id)

	var r0 types.Reservation
	if rf, ok := ret.Get(0).(func(context.Context, int64) types.Reservation); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(types.Reservation)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectReservations provides a mock function with given fields: ctx
func (_m *Client) SelectReservations(ctx context.Context) ([]types.Reservation, error) {
	ret := _m.Called(ctx)

	var r0 []types.Reservation
	if rf, ok := ret.Get(0).(func(context.Context) []types.Reservation); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Reservation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectReservationsByHotel provides a mock function with given fields: ctx, hotelID
func (_m *Client) SelectReservationsByHotel(ctx context.Context, hotelID int64) ([]types.Reservation, error) {
	ret := _m.Called(ctx, hotelID)

	var r0 []types.Reservation
	if rf, ok := ret.Get(0).(func(context.Context, int64) []types.Reservation); ok {
		r0 = rf(ctx, hotelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Reservation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, hotelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
