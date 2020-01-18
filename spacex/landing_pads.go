package spacex

import (
	"fmt"
)

type LandingPadsService service

type LandingPad struct {
	ID                 string   `json:"id"`
	FullName           string   `json:"full_name"`
	Status             string   `json:"status"`
	Location           Location `json:"location"`
	LandingType        string   `json:"landing_type"`
	AttemptedLandings  int      `json:"attempted_landings"`
	SuccessfulLandings int      `json:"successful_landings"`
	Wikipedia          string   `json:"wikipedia"`
	Details            string   `json:"details"`
}

func (s *LandingPadsService) Get(ID string) (*LandingPad, error) {
	if ID == "" {
		return nil, ErrInvalidID
	}

	u := fmt.Sprintf("landpads/%v", ID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	c := new(LandingPad)
	err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *LandingPadsService) ListAll(opt *ListOptions) ([]*LandingPad, error) {
	u := "landpads"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var c []*LandingPad
	err = s.client.Do(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
