package spacex

import (
	"fmt"
)

type DragonsService service

type HeatShield struct {
	Material    string  `json:"material"`
	SizeMeters  float32 `json:"size_meters"`
	TempDegrees float32 `json:"temp_degrees"`
	DevPartner  string  `json:"dev_partner"`
}

type Thruster struct {
	Type   string `json:"type"`
	Amount int    `json:"amount"`
	Pods   int    `json:"pods"`
	Fuel1  string `json:"fuel_1"`
	Fuel2  string `json:"fuel_2"`
	Thrust Thrust `json:"thrust"`
}

type Thrust struct {
	KN  float32 `json:"kN"`
	Lbf int     `json:"lbf"`
}

type Mass struct {
	Kg int `json:"kg"`
	Lb int `json:"lb"`
}

type Volume struct {
	CubicMeters int `json:"cubic_meters"`
	CubicFeet   int `json:"cubic_feet"`
}

type PressurizedCapsule struct {
	PayloadVolume Volume `json:"payload_volume"`
}

type Cargo struct {
	SolarArray         int  `json:"solar_array"`
	UnpressurizedCargo bool `json:"unpressurized_cargo"`
}

type Trunk struct {
	TrunkVolume Volume `json:"trunk_volume"`
	Cargo       Cargo  `json:"cargo"`
}

type Dimensions struct {
	Meters float32 `json:"meters"`
	Feet   float32 `json:"feet"`
}

type Dragon struct {
	ID                 string             `json:"id"`
	Name               string             `json:"name"`
	Type               string             `json:"type"`
	Active             bool               `json:"active"`
	CrewCapacity       int                `json:"crew_capacity"`
	SidewallAngleDeg   int                `json:"sidewall_angle_deg"`
	OrbitDurationYr    int                `json:"orbit_duration_yr"`
	DryMassKg          int                `json:"dry_mass_kg"`
	DryMassLb          int                `json:"dry_mass_lb"`
	FirstFlight        string             `json:"first_flight"`
	HeatShield         HeatShield         `json:"heat_shield"`
	WaterLanding       bool               `json:"water_landing"`
	Details            string             `json:"details"`
	Thrusters          []Thruster         `json:"thrusters"`
	LaunchPayloadMass  Mass               `json:"launch_payload_mass"`
	LaunchPayloadVol   Volume             `json:"launch_payload_vol"`
	ReturnPayloadMass  Mass               `json:"return_payload_mass"`
	ReturnPayloadVol   Volume             `json:"return_payload_vol"`
	PressurizedCapsule PressurizedCapsule `json:"pressurized_capsule"`
	Trunk              Trunk              `json:"trunk"`
	HeightWTrunk       Dimensions         `json:"height_w_trunk"`
	Diameter           Dimensions         `json:"diameter"`
	Wikipedia          string             `json:"wikipedia"`
	Description        string             `json:"description"`
}

type DragonListOptions struct {
	ID     bool `url:"id,omitempty"`     // Set as true to show mongo document id's
	Limit  int  `url:"limit,omitempty"`  // Limit results returned, defaults to all documents returned
	Offset int  `url:"offset,omitempty"` // Offset or skip results from the beginning of the query
}

func (s *DragonsService) Get(serial string) (*Dragon, error) {
	if serial == "" {
		return nil, ErrInvalidID
	}

	u := fmt.Sprintf("dragons/%v", serial)
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

func (s *DragonsService) ListAll(opt *DragonListOptions) ([]*Dragon, error) {
	u := "dragons"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var c []*Dragon
	err = s.client.Do(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
