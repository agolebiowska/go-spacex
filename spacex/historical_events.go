package spacex

import (
	"fmt"
)

type HistoricalEventsService service

type Links struct {
	Reddit 		string `json:"reddit"`
	Article		string `json:"article"`
	Wikipedia 	string `json:"wikipedia"`
}

type HistoricalEvent struct {
	ID 				int 	`json:"id"`
	Title 			string 	`json:"title"`
	EventDateUtc 	string 	`json:"event_date_utc"`
	EventDateUnix 	int 	`json:"event_date_unix"`
	FlightNumber 	int 	`json:"flight_number"`
	Details 		string 	`json:"details"`
	Links 			Links 	`json:"links"`
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

	c:= new(HistoricalEvent)
	err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *HistoricalEventsService) ListAll() ([]*HistoricalEvent, error) {
	u := "history"
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