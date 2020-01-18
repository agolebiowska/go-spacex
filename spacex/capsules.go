package spacex

import (
	"fmt"
	"time"
)

type CapsulesService service

type Capsule struct {
	CapsuleSerial      string             `json:"capsule_serial"`
	CapsuleID          string             `json:"capsule_id"`
	Status             string             `json:"status"`
	OriginalLaunch     string             `json:"original_launch"`
	OriginalLaunchUnix int                `json:"original_launch_unix"`
	Missions           []MinimisedMission `json:"missions"`
	Landings           int                `json:"landings"`
	Type               int                `json:"type"`
	Details            string             `json:"details"`
	ReuseCount         int                `json:"reuse_count"`
}

type CapsuleListOptions struct {
	CapsuleSerial  string    `url:"capsule_serial,omitempty"`
	CapsuleID      string    `url:"capsule_id,omitempty"`
	Status         string    `url:"status,omitempty"`
	OriginalLaunch time.Time `url:"original_launch,omitempty"`
	Mission        string    `url:"mission,omitempty"`
	Landings       int       `url:"landings,omitempty"`
	Type           string    `url:"type,omitempty"`
	ReuseCount     int       `url:"reuse_count,omitempty"`
}

func (s *CapsulesService) Get(serial string) (*Capsule, error) {
	if serial == "" {
		return nil, ErrInvalidID
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

func (s *CapsulesService) ListAll(opt *CapsuleListOptions) ([]*Capsule, error) {
	u := "capsules"
	return s.list(u, opt)
}

func (s *CapsulesService) ListUpcoming(opt *CapsuleListOptions) ([]*Capsule, error) {
	u := "capsules/upcoming"
	return s.list(u, opt)
}

func (s *CapsulesService) ListPast(opt *CapsuleListOptions) ([]*Capsule, error) {
	u := "capsules/past"
	return s.list(u, opt)
}

func (s *CapsulesService) list(u string, opt *CapsuleListOptions) ([]*Capsule, error) {
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

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
