package hotomata

import (
	"bytes"
)

type TaskAction string

const (
	TaskActionAbort    = "abort"
	TaskActionContinue = "continue"
)

type TaskStatus string

const (
	TaskStatusSuccess = "success"
	TaskStatusWarning = "warning"
	TaskStatusError   = "error"
	TaskStatusSkip    = "skip"
)

type TaskResponse struct {
	Log    *bytes.Buffer
	Action TaskAction
	Status TaskStatus
}

type Runner interface {
	Run(string) *TaskResponse
}
