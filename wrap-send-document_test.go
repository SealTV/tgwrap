package tgwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendDocument(t *testing.T) {
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

	inputFile := NewInputFileLocal("wrap-send-document.go")

	_, err4 := bot.SendDocument(chatID, inputFile,
		&SendDocumentOpt{
			DisableNotification: true,
		})

	assert.Nil(t, err4, "SendMessage err")
}
