// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2021_1_YSNP/internal/app/product (interfaces: ProductUsecase)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	errors "github.com/go-park-mail-ru/2021_1_YSNP/internal/app/errors"
	models "github.com/go-park-mail-ru/2021_1_YSNP/internal/app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockProductUsecase is a mock of ProductUsecase interface.
type MockProductUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockProductUsecaseMockRecorder
}

// MockProductUsecaseMockRecorder is the mock recorder for MockProductUsecase.
type MockProductUsecaseMockRecorder struct {
	mock *MockProductUsecase
}

// NewMockProductUsecase creates a new mock instance.
func NewMockProductUsecase(ctrl *gomock.Controller) *MockProductUsecase {
	mock := &MockProductUsecase{ctrl: ctrl}
	mock.recorder = &MockProductUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductUsecase) EXPECT() *MockProductUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductUsecase) Create(arg0 *models.ProductData) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockProductUsecaseMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductUsecase)(nil).Create), arg0)
}

// DislikeProduct mocks base method.
func (m *MockProductUsecase) DislikeProduct(arg0, arg1 uint64) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DislikeProduct", arg0, arg1)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// DislikeProduct indicates an expected call of DislikeProduct.
func (mr *MockProductUsecaseMockRecorder) DislikeProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DislikeProduct", reflect.TypeOf((*MockProductUsecase)(nil).DislikeProduct), arg0, arg1)
}

// GetByID mocks base method.
func (m *MockProductUsecase) GetByID(arg0 uint64) (*models.ProductData, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*models.ProductData)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockProductUsecaseMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockProductUsecase)(nil).GetByID), arg0)
}

// GetUserFavorite mocks base method.
func (m *MockProductUsecase) GetUserFavorite(arg0 uint64, arg1 *models.Page) ([]*models.ProductListData, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFavorite", arg0, arg1)
	ret0, _ := ret[0].([]*models.ProductListData)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetUserFavorite indicates an expected call of GetUserFavorite.
func (mr *MockProductUsecaseMockRecorder) GetUserFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFavorite", reflect.TypeOf((*MockProductUsecase)(nil).GetUserFavorite), arg0, arg1)
}

// LikeProduct mocks base method.
func (m *MockProductUsecase) LikeProduct(arg0, arg1 uint64) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LikeProduct", arg0, arg1)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// LikeProduct indicates an expected call of LikeProduct.
func (mr *MockProductUsecaseMockRecorder) LikeProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LikeProduct", reflect.TypeOf((*MockProductUsecase)(nil).LikeProduct), arg0, arg1)
}

// ListLatest mocks base method.
func (m *MockProductUsecase) ListLatest(arg0 *uint64, arg1 *models.Page) ([]*models.ProductListData, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLatest", arg0, arg1)
	ret0, _ := ret[0].([]*models.ProductListData)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// ListLatest indicates an expected call of ListLatest.
func (mr *MockProductUsecaseMockRecorder) ListLatest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLatest", reflect.TypeOf((*MockProductUsecase)(nil).ListLatest), arg0, arg1)
}

// SetTariff mocks base method.
func (m *MockProductUsecase) SetTariff(arg0 uint64, arg1 int) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTariff", arg0, arg1)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// SetTariff indicates an expected call of SetTariff.
func (mr *MockProductUsecaseMockRecorder) SetTariff(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTariff", reflect.TypeOf((*MockProductUsecase)(nil).SetTariff), arg0, arg1)
}

// UpdatePhoto mocks base method.
func (m *MockProductUsecase) UpdatePhoto(arg0 uint64, arg1 []string) (*models.ProductData, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhoto", arg0, arg1)
	ret0, _ := ret[0].(*models.ProductData)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// UpdatePhoto indicates an expected call of UpdatePhoto.
func (mr *MockProductUsecaseMockRecorder) UpdatePhoto(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhoto", reflect.TypeOf((*MockProductUsecase)(nil).UpdatePhoto), arg0, arg1)
}

// UserAdList mocks base method.
func (m *MockProductUsecase) UserAdList(arg0 uint64, arg1 *models.Page) ([]*models.ProductListData, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserAdList", arg0, arg1)
	ret0, _ := ret[0].([]*models.ProductListData)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// UserAdList indicates an expected call of UserAdList.
func (mr *MockProductUsecaseMockRecorder) UserAdList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserAdList", reflect.TypeOf((*MockProductUsecase)(nil).UserAdList), arg0, arg1)
}
