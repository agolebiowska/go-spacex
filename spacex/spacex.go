package spacex

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.spacexdata.com/v3/"
)

type service struct {
	client *Client
}

type Client struct {
	http *http.Client
	BaseURL *url.URL

	common service

	// Services
	Capsules *CapsulesService
	Cores *CoresService
	Dragons *DragonsService
	HistoricalEvents *HistoricalEventsService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{http: httpClient, BaseURL: baseURL}
	c.common.client = c
	c.Capsules = (*CapsulesService)(&c.common)
	c.Cores = (*CoresService)(&c.common)
	c.Dragons = (*DragonsService)(&c.common)
	c.HistoricalEvents = (*HistoricalEventsService)(&c.common)
	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
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

	return req, nil
}

func (c *Client) Do(req *http.Request, result interface{}) error {
	res, err := c.http.Do(req)
	if err != nil {
		return errors.Wrap(err, "http client cannot send request")
	}
	defer res.Body.Close()

	if err = checkResponse(res); err != nil {
		return err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "cannot read response body")
	}

	if isBracketPair(b) {
		return ErrNoResults
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		return errors.Wrap(ErrInvalidJSON, err.Error())
	}

	return nil
}