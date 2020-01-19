package spacex

import (
	"fmt"
)

type MissionsService service

type MinimisedMission struct {
	Name   string `json:"name"`
	Flight int    `json:"flight"`
}

type Mission struct {
	MissionName   string   `json:"mission_name"`
	MissionID     string   `json:"mission_id"`
	Manufacturers []string `json:"manufacturers"`
	PayloadIds    []string `json:"payload_ids"`
	Wikipedia     string   `json:"wikipedia"`
	Website       string   `json:"website"`
	Twitter       *string  `json:"twitter"`
	Description   string   `json:"description"`
}

type MissionsListOptions struct {
}

func (s *MissionsService) Get(missionID string) (*Mission, error) {
	if missionID == "" {
		return nil, ErrInvalidID
	}

	u := fmt.Sprintf("missions/%v", missionID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	c := new(Mission)
	err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *MissionsService) ListAll(baseOpt *ListOptions, extOpt *MissionsListOptions) ([]*Mission, error) {
	u := "missions"
	u, err := addOptions(u, baseOpt, extOpt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var c []*Mission
	err = s.client.Do(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
