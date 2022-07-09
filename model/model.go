package model

import "time"

type Operation struct {
	Type      string
	X         int
	Y         int
	MouseType string
	InputMsg  string
	WaitTime  time.Duration
}
