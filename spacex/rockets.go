package spacex

import (
	"fmt"
)

type RocketsService service

type MinimisedRocket struct {
	RocketID    string      `json:"rocket_id"`
	RocketName  string      `json:"rocket_name"`
	RocketType  string      `json:"rocket_type"`
	FirstStage  FirstStage  `json:"first_stage"`
	SecondStage SecondStage `json:"second_stage"`
	Fairings    Fairings    `json:"fairings"`
}

type PayloadWeight struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Kg   int    `json:"kg"`
	Lb   int    `json:"lb"`
}

type RocketFirstStage struct {
	Reusable       bool    `json:"reusable"`
	Engines        int     `json:"engines"`
	FuelAmountTons float64 `json:"fuel_amount_tons"`
	BurnTimeSec    int     `json:"burn_time_sec"`
	ThrustSeaLevel Thrust  `json:"thrust_sea_level"`
	ThrustVacuum   Thrust  `json:"thrust_vacuum"`
}

type CompositeFairing struct {
	Height   Height   `json:"height"`
	Diameter Diameter `json:"diameter"`
}

type Payloads struct {
	Option1          string           `json:"option_1"`
	Option2          string           `json:"option_2"`
	CompositeFairing CompositeFairing `json:"composite_fairing"`
}

type RocketSecondStage struct {
	Reusable       bool     `json:"reusable"`
	Engines        int      `json:"engines"`
	FuelAmountTons float64  `json:"fuel_amount_tons"`
	BurnTimeSec    int      `json:"burn_time_sec"`
	Thrust         Thrust   `json:"thrust"`
	Payloads       Payloads `json:"payloads"`
}

type LandingLegs struct {
	Number   int     `json:"number"`
	Material *string `json:"material"`
}

type Isp struct {
	SeaLevel int `json:"sea_level"`
	Vacuum   int `json:"vacuum"`
}

type Engines struct {
	Number         int     `json:"number"`
	Type           string  `json:"type"`
	Version        string  `json:"version"`
	Layout         string  `json:"layout"`
	Isp            Isp     `json:"isp"`
	EngineLossMax  int     `json:"engine_loss_max"`
	Propellant1    string  `json:"propellant_1"`
	Propellant2    string  `json:"propellant_2"`
	ThrustSeaLevel Thrust  `json:"thrust_sea_level"`
	ThrustVacuum   Thrust  `json:"thrust_vacuum"`
	ThrustToWeight float32 `json:"thrust_to_weight"`
}

type Rocket struct {
	ID             int               `json:"id"`
	Active         bool              `json:"active"`
	Stages         int               `json:"stages"`
	Boosters       int               `json:"boosters"`
	CostPerLaunch  int               `json:"cost_per_launch"`
	SuccessRatePct int               `json:"success_rate_pct"`
	FirstFlight    string            `json:"first_flight"`
	Country        string            `json:"country"`
	Company        string            `json:"company"`
	Height         Height            `json:"height"`
	Diameter       Diameter          `json:"diameter"`
	Mass           Mass              `json:"mass"`
	PayloadWeights []PayloadWeight   `json:"payload_weights"`
	FirstStage     RocketFirstStage  `json:"first_stage"`
	SecondStage    RocketSecondStage `json:"second_stage"`
	Engines        Engines           `json:"engines"`
	LandingLegs    LandingLegs       `json:"landing_legs"`
	FlickrImages   []string          `json:"flickr_images"`
	Wikipedia      string            `json:"wikipedia"`
	Description    string            `json:"description"`
	RocketId       string            `json:"rocket_id"`
	RocketName     string            `json:"rocket_name"`
	RocketType     string            `json:"rocket_type"`
}

func (s *RocketsService) Get(rocketID string) (*Rocket, error) {
	if rocketID == "" {
		return nil, ErrInvalidID
	}

	u := fmt.Sprintf("rockets/%v", rocketID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	c := new(Rocket)
	err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *RocketsService) ListAll(opt *ListOptions) ([]*Rocket, error) {
	u := "rockets"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var c []*Rocket
	err = s.client.Do(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
