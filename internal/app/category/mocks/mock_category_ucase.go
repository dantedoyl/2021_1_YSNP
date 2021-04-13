// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2021_1_YSNP/internal/app/category (interfaces: CategoryUsecase)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	errors "github.com/go-park-mail-ru/2021_1_YSNP/internal/app/errors"
	models "github.com/go-park-mail-ru/2021_1_YSNP/internal/app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockCategoryUsecase is a mock of CategoryUsecase interface.
type MockCategoryUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryUsecaseMockRecorder
}

// MockCategoryUsecaseMockRecorder is the mock recorder for MockCategoryUsecase.
type MockCategoryUsecaseMockRecorder struct {
	mock *MockCategoryUsecase
}

// NewMockCategoryUsecase creates a new mock instance.
func NewMockCategoryUsecase(ctrl *gomock.Controller) *MockCategoryUsecase {
	mock := &MockCategoryUsecase{ctrl: ctrl}
	mock.recorder = &MockCategoryUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryUsecase) EXPECT() *MockCategoryUsecaseMockRecorder {
	return m.recorder
}

// GetAllCategories mocks base method.
func (m *MockCategoryUsecase) GetAllCategories() ([]*models.Category, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCategories")
	ret0, _ := ret[0].([]*models.Category)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetAllCategories indicates an expected call of GetAllCategories.
func (mr *MockCategoryUsecaseMockRecorder) GetAllCategories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCategories", reflect.TypeOf((*MockCategoryUsecase)(nil).GetAllCategories))
}