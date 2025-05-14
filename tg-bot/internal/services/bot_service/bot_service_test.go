package bot_service

/*
Unit tests не сработают, пока вы не измените зависимость к BotImpl, вместо этого, используйте интерфейс bot.Bot.
Сейчас в коде используется приведение типов, которое не сможет работать с мок-объектом. (deleteBotMessage нужно привести MockBot к *bot.BotImpl).
*/
import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBot is a mock implementation of the bot.Bot interface.
type MockBot struct {
	mock.Mock
}

func (m *MockBot) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	args := m.Called(c)
	return args.Get(0).(tgbotapi.Message), args.Error(1)
}

func (m *MockBot) Request(c tgbotapi.Chattable) (*tgbotapi.APIResponse, error) {
	args := m.Called(c)
	return args.Get(0).(*tgbotapi.APIResponse), args.Error(1)
}

func (m *MockBot) GetUpdatesChan(u tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel {
	args := m.Called(u)
	return args.Get(0).(tgbotapi.UpdatesChannel)
}

func (m *MockBot) GetMe() (tgbotapi.User, error) {
	args := m.Called()
	return args.Get(0).(tgbotapi.User), args.Error(1)
}

func TestCreateReminder(t *testing.T) {
	mockBot := new(MockBot)
	botService := NewBotService(mockBot)
	msg := &tgbotapi.Message{
		From: &tgbotapi.User{
			ID:       123,
			UserName: "testuser",
		},
		Chat: &tgbotapi.Chat{
			ID: 456,
		},
	}

	mockBot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)

	botService.CreateReminder(msg)

	rmdr, exists := reminders[msg.From.ID]
	assert.True(t, exists)
	assert.Equal(t, msg.From.ID, rmdr.GetUserID())
}

func TestUpdateReminder_Task(t *testing.T) {
	mockBot := new(MockBot)
	botService := NewBotService(mockBot)
	msg := &tgbotapi.Message{
		From: &tgbotapi.User{
			ID:       123,
			UserName: "testuser",
		},
		Chat: &tgbotapi.Chat{
			ID: 456,
		},
		Text: "Test task",
	}

	mockBot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)

	botService.CreateReminder(msg)
	result := botService.UpdateReminder(msg)

	rmdr, exists := reminders[msg.From.ID]
	assert.True(t, exists)
	assert.True(t, result)
	assert.Equal(t, "Test task", rmdr.GetTask().Text)
}

func TestUpdateReminder_Interval(t *testing.T) {
	mockBot := new(MockBot)
	botService := NewBotService(mockBot)
	msg := &tgbotapi.Message{
		From: &tgbotapi.User{
			ID:       123,
			UserName: "testuser",
		},
		Chat: &tgbotapi.Chat{
			ID: 456,
		},
		Text: "Test task",
	}

	mockBot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)

	botService.CreateReminder(msg)
	botService.UpdateReminder(msg)

	intervalMsg := &tgbotapi.Message{
		From: &tgbotapi.User{
			ID:       123,
			UserName: "testuser",
		},
		Chat: &tgbotapi.Chat{
			ID: 456,
		},
		Text: "5",
	}

	result := botService.UpdateReminder(intervalMsg)

	rmdr, exists := reminders[intervalMsg.From.ID]
	assert.True(t, exists)
	assert.True(t, result)
	assert.Equal(t, "5", rmdr.GetInterval())
}

func TestHandleCommand(t *testing.T) {
	mockBot := new(MockBot)
	botService := NewBotService(mockBot)
	msg := &tgbotapi.Message{
		From: &tgbotapi.User{
			ID:       123,
			UserName: "testuser",
		},
		Chat: &tgbotapi.Chat{
			ID: 456,
		},
		Text: "@bot ctrl 5d",
	}

	task := &tgbotapi.Message{
		From: &tgbotapi.User{
			ID:       123,
			UserName: "testuser",
		},
		Chat: &tgbotapi.Chat{
			ID: 456,
		},
		Text: "Test task",
	}

	mockBot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)

	botService.HandleCommand(msg, task)

	// Add more assertions based on the expected behavior
}

func TestHandleCallbackQuery(t *testing.T) {
	mockBot := new(MockBot)
	botService := NewBotService(mockBot)
	msg := &tgbotapi.Message{
		From: &tgbotapi.User{
			ID:       123,
			UserName: "testuser",
		},
		Chat: &tgbotapi.Chat{
			ID: 456,
		},
		Text: "Test task",
	}

	mockBot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)

	botService.CreateReminder(msg)
	botService.UpdateReminder(msg)

	callback := &tgbotapi.CallbackQuery{
		From: &tgbotapi.User{
			ID:       123,
			UserName: "testuser",
		},
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				ID: 456,
			},
		},
		Data: "d",
	}

	botService.HandleCallbackQuery(callback)

	// Add more assertions based on the expected behavior
}
