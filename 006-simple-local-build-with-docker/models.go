package main

type appInfo struct {
	Running  bool     `json:"running"`
	TimeInfo timeInfo `json:"time"`
}

type timeInfo struct {
	StartedAt     string  `json:"startedAt"`
	Now           string  `json:"now"`
	UptimeInSec   int64   `json:"uptimeInSec"`
	UptimeInHours float32 `json:"uptimeInHours"`
}
