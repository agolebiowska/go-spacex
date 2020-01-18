package spacex

import (
	"fmt"
	"time"
)

type LaunchesService service

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

type FirstStageCore struct {
	CoreSerial     string `json:"core_serial"`
	Flight         int    `json:"flight"`
	Block          int    `json:"block"`
	Gridfins       bool   `json:"gridfins"`
	Legs           bool   `json:"legs"`
	Reused         bool   `json:"reused"`
	LandSuccess    bool   `json:"land_success"`
	LandingIntent  bool   `json:"landing_intent"`
	LandingType    string `json:"landing_type"`
	LandingVehicle string `json:"landing_vehicle"`
}

type FirstStage struct {
	FirstStageCores []FirstStageCore `json:"cores"`
}

type SecondStage struct {
	Block    int       `json:"block"`
	Payloads []Payload `json:"payloads"`
}

type Fairings struct {
	Reused          bool `json:"reused"`
	RecoveryAttempt bool `json:"recovery_attempt"`
	Recovered       bool `json:"recovered"`
	Ship            int  `json:"ship"`
}

type Rocket struct {
	RocketID    string      `json:"rocket_id"`
	RocketName  string      `json:"rocket_name"`
	RocketType  string      `json:"rocket_type"`
	FirstStage  FirstStage  `json:"first_stage"`
	SecondStage SecondStage `json:"second_stage"`
	Fairings    Fairings    `json:"fairings"`
}

type Telemetry struct {
	FlightClub string `json:"flight_club"`
}

type LaunchSite struct {
	SiteID       string `json:"site_id"`
	SiteName     string `json:"site_name"`
	SiteNameLong string `json:"site_name_long"`
}

type LaunchLinks struct {
	MissionPatch      string   `json:"mission_patch"`
	MissionPatchSmall string   `json:"mission_patch_small"`
	RedditCampaign    string   `json:"reddit_campaign"`
	RedditLaunch      string   `json:"reddit_launch"`
	RedditRecovery    int      `json:"reddit_recovery"`
	RedditMedia       string   `json:"reddit_media"`
	Presskit          string   `json:"presskit"`
	ArticleLink       string   `json:"article_link"`
	Wikipedia         string   `json:"wikipedia"`
	VideoLink         string   `json:"video_link"`
	YoutubeId         string   `json:"youtube_id"`
	FlickrImages      []string `json:"flickr_images"`
}

type Timeline struct {
	WebcastLiftoff           int `json:"webcast_liftoff"`
	GoForPropLoading         int `json:"go_for_prop_loading"`
	Rp1Loading               int `json:"rp1_loading"`
	Stage1LoxLoading         int `json:"stage1_lox_loading"`
	Stage2LoxLoading         int `json:"stage2_lox_loading"`
	EngineChill              int `json:"engine_chill"`
	PrelaunchChecks          int `json:"prelaunch_checks"`
	PropellantPressurization int `json:"propellant_pressurization"`
	GoForLaunch              int `json:"go_for_launch"`
	Ignition                 int `json:"ignition"`
	Liftoff                  int `json:"liftoff"`
	Maxq                     int `json:"maxq"`
	Meco                     int `json:"meco"`
	StageSep                 int `json:"stage_sep"`
	SecondStageIgnition      int `json:"second_stage_ignition"`
	FairingDeploy            int `json:"fairing_deploy"`
	FirstStageEntryBurn      int `json:"first_stage_entry_burn"`
	Seco1                    int `json:"seco-1"`
	FirstStageLanding        int `json:"first_stage_landing"`
	SecondStageRestart       int `json:"second_stage_restart"`
	Seco2                    int `json:"seco-2"`
	PayloadDeploy            int `json:"payload_deploy"`
}

type Launch struct {
	FlightNumber          int         `json:"flight_number"`
	MissionName           string      `json:"mission_name"`
	MissionId             string      `json:"mission_id"`
	LaunchYear            string      `json:"launch_year"`
	LaunchDateUnix        int         `json:"launch_date_unix"`
	LaunchDateUtc         time.Time   `json:"launch_date_utc"`
	LaunchDateLocal       time.Time   `json:"launch_date_local"`
	IsTentative           bool        `json:"is_tentative"`
	TentativeMaxPrecision string      `json:"tentative_max_precision"`
	Tbd                   bool        `json:"tbd"`
	LaunchWindow          int         `json:"launch_window"`
	Rocket                Rocket      `json:"rocket"`
	Ships                 []string    `json:"ships"`
	Telemetry             Telemetry   `json:"telemetry"`
	LaunchSite            LaunchSite  `json:"launch_site"`
	LaunchSuccess         bool        `json:"launch_success"`
	LaunchLinks           LaunchLinks `json:"links"`
	Details               string      `json:"details"`
	Upcoming              bool        `json:"upcoming"`
	StaticFireDateUtc     string      `json:"static_fire_date_utc"`
	StaticFireDateUnix    int         `json:"static_fire_date_unix"`
	TimeLine              Timeline    `json:"timeline"`
}

