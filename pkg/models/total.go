package models

// Total struct contains the report totals data.
type Total struct {
	Confirmed  int    `json:"confirmed"`
	Recovered  int    `json:"recovered"`
	Critical   int    `json:"critical"`
	Deaths     int    `json:"deaths"`
	LastChange string `json:"lastChange"`
	LastUpdate string `json:"lastUpdate"`
}
