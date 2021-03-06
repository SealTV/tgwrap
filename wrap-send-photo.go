package tgwrap

import (
	"fmt"
	"reflect"

	"github.com/rogozhka/tgwrap/internal/thestruct"
)

//
// SendPhotoOpt represents optional params for SendPhoto
//
type SendPhotoOpt struct {

	//
	// Photo caption (may also be used when resending photos by file_id),
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
// SendPhoto is used to send photos.
//
// chatID: (Integer or String) Unique identifier
// for the target chat or username of the target channel (in the format @channelusername)
//
// photo: (*InputFile or string) Photo to send. Pass a file_id as string to send
// a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a string
// for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data
// using &NewInputFileLocal("<file path>")
//
// opt: (can be nil) optional params
//
func (p *bot) SendPhoto(chatID interface{}, photo interface{}, opt *SendPhotoOpt) (*Message, error) {

	type sendFormat struct {
		ChatID string `json:"chat_id"`

		SendPhotoOpt `json:",omitempty"`

		//
		// Photo to send. Pass a file_id as String to send a photo that exists
		// on the Telegram servers (recommended), pass an HTTP URL as a String
		// for Telegram to get a photo from the Internet,
		// or upload a new photo using multipart/form-data.
		//
		// InputFile should have MarshalText interface
		//
		Photo interface{} `json:"photo" form:"file"`
	}

	dataSend := sendFormat{
		ChatID: fmt.Sprint(chatID), // don't care about checking fmt, Telegram will response with error if invalid ID
		Photo:  photo,
	}

	if opt != nil {
		dataSend.SendPhotoOpt = *opt
	}

	var resp struct {
		GenericResponse

		Result *Message `json:"result"`
	}

	var sender fCommandSender = p.sendJSON

	tt := thestruct.Type(reflect.TypeOf(photo))
	if "InputFile" == tt.Name() && len(photo.(*InputFile).Name()) > 0 {
		sender = p.sendFormData
	}

	err := p.getAPIResponse("sendPhoto", sender, dataSend, &resp)
	if err != nil {
		return nil, fmt.Errorf("getAPIResponse ERROR:%v", err)
	}

	return resp.Result, nil
}
