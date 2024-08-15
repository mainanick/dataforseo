package dataforseo

import (
	"net/http"
)

func redirectPolicyFunc(req *http.Request, _ []*http.Request) error {
	req.SetBasicAuth(Username, Password)
	return nil
}

type Client struct {
	HTTPClent         *http.Client
	MaxNetworkRetries int
}

func NewClient() *Client {
	return &Client{
		HTTPClent:         http.DefaultClient,
		MaxNetworkRetries: 4,
	}
}
