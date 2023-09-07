package hr_system

type PunchStatus string

const (
	PunchIn  PunchStatus = "in"
	PunchOut PunchStatus = "out"
)

func (p PunchStatus) IsValid() bool {
	switch p {
	case PunchStatus(PunchIn), PunchStatus(PunchOut):
		return true
	default:
		return false
	}
}
