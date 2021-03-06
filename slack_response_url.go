package slack

import (
	"context"
	"encoding/json"
)

// The implementation has initially be taken from https://github.com/nlopes/slack/issues/387#issuecomment-480961473
// author: gtrindade

// DeleteEphemeral deletes the ephemeral message, use the responseURL from the action payload in your ephemeral message. Only works for the interactive ones.
func (api *Client) DeleteEphemeral(responseURL string) (*SlackResponse, error) {
	message := Msg{
		ResponseType:    "ephemeral",
		Text:            "",
		ReplaceOriginal: true,
		DeleteOriginal:  true,
	}
	return api.SendResponse(responseURL, message)
}

// SendResponse Will send a json response marshalled from Msg as string using the responseURL as endpoint.
func (api *Client) SendResponse(responseURL string, message Msg) (*SlackResponse, error) {
	payload, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	json := []byte(payload)
	response := &SlackResponse{}
	err = postJSON(context.Background(), api.httpclient, responseURL, api.token, json, response, api)
	return response, err
}
