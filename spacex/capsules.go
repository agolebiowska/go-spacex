package spacex

import (
	"fmt"
)

type CapsulesService service

type Capsule struct {
	CapsuleSerial      string    `json:"capsule_serial"`
	CapsuleID          string    `json:"capsule_id"`
	Status             string    `json:"status"`
	OriginalLaunch     string    `json:"original_launch"`
	OriginalLaunchUnix int       `json:"original_launch_unix"`
	Missions           []Mission `json:"missions"`
	Landings           int       `json:"landings"`
	Type               int       `json:"type"`
	Details            string    `json:"details"`
	ReuseCount         int       `json:"reuse_count"`
}

func (s *CapsulesService) Get(serial string) (*Capsule, error) {
	if serial == "" {
		return nil, ErrInvalidSerial
	}

	u := fmt.Sprintf("capsules/%v", serial)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	c := new(Capsule)
	err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CapsulesService) ListAll() ([]*Capsule, error) {
	u := "capsules"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var c []*Capsule
	err = s.client.Do(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CapsulesService) ListUpcoming() ([]*Capsule, error) {
	u := "capsules/upcoming"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var c []*Capsule
	err = s.client.Do(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CapsulesService) ListPast() ([]*Capsule, error) {
	u := "capsules/past"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var c []*Capsule
	err = s.client.Do(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
