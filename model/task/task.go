package task

import (
	"go_clock/model/common"
)

type TaskStatus uint8

const (
	Planning TaskStatus = iota
	Finished
)

type Task struct {
	ID           string
	TaskName     string
	TaskContent  string
	AlertContent string
	TimeStamp    int64
	Status       TaskStatus
}

func (task Task) Type() common.ClockDataType {
	return common.TaskType
}

func UpdateStatus(t *Task, status TaskStatus) {
	t.Status = status
}
