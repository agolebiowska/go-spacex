package spacex

import (
	"fmt"
	"time"
)

type LaunchesService service

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
	FlightNumber          int             `json:"flight_number"`
	MissionName           string          `json:"mission_name"`
	MissionId             string          `json:"mission_id"`
	LaunchYear            string          `json:"launch_year"`
	LaunchDateUnix        int             `json:"launch_date_unix"`
	LaunchDateUtc         time.Time       `json:"launch_date_utc"`
	LaunchDateLocal       time.Time       `json:"launch_date_local"`
	IsTentative           bool            `json:"is_tentative"`
	TentativeMaxPrecision string          `json:"tentative_max_precision"`
	Tbd                   bool            `json:"tbd"`
	LaunchWindow          int             `json:"launch_window"`
	Rocket                MinimisedRocket `json:"rocket"`
	Ships                 []string        `json:"ships"`
	Telemetry             Telemetry       `json:"telemetry"`
	LaunchSite            LaunchSite      `json:"launch_site"`
	LaunchSuccess         bool            `json:"launch_success"`
	LaunchLinks           LaunchLinks     `json:"links"`
	Details               string          `json:"details"`
	Upcoming              bool            `json:"upcoming"`
	StaticFireDateUtc     string          `json:"static_fire_date_utc"`
	StaticFireDateUnix    int             `json:"static_fire_date_unix"`
	TimeLine              Timeline        `json:"timeline"`
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
