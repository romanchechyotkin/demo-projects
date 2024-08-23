// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/service.go
//
// Generated by this command:
//
//	mockgen -source internal/service/service.go -destination internal/service/mocks/mocks.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/romanchechyotkin/avito_test_task/internal/entity"
	service "github.com/romanchechyotkin/avito_test_task/internal/service"
	gomock "go.uber.org/mock/gomock"
)

// MockSender is a mock of Sender interface.
type MockSender struct {
	ctrl     *gomock.Controller
	recorder *MockSenderMockRecorder
}

// MockSenderMockRecorder is the mock recorder for MockSender.
type MockSenderMockRecorder struct {
	mock *MockSender
}

// NewMockSender creates a new mock instance.
func NewMockSender(ctrl *gomock.Controller) *MockSender {
	mock := &MockSender{ctrl: ctrl}
	mock.recorder = &MockSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSender) EXPECT() *MockSenderMockRecorder {
	return m.recorder
}

// Notify mocks base method.
func (m *MockSender) Notify() chan<- uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Notify")
	ret0, _ := ret[0].(chan<- uint)
	return ret0
}

// Notify indicates an expected call of Notify.
func (mr *MockSenderMockRecorder) Notify() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockSender)(nil).Notify))
}

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuth) CreateUser(ctx context.Context, input *service.AuthCreateUserInput) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, input)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthMockRecorder) CreateUser(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuth)(nil).CreateUser), ctx, input)
}

// GenerateToken mocks base method.
func (m *MockAuth) GenerateToken(ctx context.Context, input *service.AuthGenerateTokenInput) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", ctx, input)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthMockRecorder) GenerateToken(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuth)(nil).GenerateToken), ctx, input)
}

// ParseToken mocks base method.
func (m *MockAuth) ParseToken(accessToken string) (*service.TokenClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", accessToken)
	ret0, _ := ret[0].(*service.TokenClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthMockRecorder) ParseToken(accessToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuth)(nil).ParseToken), accessToken)
}

// MockHouse is a mock of House interface.
type MockHouse struct {
	ctrl     *gomock.Controller
	recorder *MockHouseMockRecorder
}

// MockHouseMockRecorder is the mock recorder for MockHouse.
type MockHouseMockRecorder struct {
	mock *MockHouse
}

// NewMockHouse creates a new mock instance.
func NewMockHouse(ctrl *gomock.Controller) *MockHouse {
	mock := &MockHouse{ctrl: ctrl}
	mock.recorder = &MockHouseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHouse) EXPECT() *MockHouseMockRecorder {
	return m.recorder
}

// CreateHouse mocks base method.
func (m *MockHouse) CreateHouse(ctx context.Context, input *service.HouseCreateInput) (*entity.House, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHouse", ctx, input)
	ret0, _ := ret[0].(*entity.House)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHouse indicates an expected call of CreateHouse.
func (mr *MockHouseMockRecorder) CreateHouse(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHouse", reflect.TypeOf((*MockHouse)(nil).CreateHouse), ctx, input)
}

// CreateSubscription mocks base method.
func (m *MockHouse) CreateSubscription(ctx context.Context, input *service.CreateSubscriptionInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscription", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSubscription indicates an expected call of CreateSubscription.
func (mr *MockHouseMockRecorder) CreateSubscription(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscription", reflect.TypeOf((*MockHouse)(nil).CreateSubscription), ctx, input)
}

// GetHouseFlats mocks base method.
func (m *MockHouse) GetHouseFlats(ctx context.Context, input *service.GetHouseFlatsInput) ([]*entity.Flat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHouseFlats", ctx, input)
	ret0, _ := ret[0].([]*entity.Flat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHouseFlats indicates an expected call of GetHouseFlats.
func (mr *MockHouseMockRecorder) GetHouseFlats(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHouseFlats", reflect.TypeOf((*MockHouse)(nil).GetHouseFlats), ctx, input)
}

// MockFlat is a mock of Flat interface.
type MockFlat struct {
	ctrl     *gomock.Controller
	recorder *MockFlatMockRecorder
}

// MockFlatMockRecorder is the mock recorder for MockFlat.
type MockFlatMockRecorder struct {
	mock *MockFlat
}

// NewMockFlat creates a new mock instance.
func NewMockFlat(ctrl *gomock.Controller) *MockFlat {
	mock := &MockFlat{ctrl: ctrl}
	mock.recorder = &MockFlatMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlat) EXPECT() *MockFlatMockRecorder {
	return m.recorder
}

// CreateFlat mocks base method.
func (m *MockFlat) CreateFlat(ctx context.Context, input *service.FlatCreateInput) (*entity.Flat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFlat", ctx, input)
	ret0, _ := ret[0].(*entity.Flat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFlat indicates an expected call of CreateFlat.
func (mr *MockFlatMockRecorder) CreateFlat(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFlat", reflect.TypeOf((*MockFlat)(nil).CreateFlat), ctx, input)
}

// UpdateFlat mocks base method.
func (m *MockFlat) UpdateFlat(ctx context.Context, input *service.FlatUpdateInput) (*entity.Flat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFlat", ctx, input)
	ret0, _ := ret[0].(*entity.Flat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFlat indicates an expected call of UpdateFlat.
func (mr *MockFlatMockRecorder) UpdateFlat(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFlat", reflect.TypeOf((*MockFlat)(nil).UpdateFlat), ctx, input)
}
