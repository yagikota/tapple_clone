// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/service/user_service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockIUserService is a mock of IUserService interface.
type MockIUserService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserServiceMockRecorder
}

// MockIUserServiceMockRecorder is the mock recorder for MockIUserService.
type MockIUserServiceMockRecorder struct {
	mock *MockIUserService
}

// NewMockIUserService creates a new mock instance.
func NewMockIUserService(ctrl *gomock.Controller) *MockIUserService {
	mock := &MockIUserService{ctrl: ctrl}
	mock.recorder = &MockIUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserService) EXPECT() *MockIUserServiceMockRecorder {
	return m.recorder
}

// FindAllRooms mocks base method.
func (m *MockIUserService) FindAllRooms(ctx context.Context, userID int) (entity.RoomSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllRooms", ctx, userID)
	ret0, _ := ret[0].(entity.RoomSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllRooms indicates an expected call of FindAllRooms.
func (mr *MockIUserServiceMockRecorder) FindAllRooms(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllRooms", reflect.TypeOf((*MockIUserService)(nil).FindAllRooms), ctx, userID)
}

// FindAllUsers mocks base method.
func (m *MockIUserService) FindAllUsers(ctx context.Context) (entity.UserSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllUsers", ctx)
	ret0, _ := ret[0].(entity.UserSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllUsers indicates an expected call of FindAllUsers.
func (mr *MockIUserServiceMockRecorder) FindAllUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllUsers", reflect.TypeOf((*MockIUserService)(nil).FindAllUsers), ctx)
}

// FindRoomDetailByRoomID mocks base method.
func (m *MockIUserService) FindRoomDetailByRoomID(ctx context.Context, userID, roomID int) (*entity.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRoomDetailByRoomID", ctx, userID, roomID)
	ret0, _ := ret[0].(*entity.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRoomDetailByRoomID indicates an expected call of FindRoomDetailByRoomID.
func (mr *MockIUserServiceMockRecorder) FindRoomDetailByRoomID(ctx, userID, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRoomDetailByRoomID", reflect.TypeOf((*MockIUserService)(nil).FindRoomDetailByRoomID), ctx, userID, roomID)
}

// FindUserByUserID mocks base method.
func (m *MockIUserService) FindUserByUserID(ctx context.Context, userID int) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByUserID", ctx, userID)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByUserID indicates an expected call of FindUserByUserID.
func (mr *MockIUserServiceMockRecorder) FindUserByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUserID", reflect.TypeOf((*MockIUserService)(nil).FindUserByUserID), ctx, userID)
}

// SendMessage mocks base method.
func (m_2 *MockIUserService) SendMessage(ctx context.Context, m *entity.Message) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMessage", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockIUserServiceMockRecorder) SendMessage(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockIUserService)(nil).SendMessage), ctx, m)
}
