package spacex

import (
	"fmt"
)

type HistoricalEventsService service

type Links struct {
	Reddit    *string `json:"reddit"`
	Article   *string `json:"article"`
	Wikipedia *string `json:"wikipedia"`
}

type HistoricalEvent struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	EventDateUtc  string `json:"event_date_utc"`
	EventDateUnix int    `json:"event_date_unix"`
	FlightNumber  *int   `json:"flight_number"`
	Details       string `json:"details"`
	Links         Links  `json:"links"`
}

type HistoricalEventListOptions struct {
	ID           int    `url:"id,omitempty"` // Filter by historical event id
	Start        string `url:"start,omitempty"`
	End          string `url:"end,omitempty"`
	FlightNumber int    `url:"flight_number,omitempty"`
}

func (s *HistoricalEventsService) Get(ID int) (*HistoricalEvent, error) {
	if ID <= 0 {
		return nil, ErrInvalidID
	}

	u := fmt.Sprintf("history/%v", ID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	c := new(HistoricalEvent)
	err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *HistoricalEventsService) ListAll(baseOpt *ListOptions, extOpt *HistoricalEventListOptions) ([]*HistoricalEvent, error) {
	u := "history"
	u, err := addOptions(u, baseOpt, extOpt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var c []*HistoricalEvent
	err = s.client.Do(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
