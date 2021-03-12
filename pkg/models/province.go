package models

// Province struct for province location data.
type Province struct {
	Province  string `json:"province"`
	Confirmed int    `json:"confirmed"`
	Recovered int    `json:"recovered"`
	Deaths    int    `json:"deaths"`
	Active    int    `json:"active"`
}
