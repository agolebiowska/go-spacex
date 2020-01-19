package spacex

import (
	"fmt"
)

type LaunchPadsService service

type LaunchPad struct {
	ID                 int      `json:"id"`
	Name               string   `json:"name"`
	Status             string   `json:"status"`
	Location           Location `json:"location"`
	VehiclesLaunched   []string `json:"vehicles_launched"`
	AttemptedLaunches  int      `json:"attempted_launches"`
	SuccessfulLaunches int      `json:"successful_launches"`
	Wikipedia          string   `json:"wikipedia"`
	Details            string   `json:"details"`
	SiteId             string   `json:"site_id"`
	SiteNameLong       string   `json:"site_name_long"`
}

type LaunchPadsListOptions struct {
}

func (s *LaunchPadsService) Get(siteID string) (*LaunchPad, error) {
	if siteID == "" {
		return nil, ErrInvalidID
	}

	u := fmt.Sprintf("launchpads/%v", siteID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	c := new(LaunchPad)
	err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *LaunchPadsService) ListAll(baseOpt *ListOptions, extOpt *LaunchPadsListOptions) ([]*LaunchPad, error) {
	u := "launchpads"
	u, err := addOptions(u, baseOpt, extOpt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var c []*LaunchPad
	err = s.client.Do(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
