package hr_system

type HrSystem interface {
	Punch(status PunchStatus) error
}
