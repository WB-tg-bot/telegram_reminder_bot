package handlers

import (
	"testing"
	"tg-bot/internal/services/mocks"
	"tg-bot/internal/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandlerImpl_HandleUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockBotService(ctrl)
	handler := NewHandler()

	t.Run("HandleMessage", func(t *testing.T) {
		update := tgbotapi.Update{Message: &tgbotapi.Message{}}
		handler.HandleUpdate(mockService, update)
	})

	t.Run("HandleCallbackQuery", func(t *testing.T) {
		update := tgbotapi.Update{CallbackQuery: &tgbotapi.CallbackQuery{}}
		handler.HandleUpdate(mockService, update)
	})

	t.Run("HandleMyChatMemberUpdate", func(t *testing.T) {
		update := tgbotapi.Update{MyChatMember: &tgbotapi.ChatMemberUpdated{}}
		handler.HandleUpdate(mockService, update)
	})

	t.Run("HandleEditedMessage", func(t *testing.T) {
		update := tgbotapi.Update{EditedMessage: &tgbotapi.Message{}}
		handler.HandleUpdate(mockService, update)
	})
}

func TestHandlerImpl_HandleMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockBotService(ctrl)
	handler := NewHandler()

	t.Run("StartCommand", func(t *testing.T) {
		message := &tgbotapi.Message{Text: "/start", Chat: &tgbotapi.Chat{}, From: &tgbotapi.User{ID: 1}}
		mockService.EXPECT().HandleStart(message.Chat)
		utils.DeleteMessage = func(api *tgbotapi.BotAPI, message *tgbotapi.Message) {}
		handler.HandleMessage(mockService, message)
	})

	t.Run("AddReminder", func(t *testing.T) {
		message := &tgbotapi.Message{Text: "Добавить напоминание", From: &tgbotapi.User{ID: 1}}
		mockService.EXPECT().CreateReminder(message).Times(1)
		utils.DeleteMessage = func(api *tgbotapi.BotAPI, message *tgbotapi.Message) {}
		handler.HandleMessage(mockService, message)
		assert.True(t, flags[message.From.ID])
	})

	t.Run("CommandMatch", func(t *testing.T) {
		message := &tgbotapi.Message{Text: "@username ctrl 123a", From: &tgbotapi.User{ID: 1}}
		mockService.EXPECT().HandleCommand(message, gomock.Any()).Times(1)
		handler.HandleMessage(mockService, message)
	})

	t.Run("UpdateReminder", func(t *testing.T) {
		message := &tgbotapi.Message{Text: "Some text", From: &tgbotapi.User{ID: 1}}
		flags[message.From.ID] = true
		mockService.EXPECT().UpdateReminder(message).Return(false).Times(1)
		handler.HandleMessage(mockService, message)
		assert.False(t, flags[message.From.ID])
	})

	t.Run("SaveMessage", func(t *testing.T) {
		message := &tgbotapi.Message{Text: "Some text", From: &tgbotapi.User{ID: 1}}
		handler.HandleMessage(mockService, message)
		assert.Equal(t, message, msgs[message.From.ID])
	})
}

func TestHandlerImpl_HandleCallbackQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockBotService(ctrl)
	handler := NewHandler()

	callback := &tgbotapi.CallbackQuery{From: &tgbotapi.User{ID: 1}}
	mockService.EXPECT().HandleCallbackQuery(callback).Times(1)
	handler.HandleCallbackQuery(mockService, callback)
	assert.False(t, flags[callback.From.ID])
}

func TestHandlerImpl_HandleMyChatMemberUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockBotService(ctrl)
	handler := NewHandler()

	myChatMember := &tgbotapi.ChatMemberUpdated{}
	mockService.EXPECT().HandleMyChatMemberUpdate(myChatMember).Times(1)
	handler.HandleMyChatMemberUpdate(mockService, myChatMember)
}

func TestHandlerImpl_HandleEditedMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockBotService(ctrl)
	handler := NewHandler()

	editedMessage := &tgbotapi.Message{From: &tgbotapi.User{ID: 1}}
	handler.HandleEditedMessage(mockService, editedMessage)
	assert.Equal(t, editedMessage, msgs[editedMessage.From.ID])
}
