package engine

import (
	"math"
	"time"
)

// Telemetry represents the live payload
type Telemetry struct {
	DayOfWeek       time.Weekday `json:"day_of_week"`
	HourOfDay       int          `json:"hour_of_day"`
	IsPeakSeason    bool         `json:"is_peak_season"`
	CurrentWeather  string       `json:"current_weather"`
	LiveUserTraffic int          `json:"live_user_traffic"`
}

// SystemResponse defines the structural output
type SystemResponse struct {
	RiskCoefficient float64 `json:"risk_coefficient"`
	TargetNodes     int     `json:"target_nodes"`
	HourlySpendRate float64 `json:"hourly_spend_rate_usd"`
	ActionTaken     string  `json:"action_taken"`
}

// CalculatePredictiveRisk computes the real-time node requirement
func CalculatePredictiveRisk(t Telemetry) float64 {
	coefficient := 1.0

	// Seasonal & Calendar Multipliers
	if t.IsPeakSeason {
		coefficient *= 2.5
	}
	if t.DayOfWeek == time.Saturday || t.DayOfWeek == time.Sunday {
		coefficient *= 1.8
	}
	if t.HourOfDay >= 8 && t.HourOfDay <= 11 {
		coefficient *= 1.4
	} else if t.HourOfDay >= 17 && t.HourOfDay <= 21 {
		coefficient *= 1.3
	}

	// Live Weather Volatility
	switch t.CurrentWeather {
	case "Blizzard":
		coefficient *= 3.5
	case "Sunny":
		coefficient *= 1.6
	case "Rain":
		coefficient *= 0.4
	default:
		coefficient *= 1.0
	}

	// Traffic Volatility (Logarithmic floor to prevent system overflow)
	if t.LiveUserTraffic > 0 {
		coefficient += math.Log10(float64(t.LiveUserTraffic))
	}

	return coefficient
}