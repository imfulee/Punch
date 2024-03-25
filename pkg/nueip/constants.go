package nueip

import "time"

type PunchStatus string

const (
	PunchIn        PunchStatus = "in"
	PunchOut       PunchStatus = "out"
	DefaultTimeout             = 10 * time.Second
)

func (p PunchStatus) IsValid() bool {
	switch p {
	case PunchIn, PunchOut:
		return true
	default:
		return false
	}
}
