package http_client

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
)

type Client struct {
	client  *http.Client
	baseUrl string
}

func NewClient(baseUrl string) *Client {
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := new(http.Client)
	client.Transport = transport
	return &Client{
		client:  client,
		baseUrl: baseUrl,
	}
}

func (c *Client) GetRawBody(ctx context.Context, path string, headers map[string]string, redirectFunc func(req *http.Request, via []*http.Request) error) ([]byte, int, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseUrl+path, nil)
	if err != nil {
		return nil, 0, err
	}
	//req.Header.Add("content-type", c.contentType)
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	var client http.Client
	client = *c.client
	client.CheckRedirect = redirectFunc
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return body, resp.StatusCode, nil
}
