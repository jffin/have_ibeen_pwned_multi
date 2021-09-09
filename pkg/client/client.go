package client

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// RLHTTPClient Rate Limited HTTP Client
type RLHTTPClient struct {
	client      *http.Client
	RateLimiter *rate.Limiter
}

// Do dispatches the HTTP request to the network
func (c *RLHTTPClient) Do(req *http.Request) (*http.Response, error) {
	// Comment out the below 5 lines to turn off rateLimiting
	ctx := context.Background()
	err := c.RateLimiter.Wait(ctx) // This is a blocking call. Honors the rate limit
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// NewClient return http client with a rateLimiter
func NewClient(rl *rate.Limiter) *RLHTTPClient {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	c := &RLHTTPClient{
		client:      client,
		RateLimiter: rl,
	}
	return c
}
