// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/user_usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIUserUsecase is a mock of IUserUsecase interface.
type MockIUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIUserUsecaseMockRecorder
}

// MockIUserUsecaseMockRecorder is the mock recorder for MockIUserUsecase.
type MockIUserUsecaseMockRecorder struct {
	mock *MockIUserUsecase
}

// NewMockIUserUsecase creates a new mock instance.
func NewMockIUserUsecase(ctrl *gomock.Controller) *MockIUserUsecase {
	mock := &MockIUserUsecase{ctrl: ctrl}
	mock.recorder = &MockIUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserUsecase) EXPECT() *MockIUserUsecaseMockRecorder {
	return m.recorder
}

// FindAllRooms mocks base method.
func (m *MockIUserUsecase) FindAllRooms(ctx context.Context, userID int) (*model.Rooms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllRooms", ctx, userID)
	ret0, _ := ret[0].(*model.Rooms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllRooms indicates an expected call of FindAllRooms.
func (mr *MockIUserUsecaseMockRecorder) FindAllRooms(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllRooms", reflect.TypeOf((*MockIUserUsecase)(nil).FindAllRooms), ctx, userID)
}

// FindAllUsers mocks base method.
func (m *MockIUserUsecase) FindAllUsers(ctx context.Context) (model.UserSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllUsers", ctx)
	ret0, _ := ret[0].(model.UserSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllUsers indicates an expected call of FindAllUsers.
func (mr *MockIUserUsecaseMockRecorder) FindAllUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllUsers", reflect.TypeOf((*MockIUserUsecase)(nil).FindAllUsers), ctx)
}

// FindRoomDetailByRoomID mocks base method.
func (m *MockIUserUsecase) FindRoomDetailByRoomID(ctx context.Context, userID, roomID, messageID int) (*model.RoomDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRoomDetailByRoomID", ctx, userID, roomID, messageID)
	ret0, _ := ret[0].(*model.RoomDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRoomDetailByRoomID indicates an expected call of FindRoomDetailByRoomID.
func (mr *MockIUserUsecaseMockRecorder) FindRoomDetailByRoomID(ctx, userID, roomID, messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRoomDetailByRoomID", reflect.TypeOf((*MockIUserUsecase)(nil).FindRoomDetailByRoomID), ctx, userID, roomID, messageID)
}

// FindUserByUserID mocks base method.
func (m *MockIUserUsecase) FindUserByUserID(ctx context.Context, userID int) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByUserID", ctx, userID)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByUserID indicates an expected call of FindUserByUserID.
func (mr *MockIUserUsecaseMockRecorder) FindUserByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUserID", reflect.TypeOf((*MockIUserUsecase)(nil).FindUserByUserID), ctx, userID)
}

// FindUserDetailByUserID mocks base method.
func (m *MockIUserUsecase) FindUserDetailByUserID(ctx context.Context, userID int) (*model.UserDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserDetailByUserID", ctx, userID)
	ret0, _ := ret[0].(*model.UserDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserDetailByUserID indicates an expected call of FindUserDetailByUserID.
func (mr *MockIUserUsecaseMockRecorder) FindUserDetailByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserDetailByUserID", reflect.TypeOf((*MockIUserUsecase)(nil).FindUserDetailByUserID), ctx, userID)
}

// SendMessage mocks base method.
func (m_2 *MockIUserUsecase) SendMessage(ctx context.Context, userID, roomID int, m *model.NewMessage) (*model.Message, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMessage", ctx, userID, roomID, m)
	ret0, _ := ret[0].(*model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockIUserUsecaseMockRecorder) SendMessage(ctx, userID, roomID, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockIUserUsecase)(nil).SendMessage), ctx, userID, roomID, m)
}
