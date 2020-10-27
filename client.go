package stravaclient

import (
	"encoding/json"
	"fmt"
	"io"
)

type api interface {
	GetAthlete(accessToken string) (io.Reader, error)
}

type Client struct {
	AccessToken string
	Api         api
}

func (c *Client) ShowAthlete() (*Athlete, error) {
	responseBody, err := c.Api.GetAthlete(c.AccessToken)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var athlete Athlete
	jsonErr := json.NewDecoder(responseBody).Decode(&athlete)
	return &athlete, jsonErr
}
