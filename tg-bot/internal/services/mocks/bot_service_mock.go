package mocks

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	"reflect"
)

type MockBotService struct {
	ctrl     *gomock.Controller
	recorder *MockBotServiceMockRecorder
}

type MockBotServiceMockRecorder struct {
	mock *MockBotService
}

func NewMockBotService(ctrl *gomock.Controller) *MockBotService {
	mock := &MockBotService{ctrl: ctrl}
	mock.recorder = &MockBotServiceMockRecorder{mock}
	return mock
}

func (m *MockBotService) EXPECT() *MockBotServiceMockRecorder {
	return m.recorder
}

func (m *MockBotService) CreateReminder(msg *tgbotapi.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateReminder", msg)
}

func (mr *MockBotServiceMockRecorder) CreateReminder(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReminder", reflect.TypeOf((*MockBotService)(nil).CreateReminder), msg)
}

func (m *MockBotService) UpdateReminder(msg *tgbotapi.Message) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReminder", msg)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (mr *MockBotServiceMockRecorder) UpdateReminder(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReminder", reflect.TypeOf((*MockBotService)(nil).UpdateReminder), msg)
}

func (m *MockBotService) HandleCallbackQuery(callback *tgbotapi.CallbackQuery) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleCallbackQuery", callback)
}

func (mr *MockBotServiceMockRecorder) HandleCallbackQuery(callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleCallbackQuery", reflect.TypeOf((*MockBotService)(nil).HandleCallbackQuery), callback)
}

func (m *MockBotService) HandleCommand(message *tgbotapi.Message, task *tgbotapi.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleCommand", message, task)
}

func (mr *MockBotServiceMockRecorder) HandleCommand(message, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleCommand", reflect.TypeOf((*MockBotService)(nil).HandleCommand), message, task)
}

func (m *MockBotService) HandleMyChatMemberUpdate(myChatMember *tgbotapi.ChatMemberUpdated) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleMyChatMemberUpdate", myChatMember)
}

func (mr *MockBotServiceMockRecorder) HandleMyChatMemberUpdate(myChatMember interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMyChatMemberUpdate", reflect.TypeOf((*MockBotService)(nil).HandleMyChatMemberUpdate), myChatMember)
}

func (m *MockBotService) HandleStart(chat *tgbotapi.Chat) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleStart", chat)
}

func (mr *MockBotServiceMockRecorder) HandleStart(chat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleStart", reflect.TypeOf((*MockBotService)(nil).HandleStart), chat)
}
