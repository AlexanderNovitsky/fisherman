// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	commands "fisherman/commands"

	mock "github.com/stretchr/testify/mock"
)

// CliCommand is an autogenerated mock type for the CliCommand type
type CliCommand struct {
	mock.Mock
}

// Name provides a mock function with given fields:
func (_m *CliCommand) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Run provides a mock function with given fields: ctx, args
func (_m *CliCommand) Run(ctx *commands.CommandContext, args []string) error {
	ret := _m.Called(ctx, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(*commands.CommandContext, []string) error); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
