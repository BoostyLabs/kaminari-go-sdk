package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type client struct {
	apiUrl      string
	restyClient *resty.Client
}

func newKaminariClient(apiUrl string, restyClient *resty.Client) *client {
	return &client{
		apiUrl:      apiUrl,
		restyClient: restyClient,
	}
}

func (c *client) onChainURL() string {
	return fmt.Sprintf("%v/api/bitcoin/v1", c.apiUrl)
}

func (c *client) lightningURL() string {
	return fmt.Sprintf("%v/api/lightning/v1", c.apiUrl)
}
