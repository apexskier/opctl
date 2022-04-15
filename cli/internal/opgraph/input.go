package opgraph

import (
	"errors"
)

type InputEvent int

const (
	UpInputEvent    InputEvent = 0
	DownInputEvent  InputEvent = 1
	LeftInputEvent  InputEvent = 2
	RightInputEvent InputEvent = 3
	ClearInputEvent InputEvent = 4
)

// Input manages keyboard input
type Input struct {
	processing []byte
	Events     chan InputEvent
}

var (
	// ErrCtlC is returned when the user typed ctl-c, the program should handle SIGINT
	ErrCtlC = errors.New("ctl-c")
	// ErrCtlD is returned when the user typed ctl-d, the program should handle EOF
	ErrCtlD = errors.New("ctl-d")
)

// Consume takes in a byte of date from keyboard entry, and returns true if it's
// the last one expected to be consumed
func (i *Input) Consume(b byte) error {
	if len(i.processing) == 2 {
		// we're in a control sequence
		switch b {
		case 65: // up
			i.Events <- UpInputEvent
		case 66: // down
			i.Events <- DownInputEvent
		case 67: // right
			i.Events <- RightInputEvent
		case 68: // left
			i.Events <- LeftInputEvent
		}
		return nil
	}
	switch b {
	case 3: // ctl+c
		return ErrCtlC
	case 4: // ctl+d
		i.Events <- ClearInputEvent
	case 127: // backspace
	case 27: // esc
		i.processing = append(i.processing, b)
	case 91: // [
		if len(i.processing) == 1 {
			i.processing = append(i.processing, b)
			return nil
		}
	}
	return nil
}
