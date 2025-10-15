package common

type ClockDataType string

const (
	TaskType ClockDataType = "Task"
)

// Common a simple struct to description the value‘s basic info
type Common struct {
	Key  string
	Name string
	Type ClockDataType
}
