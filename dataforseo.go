package dataforseo

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
)

var (
	DefaultBaseURL string = "https://api.dataforseo.com/v3/"
	UserAgent             = "github.com/mainanick/dataforseo"
)

type service struct {
	client *Client
}

type Client struct {
	clientMu          sync.Mutex // clientMu protects the client during calls that modify the CheckRedirect func.
	client            *http.Client
	MaxNetworkRetries int
	BaseURL           *url.URL
	UserAgent         string

	srv service

	Keyword *SiteKeywordService
	Serp    *SerpService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	c := &Client{
		client: httpClient,
	}
	c.init()
	return c
}

func (c *Client) init() {
	if c.client == nil {
		c.client = &http.Client{}
	}
	if c.BaseURL == nil {
		c.BaseURL, _ = url.Parse(DefaultBaseURL)
	}
	c.UserAgent = UserAgent

	c.srv.client = c
	c.Keyword = (*SiteKeywordService)(&c.srv)
	c.Serp = (*SerpService)(&c.srv)
}

// WithAuthToken returns a copy of the client configured to use the provided token for the Authorization header.
func (c *Client) WithAuthToken(username string, password string) *Client {
	c2 := c.copy()
	defer c2.init()
	transport := c2.client.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	c2.client.Transport = roundTripperFunc(
		func(req *http.Request) (*http.Response, error) {
			req = req.Clone(req.Context())
			req.SetBasicAuth(username, password)
			return transport.RoundTrip(req)
		},
	)
	return c2
}

// copy returns a copy of the current client. It must be initialized before use.
func (c *Client) copy() *Client {
	c.clientMu.Lock()
	// can't use *c here because that would copy mutexes by value.
	clone := Client{
		client:    &http.Client{},
		UserAgent: c.UserAgent,
		BaseURL:   c.BaseURL,
	}
	c.clientMu.Unlock()
	if c.client != nil {
		clone.client.Transport = c.client.Transport
		clone.client.CheckRedirect = c.client.CheckRedirect
		clone.client.Jar = c.client.Jar
		clone.client.Timeout = c.client.Timeout
	}
	return &clone
}

// RequestOption represents an option that can modify an http.Request.
type RequestOption func(req *http.Request)

func WithUserAgent(useragent string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set("User-Agent", useragent)
	}
}

func (c *Client) NewRequest(method, urlStr string, body interface{}, opts ...RequestOption) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	log.Println(u.String())
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	for _, opt := range opts {
		opt(req)
	}

	return req, nil
}

type DataForSEOError struct{}

func (*DataForSEOError) Error() string {
	return "DataForSEO Error"
}
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	return &DataForSEOError{}
}
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	err = CheckResponse(resp)
	if err != nil {
		defer resp.Body.Close()
		return resp, err
	}

	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}

		// file, _ := json.MarshalIndent(v, "", " ")
		// _ = os.WriteFile("local/test.json", file, 0644)
	}
	return resp, err
}

// roundTripperFunc creates a RoundTripper (transport)
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}
