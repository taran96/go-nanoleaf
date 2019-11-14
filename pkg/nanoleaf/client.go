package nanoleaf

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type Client interface {
	Authorize(ctx context.Context) error
	GetInfo(ctx context.Context) (*DeviceInfo, error)
	SetState(ctx context.Context, state *State) error
	GetPower(ctx context.Context) (*OnValue, error)
	SetPower(ctx context.Context, value bool) error
	GetBrightness(ctx context.Context) (*RangedValue, error)
	SetBrightness(ctx context.Context, value int) error
	SetBrightnessWithDuration(ctx context.Context, value int, duration int) error
	GetHue(ctx context.Context) (*RangedValue, error)
	SetHue(ctx context.Context, value int) error
	GetSaturation(ctx context.Context) (*RangedValue, error)
	SetSaturation(ctx context.Context, value int) error
	GetColorTemperature(ctx context.Context) (*RangedValue, error)
	SetColorTemperature(ctx context.Context, value int) error
	GetColorMode(ctx context.Context) (*string, error)
	GetCurrentEffect(ctx context.Context) (*string, error)
	GetAllEffects(ctx context.Context) ([]*string, error)
	SelectEffect(ctx context.Context, effect string) error
}

type client struct {
	BaseURL    *url.URL
	Token      *string
	httpClient *http.Client
}

func NewClient(httpClient *http.Client, deviceIp string) (Client, error) {
	baseURL, err := url.Parse("http://" + deviceIp + "/api/v1/")
	if err != nil {
		return nil, err
	}
	return &client{
		BaseURL:    baseURL,
		Token:      nil,
		httpClient: httpClient,
	}, nil
}

func NewClientWithToken(httpClient *http.Client, deviceIp string, token string) (Client, error) {
	baseURL, err := url.Parse("http://" + deviceIp + "/api/v1/" + token + "/")
	if err != nil {
		return nil, err
	}
	return &client{
		BaseURL:    baseURL,
		Token:      &token,
		httpClient: httpClient,
	}, nil
}

func (c *client) newRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	rel := url.URL{Path: path}
	url := c.BaseURL.ResolveReference(&rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}

func (c *client) makeRequest(ctx context.Context, method, path string, data interface{}, resp interface{}) (*http.Response, error) {
	req, err := c.newRequest(ctx, method, path, data)

	if err != nil {
		return nil, err
	}

	httpResp, err := c.do(req, resp)

	switch httpResp.StatusCode {
	case 404:
		err = errors.New("NOT FOUND")
	case 500:
		err = errors.New("SERVER ERROR")
	}

	return httpResp, err
}

func (c *client) makeRequestRaw(ctx context.Context, method, path string, data interface{}) (*http.Response, error) {
	req, err := c.newRequest(ctx, method, path, data)

	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)

	switch resp.StatusCode {
	case 404:
		err = errors.New("NOT FOUND")
	case 500:
		err = errors.New("SERVER ERROR")
	}

	return resp, err
}
