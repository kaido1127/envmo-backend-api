package model

type EnvInfoModel struct {
	MacAddr           string   `json:"device_id"`
	TemperatureInC    *float32 `json:"temperature_in_celcius,omitempty"`
	HumidityInPercent *float32 `json:"humidity_in_percent,omitempty"`
	PressureInHpa     *float32 `json:"pressure_in_hpa,omitempty"`
	GasInPpm          *float32 `json:"gas_in_ppm,omitempty"`
	CreatedAt         *int64   `json:"created_at,omitempty"`
	Location          *string  `json:"location_from_ip,omitempty"`
}
