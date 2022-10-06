package clioutput

import (
	"fmt"
	"io"
	"strings"

	"github.com/opctl/opctl/cli/internal/clicolorer"
	"github.com/opctl/opctl/sdks/go/model"
)

// CliOutput allows mocking/faking output
type CliOutput interface {
	// silently disables coloring
	DisableColor()

	// outputs a msg requiring attention
	Attention(s string)

	// outputs a warning message (looks like an error but on stdout)
	Warning(s string)

	// outputs an error msg
	Error(s string)

	// outputs an event
	// @TODO: not generic
	Event(event *model.Event)

	// outputs a success msg
	Success(s string)
}

func New(
	cliColorer clicolorer.CliColorer,
	opFormatter OpFormatter,
	errWriter io.Writer,
	stdWriter io.Writer,
) (CliOutput, error) {
	return _cliOutput{
		cliColorer:  cliColorer,
		opFormatter: opFormatter,
		errWriter:   errWriter,
		stdWriter:   stdWriter,
	}, nil
}

type _cliOutput struct {
	cliColorer  clicolorer.CliColorer
	opFormatter OpFormatter
	errWriter   io.Writer
	stdWriter   io.Writer
}

func (clio _cliOutput) DisableColor() {
	clio.cliColorer.DisableColor()
}

func (clio _cliOutput) Attention(s string) {
	io.WriteString(
		clio.stdWriter,
		fmt.Sprintln(
			clio.cliColorer.Attention(s),
		),
	)
}

func (clio _cliOutput) Warning(s string) {
	io.WriteString(
		clio.stdWriter,
		fmt.Sprintln(
			clio.cliColorer.Error(s),
		),
	)
}

func (clio _cliOutput) Error(s string) {
	io.WriteString(
		clio.errWriter,
		fmt.Sprintln(
			clio.cliColorer.Error(s),
		),
	)
}

func (clio _cliOutput) Event(event *model.Event) {
	switch {
	case event.CallEnded != nil &&
		event.CallEnded.Call.Container != nil:
		clio.containerExited(event)

	case event.CallStarted != nil &&
		event.CallStarted.Call.Container != nil:
		clio.containerStarted(event)

	case event.ContainerStdErrWrittenTo != nil:
		clio.containerStdErrWrittenTo(event.ContainerStdErrWrittenTo)

	case event.ContainerStdOutWrittenTo != nil:
		clio.containerStdOutWrittenTo(event.ContainerStdOutWrittenTo)

	case event.CallEnded != nil &&
		event.CallEnded.Call.Op != nil:
		clio.opEnded(event)

	case event.CallStarted != nil && event.CallStarted.Call.Op != nil:
		clio.opStarted(event.CallStarted)
	}
}

func (clio _cliOutput) containerExited(event *model.Event) {
	var color func(s string) string
	var writer io.Writer
	var message string
	switch event.CallEnded.Outcome {
	case model.OpOutcomeSucceeded:
		message = "exited"
		color = clio.cliColorer.Success
		writer = clio.stdWriter
	case model.OpOutcomeKilled:
		message = "killed"
		color = clio.cliColorer.Info
		writer = clio.stdWriter
	default:
		message = "crashed"
		color = clio.cliColorer.Error
		writer = clio.errWriter
	}

	if event.CallEnded.Call.Container.Image.Ref != nil {
		message = fmt.Sprintf("%s ", *event.CallEnded.Call.Container.Image.Ref) + message
	} else {
		message += "unknown container " + message
	}
	message = color(message)
	if event.CallEnded.Error != nil {
		message += color(":") + " " + event.CallEnded.Error.Message
	}

	io.WriteString(
		writer,
		fmt.Sprintf(
			"%s%s\n",
			clio.outputPrefix(event.CallEnded),
			message,
		),
	)
}

func (clio _cliOutput) containerStarted(event *model.Event) {
	message := "started "
	if event.CallStarted.Call.Container.Image.Ref != nil {
		message += *event.CallStarted.Call.Container.Image.Ref
	} else {
		message += "unknown container"
	}

	io.WriteString(
		clio.stdWriter,
		fmt.Sprintf(
			"%s%s\n",
			clio.outputPrefix(event.CallStarted),
			clio.cliColorer.Info(message),
		),
	)
}

func (clio _cliOutput) outputPrefix(event model.OpEvent) string {
	parts := []string{
		fmt.Sprintf("%.8s", fmt.Sprintf("%-8s", event.Id())),
	}
	opRef := clio.opFormatter.FormatOpRef(event.Ref())
	if opRef != "" {
		parts = append(parts, opRef)
	}
	return clio.cliColorer.Muted("["+strings.Join(parts, " ")+"]") + " "
}

func (clio _cliOutput) containerStdErrWrittenTo(event *model.ContainerStdErrWrittenTo) {
	io.WriteString(
		clio.errWriter,
		fmt.Sprintf(
			"%s%s",
			clio.outputPrefix(event),
			event.Data,
		),
	)
}

func (clio _cliOutput) containerStdOutWrittenTo(event *model.ContainerStdOutWrittenTo) {
	io.WriteString(
		clio.stdWriter,
		fmt.Sprintf(
			"%s%s",
			clio.outputPrefix(event),
			event.Data,
		),
	)
}

func (clio _cliOutput) opEnded(event *model.Event) {
	var color func(s string) string
	var writer io.Writer
	var message string
	switch event.CallEnded.Outcome {
	case model.OpOutcomeSucceeded:
		message = "succeeded"
		color = clio.cliColorer.Success
		writer = clio.stdWriter
	case model.OpOutcomeKilled:
		message = "killed"
		color = clio.cliColorer.Info
		writer = clio.stdWriter
	default:
		message = "failed"
		color = clio.cliColorer.Error
		writer = clio.errWriter
	}

	message = color(fmt.Sprintf("op %s", message))
	if event.CallEnded.Error != nil {
		message += color(":") + " " + event.CallEnded.Error.Message
	}

	io.WriteString(
		writer,
		fmt.Sprintf(
			"%s%s\n",
			clio.outputPrefix(event.CallEnded),
			message,
		),
	)
}

func (clio _cliOutput) opStarted(event *model.CallStarted) {
	io.WriteString(
		clio.stdWriter,
		fmt.Sprintf(
			"%s%s\n",
			clio.outputPrefix(event),
			clio.cliColorer.Info("started op"),
		),
	)
}

func (clio _cliOutput) Success(s string) {
	io.WriteString(
		clio.stdWriter,
		fmt.Sprintln(
			clio.cliColorer.Success(s),
		),
	)
}
