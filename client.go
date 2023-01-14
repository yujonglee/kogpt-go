package kogpt

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type HTTPError struct {
	Detail string
	Status int
}

func (h HTTPError) Error() string {
	return fmt.Sprintf("http error with status %d: %v", h.Status, h.Detail)
}

func (h HTTPError) GetDetail() string {
	return h.Detail
}

func (h HTTPError) GetStatusCode() int {
	return h.Status
}

var _ error = (*HTTPError)(nil)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

var _ HTTPClient = (*http.Client)(nil)

type Client struct {
	client  HTTPClient
	baseUrl string
	token   string
}

func NewClient(client HTTPClient, token string) *Client {
	baseUrl := "https://api.kakaobrain.com/v1/inference/kogpt/"

	return &Client{
		client:  client,
		baseUrl: baseUrl,
		token:   token,
	}
}

func (c *Client) issueRequest(ctx context.Context, endpoint string, params, dest interface{}) error {
	if c.client == nil {
		return errors.New("client is nil")
	}

	var buf io.Reader = nil
	if params != nil {
		j, err := json.Marshal(params)
		if err != nil {
			return err
		}
		buf = bytes.NewBuffer(j)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseUrl+endpoint, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "KakaoAK "+c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return &HTTPError{
			Detail: string(body),
			Status: resp.StatusCode,
		}
	}

	if err = json.Unmarshal(body, dest); err != nil {
		return err
	}

	return nil
}
