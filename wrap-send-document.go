package tgwrap

import (
	"fmt"
	"reflect"

	"github.com/rogozhka/tgwrap/internal/thestruct"
)

//
// SendDocumentOpt represents optional params for SendDocument
//
type SendDocumentOpt struct {

	//
	// Document caption, 0-200 characters
	// 0-200 characters
	//
	Caption string `json:"caption,omitempty"`

	//
	// Sends the message silently. Users will receive a notification with no sound.
	//
	DisableNotification bool `json:"disable_notification,omitempty"`

	//
	// If the message is a reply, ID of the original message
	//
	ReplyToID uint64 `json:"reply_to_message_id,omitempty"`

	//
	// Additional interface options. A JSON-serialized object
	// for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard
	// or to force a reply from the user.
	//
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

//
// SendDocument to send general files. On success, the sent Message is returned.
// Bots can currently send files of any type of up to 50 MB in size, t
// his limit may be changed in the future.
//
// chatID: (uint64 or string) Unique identifier for the target chat
// or username of the target channel (in the format @channelusername)
//
// document: (*InputFile or string) File to send. Pass a file_id as String to send
// an audio that exists on the Telegram servers (recommended), pass an HTTP URL as a String
// for Telegram to get an audio from the Internet, or upload a new file using multipart/form-data.
// using &NewInputFileLocal("<file path>")
//
// opt: (can be nil) optional params
//
func (p *bot) SendDocument(chatID interface{}, document interface{}, opt *SendDocumentOpt) (*Message, error) {

	type sendFormat struct {
		ChatID string `json:"chat_id"`

		SendDocumentOpt `json:",omitempty"`

		Document interface{} `json:"document" form:"file"`
	}

	dataSend := sendFormat{
		ChatID:   fmt.Sprint(chatID),
		Document: document,
	}

	if opt != nil {
		dataSend.SendDocumentOpt = *opt
	}

	var resp struct {
		GenericResponse

		Result *Message `json:"result"`
	}

	var sender fCommandSender = p.sendJSON

	tt := thestruct.Type(reflect.TypeOf(document))
	if "InputFile" == tt.Name() && len(document.(*InputFile).Name()) > 0 {
		sender = p.sendFormData
	}

	err := p.getAPIResponse("sendDocument", sender, dataSend, &resp)
	if err != nil {
		return nil, fmt.Errorf("getAPIResponse ERROR:%v", err)
	}

	return resp.Result, nil
}
