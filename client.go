package greynoise

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
)

type Client struct {
	Key string

	*http.Client
	baseURL *url.URL
}

// Context will return more information about a given IP address. Returns time ranges, IP metadata (network owner,
// ASN,reverse DNS pointer, country), associated actors, activity tags, and raw port scan and web
// request information.
func (c *Client) Context(ip net.IP) (*ContextResult, error) {
	request, err := c.NewRequest("GET", fmt.Sprintf("/v2/noise/context/%s", ip.String()), nil)
	if err != nil {
		return nil, err
	}

	output := ContextResult{}
	if err := c.Do(request, &output); err != nil {
		return nil, err
	}

	return &output, nil
}

type optionFn func(*Client)

// WithKey contains the Greynoise API key
func WithKey(key string) optionFn {
	return func(c *Client) {
		c.Key = key
	}
}

// New returns a Greynoise API client
func New(options ...optionFn) (*Client, error) {
	baseURL, err := url.Parse("https://enterprise.api.greynoise.io/v2/noise/")
	if err != nil {
		panic(err)
	}

	c := &Client{
		baseURL: baseURL,
		Client:  http.DefaultClient,
	}

	for _, optionFn := range options {
		optionFn(c)

	}

	return c, nil
}
