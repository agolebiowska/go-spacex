package spacex

type InfoService service

type Headquarters struct {
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
}

type Info struct {
	Name          string       `json:"name"`
	Founder       string       `json:"founder"`
	Founded       int          `json:"founded"`
	Employees     int          `json:"employees"`
	Vehicles      int          `json:"vehicles"`
	LaunchSites   int          `json:"launch_sites"`
	TestSites     int          `json:"test_sites"`
	CEO           string       `json:"ceo"`
	CTO           string       `json:"cto"`
	COO           string       `json:"coo"`
	CtoPropulsion string       `json:"cto_propulsion"`
	Valuation     string       `json:"valuation"`
	Headquarters  Headquarters `json:"headquarters"`
	Summary       string       `json:"summary"`
}

func (s *InfoService) ListAll() ([]*Info, error) {
	u := "info"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var c []*Info
	err = s.client.Do(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