type LaunchesListOptions struct {
	FlightID                string    `url:"flight_id,omitempty"`
	Start                   string    `url:"start,omitempty"`
	End                     string    `url:"end,omitempty"`
	FlightNumber            int       `url:"flight_number,omitempty"`
	LaunchYear              int       `url:"launch_year,omitempty"`
	LaunchDateUtc           time.Time `url:"launch_date_utc,omitempty"`
	LaunchDateLocal         time.Time `url:"launch_date_local,omitempty"`
	Tbd                     bool      `url:"tbd,omitempty"`
	RocketID                string    `url:"rocket_id,omitempty"`
	RocketName              string    `url:"rocket_name,omitempty"`
	RocketType              string    `url:"rocket_type,omitempty"`
	CoreSerial              string    `url:"core_serial,omitempty"`
	LandSuccess             bool      `url:"land_success,omitempty"`
	LandingIntent           bool      `url:"landing_intent,omitempty"`
	LandingType             string    `url:"landing_type,omitempty"`
	LandingVehicle          string    `url:"landing_vehicle,omitempty"`
	CapSerial               string    `url:"cap_serial,omitempty"`
	CoreFlight              int       `url:"core_flight,omitempty"`
	Block                   int       `url:"block,omitempty"`
	Gridfins                bool      `url:"gridfins,omitempty"`
	Legs                    bool      `url:"legs,omitempty"`
	CoreReuse               bool      `url:"core_reuse,omitempty"`
	CapsuleReuse            bool      `url:"capsule_reuse,omitempty"`
	FairingsReused          bool      `url:"fairings_reused,omitempty"`
	FairingsRecoveryAttempt bool      `url:"fairings_recovery_attempt,omitempty"`
	FairingsRecovered       bool      `url:"fairings_recovered,omitempty"`
	FairingsShip            string    `url:"fairings_ship,omitempty"`
	SiteID                  string    `url:"site_id,omitempty"`
	SiteName                string    `url:"site_name,omitempty"`
	PayloadId               string    `url:"payload_id,omitempty"`
	NoradId                 int       `url:"norad_id,omitempty"`
	Customer                string    `url:"customer,omitempty"`
	Nationality             string    `url:"nationality,omitempty"`
	Manufacturer            string    `url:"manufacturer,omitempty"`
	PayloadType             string    `url:"reference_system,omitempty"`
	Orbit                   string    `url:"orbit,omitempty"`
	ReferenceSystem         string    `url:"reference_system,omitempty"`
	Regime                  string    `url:"regime,omitempty"`
	Longitude               float32   `url:"longitude,omitempty"`
	SemiMajorAxisKm         float32   `url:"semi_major_axis_km,omitempty"`
	Eccentricity            float32   `url:"eccentricity,omitempty"`
	PeriapsisKm             float32   `url:"periapsis_km,omitempty"`
	ApoapsisKm              float32   `url:"apoapsis_km,omitempty"`
	InclinationDeg          float32   `url:"inclination_deg,omitempty"`
	PeriodMin               float32   `url:"period_min,omitempty"`
	LifespanYears           int       `url:"lifespan_years,omitempty"`
	Epoch                   string    `url:"epoch,omitempty"`
	MeanMotion              float32   `url:"mean_motion,omitempty"`
	Raan                    float32   `url:"raan,omitempty"`
	LaunchSuccess           bool      `url:"launch_success,omitempty"`
}

func (s *LaunchesService) Get(flightNumber string) (*Launch, error) {
	if flightNumber == "" {
		return nil, ErrInvalidID
	}

	u := fmt.Sprintf("launches/%v", flightNumber)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	c := new(Launch)
	err = s.client.Do(req, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *LaunchesService) ListAll(opt *LaunchesListOptions) ([]*Launch, error) {
	u := "launches"
	return s.list(u, opt)
}

func (s *LaunchesService) ListUpcoming(opt *LaunchesListOptions) ([]*Launch, error) {
	u := "launches/upcoming"
	return s.list(u, opt)
}

func (s *LaunchesService) ListPast(opt *LaunchesListOptions) ([]*Launch, error) {
	u := "launches/past"
	return s.list(u, opt)
}

func (s *LaunchesService) ListLatest(opt *LaunchesListOptions) ([]*Launch, error) {
	u := "launches/latest"
	return s.list(u, opt)
}

func (s *LaunchesService) ListNext(opt *LaunchesListOptions) ([]*Launch, error) {
	u := "launches/next"
	return s.list(u, opt)
}

func (s *LaunchesService) list(u string, opt *LaunchesListOptions) ([]*Launch, error) {
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var c []*Launch
	err = s.client.Do(req, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
