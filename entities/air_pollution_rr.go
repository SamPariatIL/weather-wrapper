package entities

type AirPollutionComponents struct {
	CO   float64 `json:"co"`
	NO   float64 `json:"no"`
	NO2  float64 `json:"no2"`
	O3   float64 `json:"o3"`
	SO2  float64 `json:"so2"`
	PM25 float64 `json:"pm2_5"`
	PM10 float64 `json:"pm10"`
	NH3  float64 `json:"nh3"`
}

type AirPollution struct {
	Coord
	List []struct {
		Dt   int `json:"dt"`
		Main struct {
			AQI int `json:"aqi"`
		}
		Components AirPollutionComponents `json:"components"`
	} `json:"list"`
}
