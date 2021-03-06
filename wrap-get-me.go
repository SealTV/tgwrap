package tgwrap

//
// GetMe is used for testing your bot's auth token.
// Requires no parameters. Returns basic information
// about the bot in form of a User object.Telegram API method
//
func (p *bot) GetMe() (*User, error) {

	var resp struct {
		GenericResponse

		Result *User `json:"result"`
	}

	err := p.getAPIResponse("getMe", p.sendJSON, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
