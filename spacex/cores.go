package spacex

import (
	"fmt"
)

type CoresService service

type Core struct {
	CoreSerial 			string 		`json:"core_serial"`
	Block 				int 		`json:"block"`
	Status 				string 		`json:"status"`
	OriginalLaunch 		string 		`json:"original_launch"`
	OriginalLaunchUnix 	int 		`json:"original_launch_unix"`
	Missions			[]Mission	`json:"missions"`
	ReuseCount 			int 		`json:"reuse_count"`
	RtlsAttempts 		int 		`json:"rtls_attempts"`
	RtlsLandings 		int 		`json:"rtls_landings"`
	AsdsAttempts 		int 		`json:"asds_attempts"`
	AsdsLandings 		int 		`json:"asds_landings"`
	WaterLanding 		bool 		`json:"water_landing"`
	Details 			string 		`json:"details"`
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

	c:= new(Core)
	err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CoresService) ListAll() ([]*Core, error) {
	u := "cores"
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

func (s *CoresService) ListUpcoming() ([]*Core, error) {
	u := "cores/upcoming"
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

func (s *CoresService) ListPast() ([]*Core, error) {
	u := "cores/past"
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