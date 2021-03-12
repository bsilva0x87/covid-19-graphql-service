package models

// Country struct for country location data.
type Country struct {
	Country    string  `json:"country"`
	Code       string  `json:"code"`
	Confirmed  int     `json:"confirmed"`
	Recovered  int     `json:"recovered"`
	Critical   int     `json:"critical"`
	Deaths     int     `json:"deaths"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	LastChange string  `json:"lastChange"`
	LastUpdate string  `json:"lastUpdate"`
}
