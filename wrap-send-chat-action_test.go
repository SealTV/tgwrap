package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendChatAction(t *testing.T) {
	token, err := getTokenEnv()
	assert.Nil(t, err, "Token")
	if len(token) < 1 {
		return
	}

	bot := NewBot(token)

	arr, err2 := bot.GetUpdates(&GetUpdatesOpt{Limit: 1})
	assert.Nil(t, err2, "GetUpdates err")

	if len(arr) < 1 {
		return
	}

	up := arr[0]
	chatID := up.Message.From.ID

	b2, err4 := bot.SendChatAction(chatID, ChatActionUploadPhoto)
	assert.Nil(t, err4, "SendChatAction err")
	assert.True(t, b2, "ChatActionTyping")
}
