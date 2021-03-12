package models

// Location struct for country location data.
type Location struct {
	Country   string     `json:"country"`
	Provinces []Province `json:"provinces"`
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	Date      string     `json:"date"`
}
