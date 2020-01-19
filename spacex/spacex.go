package spacex

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

const (
	defaultBaseURL = "https://api.spacexdata.com/v3/"
)

type service struct {
	client *Client
}

type Client struct {
	http    *http.Client
	BaseURL *url.URL

	common service

	// Services
	Capsules         *CapsulesService
	Cores            *CoresService
	Dragons          *DragonsService
	HistoricalEvents *HistoricalEventsService
	Info             *InfoService
	LandingPads      *LandingPadsService
	Launches         *LaunchesService
	LaunchPads       *LaunchPadsService
	Missions         *MissionsService
	Payloads         *PayloadsService
	Rockets          *RocketsService
	Ships            *ShipsService
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
	c.Info = (*InfoService)(&c.common)
	c.LandingPads = (*LandingPadsService)(&c.common)
	c.Launches = (*LaunchesService)(&c.common)
	c.LaunchPads = (*LaunchPadsService)(&c.common)
	c.Missions = (*MissionsService)(&c.common)
	c.Payloads = (*PayloadsService)(&c.common)
	c.Rockets = (*RocketsService)(&c.common)
	c.Ships = (*ShipsService)(&c.common)
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

func addOptions(s string, opt ...interface{}) (string, error) {
	var u, err = url.Parse(s)
	if err != nil {
		return s, err
	}

	for _, o := range opt {
		v := reflect.ValueOf(o)
		if !v.IsNil() {
			qs, err := query.Values(o)
			if err != nil {
				return s, err
			}

			u.RawQuery += qs.Encode()
		}
	}

	return u.String(), nil
}
