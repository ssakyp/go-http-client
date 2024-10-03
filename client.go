package client

import "net/http"

type Client struct {
	httpClient *http.Client
}

// bootstrapped
func NewClient() *Client {
	client := &Client{
		httpClient: *&http.DefaultClient,
	}
	return client
}