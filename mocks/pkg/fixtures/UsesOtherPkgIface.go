// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	test "github.com/vektra/mockery/v2/pkg/fixtures"
)

// UsesOtherPkgIface is an autogenerated mock type for the UsesOtherPkgIface type
type UsesOtherPkgIface struct {
	mock.Mock
}

// DoSomethingElse provides a mock function with given fields: obj
func (_m *UsesOtherPkgIface) DoSomethingElse(obj test.Sibling) {
	_m.Called(obj)
}
