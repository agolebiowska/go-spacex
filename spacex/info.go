package spacex

type InfoService service

type Headquarters struct {
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
}

type InfoLinks struct {
	Website     *string `json:"website"`
	Flickr      *string `json:"flickr"`
	Twitter     *string `json:"twitter"`
	ElonTwitter *string `json:"elon_twitter"`
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
	Valuation     int64        `json:"valuation"`
	Headquarters  Headquarters `json:"headquarters"`
	Links         InfoLinks    `json:"links"`
	Summary       string       `json:"summary"`
}

func (s *InfoService) ListAll() (*Info, error) {
	u := "info"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	c := new(Info)
	err = s.client.Do(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
