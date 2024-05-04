package model

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type EnvInfoModel struct {
	CreatedAt      int64     `json:"created_at"`
	Humidity       float64   `json:"humidity"`
	TemperatureInC float64   `json:"temperature_in_c"`
	Location       *Location `json:"location"`
}
