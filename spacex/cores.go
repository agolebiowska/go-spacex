package spacex

import (
	"fmt"
	"time"
)

type CoresService service

type Core struct {
	CoreSerial         string    `json:"core_serial"`
	Block              int       `json:"block"`
	Status             string    `json:"status"`
	OriginalLaunch     string    `json:"original_launch"`
	OriginalLaunchUnix int       `json:"original_launch_unix"`
	Missions           []Mission `json:"missions"`
	ReuseCount         int       `json:"reuse_count"`
	RtlsAttempts       int       `json:"rtls_attempts"`
	RtlsLandings       int       `json:"rtls_landings"`
	AsdsAttempts       int       `json:"asds_attempts"`
	AsdsLandings       int       `json:"asds_landings"`
	WaterLanding       bool      `json:"water_landing"`
	Details            string    `json:"details"`
}

type CoresListOptions struct {
	CapsuleSerial  string    `url:"core_serial,omitempty"`
	CapsuleID      int       `url:"block,omitempty"`
	Status         string    `url:"status,omitempty"`
	OriginalLaunch time.Time `url:"original_launch,omitempty"`
	Mission        string    `url:"mission,omitempty"`
	ReuseCount     int       `url:"reuse_count,omitempty"`
	RtlsAttempts   int       `url:"rtls_attempts,omitempty"`
	RtlsLandings   int       `url:"rtls_landings,omitempty"`
	AsdsAttempts   int       `url:"asds_attempts,omitempty"`
	AsdsLandings   int       `url:"asds_landings,omitempty"`
	WaterLanding   bool      `url:"water_landing,omitempty"`
}

func (s *CoresService) Get(serial string) (*Core, error) {
	if serial == "" {
		return nil, ErrInvalidSerial
	}

	u := fmt.Sprintf("cores/%v", serial)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	c := new(Core)
	err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CoresService) ListAll(opt *CoresListOptions) ([]*Core, error) {
	u := "cores"
	return s.list(u, opt)
}

func (s *CoresService) ListUpcoming(opt *CoresListOptions) ([]*Core, error) {
	u := "cores/upcoming"
	return s.list(u, opt)
}

func (s *CoresService) ListPast(opt *CoresListOptions) ([]*Core, error) {
	u := "cores/past"
	return s.list(u, opt)
}

func (s *CoresService) list(u string, opt *CoresListOptions) ([]*Core, error) {
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var c []*Core
	err = s.client.Do(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
