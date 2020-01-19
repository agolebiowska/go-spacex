package spacex

type Height struct {
	Meters float32 `json:"meters"`
	Feet   float32 `json:"feet"`
}

type Diameter struct {
	Meters float32 `json:"meters"`
	Feet   float32 `json:"feet"`
}

type Mass struct {
	Kg int `json:"kg"`
	Lb int `json:"lb"`
}

type Thrust struct {
	KN  float32 `json:"kN"`
	Lbf int     `json:"lbf"`
}

type Position struct {
	Latitude  *float32 `json:"latitude"`
	Longitude *float32 `json:"longitude"`
}

type ListOptions struct {
	ID     bool   `url:"id,omitempty"`     // Set as true to show mongo document id's
	Limit  int    `url:"limit,omitempty"`  // Limit results returned, defaults to all documents returned
	Offset int    `url:"offset,omitempty"` // Offset or skip results from the beginning of the query
	Sort   string `url:"sort,omitempty"`   // Change result sorting by setting value to any parameter in this list
	Order  string `url:"order,omitempty"`  // Change result ordering by setting values of asc or desc
}

func NewLimitOption(l int) *ListOptions {
	return &ListOptions{false, l, 0, "", ""}
}

type LaunchesListOptions struct {
	FlightID                string  `url:"flight_id,omitempty"`
	Start                   string  `url:"start,omitempty"`
	End                     string  `url:"end,omitempty"`
	FlightNumber            int     `url:"flight_number,omitempty"`
	LaunchYear              int     `url:"launch_year,omitempty"`
	LaunchDateUtc           string  `url:"launch_date_utc,omitempty"`
	LaunchDateLocal         string  `url:"launch_date_local,omitempty"`
	Tbd                     bool    `url:"tbd,omitempty"`
	RocketID                string  `url:"rocket_id,omitempty"`
	RocketName              string  `url:"rocket_name,omitempty"`
	RocketType              string  `url:"rocket_type,omitempty"`
	CoreSerial              string  `url:"core_serial,omitempty"`
	LandSuccess             bool    `url:"land_success,omitempty"`
	LandingIntent           bool    `url:"landing_intent,omitempty"`
	LandingType             string  `url:"landing_type,omitempty"`
	LandingVehicle          string  `url:"landing_vehicle,omitempty"`
	CapSerial               string  `url:"cap_serial,omitempty"`
	CoreFlight              int     `url:"core_flight,omitempty"`
	Block                   int     `url:"block,omitempty"`
	Gridfins                bool    `url:"gridfins,omitempty"`
	Legs                    bool    `url:"legs,omitempty"`
	CoreReuse               bool    `url:"core_reuse,omitempty"`
	CapsuleReuse            bool    `url:"capsule_reuse,omitempty"`
	FairingsReused          bool    `url:"fairings_reused,omitempty"`
	FairingsRecoveryAttempt bool    `url:"fairings_recovery_attempt,omitempty"`
	FairingsRecovered       bool    `url:"fairings_recovered,omitempty"`
	FairingsShip            string  `url:"fairings_ship,omitempty"`
	SiteID                  string  `url:"site_id,omitempty"`
	SiteName                string  `url:"site_name,omitempty"`
	PayloadId               string  `url:"payload_id,omitempty"`
	NoradId                 int     `url:"norad_id,omitempty"`
	Customer                string  `url:"customer,omitempty"`
	Nationality             string  `url:"nationality,omitempty"`
	Manufacturer            string  `url:"manufacturer,omitempty"`
	PayloadType             string  `url:"reference_system,omitempty"`
	Orbit                   string  `url:"orbit,omitempty"`
	ReferenceSystem         string  `url:"reference_system,omitempty"`
	Regime                  string  `url:"regime,omitempty"`
	Longitude               float32 `url:"longitude,omitempty"`
	SemiMajorAxisKm         float32 `url:"semi_major_axis_km,omitempty"`
	Eccentricity            float32 `url:"eccentricity,omitempty"`
	PeriapsisKm             float32 `url:"periapsis_km,omitempty"`
	ApoapsisKm              float32 `url:"apoapsis_km,omitempty"`
	InclinationDeg          float32 `url:"inclination_deg,omitempty"`
	PeriodMin               float32 `url:"period_min,omitempty"`
	LifespanYears           int     `url:"lifespan_years,omitempty"`
	Epoch                   string  `url:"epoch,omitempty"`
	MeanMotion              float32 `url:"mean_motion,omitempty"`
	Raan                    float32 `url:"raan,omitempty"`
	LaunchSuccess           bool    `url:"launch_success,omitempty"`
}
