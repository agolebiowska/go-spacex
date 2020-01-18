package spacex

type Location struct {
	Name      string  `json:"name"`
	Region    string  `json:"region"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
