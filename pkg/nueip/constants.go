package nueip

type PunchStatus string

const (
	PunchIn  PunchStatus = "in"
	PunchOut PunchStatus = "out"
)

func (p PunchStatus) IsValid() bool {
	switch p {
	case PunchIn, PunchOut:
		return true
	default:
		return false
	}
}
