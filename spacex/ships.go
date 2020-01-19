package spacex

import (
	"fmt"
)

type ShipsService service

type Ship struct {
	ShipID             string             `json:"ship_id"`
	ShipName           string             `json:"ship_name"`
	ShipModel          *int               `json:"ship_model"`
	ShipType           string             `json:"ship_type"`
	Roles              []string           `json:"roles"`
	Active             bool               `json:"active"`
	Imo                int                `json:"imo"`
	Mmsi               int                `json:"mmsi"`
	Abs                int                `json:"abs"`
	Class              int                `json:"class"`
	WeightLbs          int                `json:"weight_lbs"`
	WeightKg           int                `json:"weight_kg"`
	YearBuilt          int                `json:"year_built"`
	HomePort           string             `json:"home_port"`
	Status             string             `json:"status"`
	SpeedKn            *int               `json:"speed_kn"`
	CourseDeg          *int               `json:"course_deg"`
	Position           Position           `json:"position"`
	SuccessfulLandings *int               `json:"successful_landings"`
	AttemptedLandings  *int               `json:"attempted_landings"`
	Missions           []MinimisedMission `json:"missions"`
	Url                string             `json:"url"`
	Image              string             `json:"image"`
}

type ShipsListOptions struct {
	Options            ListOptions
	ShipID             string  `url:"ship_id,omitempty"`
	ShipName           string  `url:"ship_name,omitempty"`
	ShipModel          string  `url:"ship_model,omitempty"`
	ShipType           string  `url:"ship_type,omitempty"`
	Role               string  `url:"role,omitempty"`
	Active             bool    `url:"active,omitempty"`
	Imo                int     `url:"imo,omitempty"`
	Mmsi               int     `url:"mmsi,omitempty"`
	Abs                int     `url:"abs,omitempty"`
	Class              int     `url:"class,omitempty"`
	WeightLbs          int     `url:"weight_lbs,omitempty"`
	WeightKg           int     `url:"weight_kg,omitempty"`
	YearBuilt          int     `url:"year_built,omitempty"`
	HomePort           string  `url:"home_port,omitempty"`
	Status             string  `url:"status,omitempty"`
	SpeedKn            int     `url:"speed_kn,omitempty"`
	CourseDeg          int     `url:"course_deg,omitempty"`
	Latitude           float32 `url:"latitude,omitempty"`
	Longitude          float32 `url:"longitude,omitempty"`
	SuccessfulLandings int     `url:"successful_landings,omitempty"`
	AttemptedLandings  int     `url:"attempted_landings,omitempty"`
	Mission            string  `url:"mission,omitempty"`
}

func (s *ShipsService) Get(shipID string) (*Ship, error) {
	if shipID == "" {
		return nil, ErrInvalidID
	}

	u := fmt.Sprintf("ships/%v", shipID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	c := new(Ship)
	err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *ShipsService) ListAll(baseOpt *ListOptions, extOpt *ShipsListOptions) ([]*Ship, error) {
	u := "ships"
	u, err := addOptions(u, baseOpt, extOpt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var c []*Ship
	err = s.client.Do(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
