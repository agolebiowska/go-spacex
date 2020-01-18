package spacex

import (
	"fmt"
)

type LandingPadsService service

type Location struct {
	Name      string  `json:"name"`
	Region    string  `json:"region"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

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

type LandingPadsListOptions struct {
	ID     bool `url:"id,omitempty"`     // Set as true to show mongo document id's
	Limit  int  `url:"limit,omitempty"`  // Limit results returned, defaults to all documents returned
	Offset int  `url:"offset,omitempty"` // Offset or skip results from the beginning of the query
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

func (s *LandingPadsService) ListAll(opt *LandingPadsListOptions) ([]*LandingPad, error) {
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
