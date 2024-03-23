package config

import "time"

type Configuration struct {
	AppSettings AppSettings `json:"app_settings"`
}

type AppSettings struct {
	IsOneTime       bool          `json:"is_one_time"`
	StandByDuration time.Duration `json:"standby_duration"`
	WatchDirectory  string        `json:"watch_directory"`
}
