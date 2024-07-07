package model

type UserEnvInfoModel struct {
	DeviceID          string   `json:"device_id"`
	TemperatureInC    *float32 `json:"temperature_in_c,omitempty"`
	HumidityInPercent *float32 `json:"humidity_in_percent,omitempty"`
	PressureInHpa     *float32 `json:"pressure_in_hpa,omitempty"`
	GasInPpm          *float32 `json:"gas_in_ppm,omitempty"`
	CreatedAt         int64    `json:"created_at"`
	Location          *string  `json:"location,omitempty"`
}

func (t *UserEnvInfoModel) ToMap() map[string]interface{} {
	data := map[string]interface{}{
		"temperature_in_c":    t.TemperatureInC,
		"pressure_in_hpa":     t.PressureInHpa,
		"humidity_in_percent": t.HumidityInPercent,
		"gas_in_ppm":          t.GasInPpm,
		"created_at":          t.CreatedAt,
		"location":            t.Location,
	}

	return data
}
