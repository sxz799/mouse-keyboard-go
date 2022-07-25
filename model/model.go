package model

import "time"

type Operation struct {
	Type      string
	X         int
	Y         int
	MouseType string
	Key       string
	InputStr  string
	WaitTime  time.Duration
}
