// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/service/room_service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIRoomService is a mock of IRoomService interface.
type MockIRoomService struct {
	ctrl     *gomock.Controller
	recorder *MockIRoomServiceMockRecorder
}

// MockIRoomServiceMockRecorder is the mock recorder for MockIRoomService.
type MockIRoomServiceMockRecorder struct {
	mock *MockIRoomService
}

// NewMockIRoomService creates a new mock instance.
func NewMockIRoomService(ctrl *gomock.Controller) *MockIRoomService {
	mock := &MockIRoomService{ctrl: ctrl}
	mock.recorder = &MockIRoomServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRoomService) EXPECT() *MockIRoomServiceMockRecorder {
	return m.recorder
}

// FindAllRooms mocks base method.
func (m *MockIRoomService) FindAllRooms(ctx context.Context, userID int) (model.RoomSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllRooms", ctx, userID)
	ret0, _ := ret[0].(model.RoomSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllRooms indicates an expected call of FindAllRooms.
func (mr *MockIRoomServiceMockRecorder) FindAllRooms(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllRooms", reflect.TypeOf((*MockIRoomService)(nil).FindAllRooms), ctx, userID)
}

// FindRoomDetailByRoomID mocks base method.
func (m *MockIRoomService) FindRoomDetailByRoomID(ctx context.Context, userID, roomID, messageID int) (*model.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRoomDetailByRoomID", ctx, userID, roomID, messageID)
	ret0, _ := ret[0].(*model.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRoomDetailByRoomID indicates an expected call of FindRoomDetailByRoomID.
func (mr *MockIRoomServiceMockRecorder) FindRoomDetailByRoomID(ctx, userID, roomID, messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRoomDetailByRoomID", reflect.TypeOf((*MockIRoomService)(nil).FindRoomDetailByRoomID), ctx, userID, roomID, messageID)
}

// SendMessage mocks base method.
func (m_2 *MockIRoomService) SendMessage(ctx context.Context, m *model.Message) (*model.Message, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMessage", ctx, m)
	ret0, _ := ret[0].(*model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockIRoomServiceMockRecorder) SendMessage(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockIRoomService)(nil).SendMessage), ctx, m)
}