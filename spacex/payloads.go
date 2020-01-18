package spacex

import (
	"fmt"
)

type PayloadsService service

type OrbitParams struct {
	ReferenceSystem string  `json:"reference_system"`
	Regime          string  `json:"regime"`
	Longitude       int     `json:"longitude"`
	SemiMajorAxisKm float32 `json:"semi_major_axis_km"`
	Eccentricity    float32 `json:"eccentricity"`
	PeriapsisKm     float32 `json:"periapsis_km"`
	ApoapsisKm      float32 `json:"apoapsis_km"`
	InclinationDeg  float32 `json:"inclination_deg"`
	PeriodMin       float32 `json:"period_min"`
	LifespanYears   int     `json:"lifespan_years"`
	Epoch           string  `json:"epoch"`
	MeanMotion      float32 `json:"mean_motion"`
	Raan            float32 `json:"raan"`
	ArgOfPericenter float32 `json:"arg_of_pericenter"`
	MeanAnomaly     float32 `json:"mean_anomaly"`
}

type Payload struct {
	PayloadID      string      `json:"payload_id"`
	NoradID        []int       `json:"norad_id"`
	Reused         bool        `json:"reused"`
	Customers      []string    `json:"customers"`
	Nationality    string      `json:"nationality"`
	Manufacturer   string      `json:"manufacturer"`
	PayloadType    string      `json:"payload_type"`
	PayloadMassKg  int         `json:"payload_mass_kg"`
	PayloadMassLbs int         `json:"payload_mass_lbs"`
	Orbit          string      `json:"orbit"`
	OrbitParams    OrbitParams `json:"orbit_params"`
}

func (s *PayloadsService) Get(payloadID string) (*Dragon, error) {
	if payloadID == "" {
		return nil, ErrInvalidID
	}

	u := fmt.Sprintf("payloads/%v", payloadID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	c := new(Dragon)
	err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *PayloadsService) ListAll(opt *LaunchesListOptions) ([]*Payload, error) {
	u := "payloads"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var c []*Payload
	err = s.client.Do(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
