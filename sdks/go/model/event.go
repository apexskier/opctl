package model

import "time"

// Event represents a distributed state change
type Event struct {
	CallEnded                *CallEnded                `json:"callEnded,omitempty"`
	CallStarted              *CallStarted              `json:"callStarted,omitempty"`
	ContainerStdErrWrittenTo *ContainerStdErrWrittenTo `json:"containerStdErrWrittenTo,omitempty"`
	ContainerStdOutWrittenTo *ContainerStdOutWrittenTo `json:"containerStdOutWrittenTo,omitempty"`
	Timestamp                time.Time                 `json:"timestamp"`
}

type OpOutcome string

const (
	OpOutcomeSucceeded OpOutcome = "SUCCEEDED"
	OpOutcomeFailed    OpOutcome = "FAILED"
	OpOutcomeKilled    OpOutcome = "KILLED"
)

func (o OpOutcome) String() string {
	return string(o)
}

type OpEvent interface {
	Id() string
	Ref() string
}

// CallEnded represents a call ended; no further events will occur for the call
type CallEnded struct {
	Call    Call            `json:"call"`
	OpRef   string          `json:"ref"`
	Error   *CallEndedError `json:"error,omitempty"`
	Outcome OpOutcome       `json:"outcome"`
}

func (e CallEnded) Id() string {
	return e.Call.ID
}

func (e CallEnded) Ref() string {
	if e.Call.Op != nil {
		return e.Call.Op.OpPath
	}
	return e.OpRef
}

// CallStarted represents the start of an op
type CallStarted struct {
	Call  Call   `json:"call"`
	OpRef string `json:"ref"`
}

func (e CallStarted) Id() string {
	return e.Call.ID
}

func (e CallStarted) Ref() string {
	if e.Call.Op != nil {
		return e.Call.Op.OpPath
	}
	return e.OpRef
}

// CallEndedError represents an error associated w/ an ended call
type CallEndedError struct {
	Message string `json:"message"`
}

// ContainerStdErrWrittenTo represents a single write to a containers std err.
type ContainerStdErrWrittenTo struct {
	Data        []byte `json:"data"`
	ContainerID string `json:"containerId"`
	OpRef       string `json:"opRef"`
}

func (e ContainerStdErrWrittenTo) Id() string {
	return e.ContainerID
}

func (e ContainerStdErrWrittenTo) Ref() string {
	return e.OpRef
}

// ContainerStdOutWrittenTo represents a single write to a containers std out.
type ContainerStdOutWrittenTo struct {
	Data        []byte `json:"data"`
	ContainerID string `json:"containerId"`
	OpRef       string `json:"opRef"`
}

func (e ContainerStdOutWrittenTo) Id() string {
	return e.ContainerID
}

func (e ContainerStdOutWrittenTo) Ref() string {
	return e.OpRef
}
