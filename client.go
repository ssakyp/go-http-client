package client

import "net/http"

var DefaultAPIURL = "https://pokeapi.co"

type Client struct {
	apiURL string
	httpClient *http.Client
}


type Option func(c *Client)


// bootstrapped
func NewClient(opts ...Option) *Client {
	client := &Client{
		apiURL: DefaultAPIURL,
		httpClient: http.DefaultClient,
	}

	for _, option := range opts {
		option(client)
	}

	return client
}

func WithAPIURL(url string) Option {
	return func(c *Client) {
		c.apiURL = url
	}
}

func WithHTTPClient(httpClient *http.Client) Option{
	return func(c *Client){
		c.httpClient = httpClient
	}
}