package cliprompt

import (
	"fmt"
	"os"

	"github.com/opctl/opctl/cli/internal/clioutput"
	"github.com/opctl/opctl/cli/internal/cliparamsatisfier/inputsrc"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/peterh/liner"
	"golang.org/x/term"
)

func New(
	cliOutput clioutput.CliOutput,
	inputs map[string]*model.ParamSpec,
) inputsrc.InputSrc {
	return cliPromptInputSrc{
		inputs:    inputs,
		cliOutput: cliOutput,
	}
}

// cliPromptInputSrc implements InputSrc interface by sourcing inputs from std in
type cliPromptInputSrc struct {
	inputs    map[string]*model.ParamSpec
	cliOutput clioutput.CliOutput
}

func (this cliPromptInputSrc) ReadString(
	inputName string,
) (*string, bool) {
	if param := this.inputs[inputName]; param != nil {
		var (
			isSecret    bool
			description string
			prompt      string
		)

		if !term.IsTerminal(int(os.Stdout.Fd())) {
			return nil, false
		}

		switch {
		case param.Array != nil:
			isSecret = param.Array.IsSecret
			// @TODO remove after deprecation period
			description = param.Array.Description
			prompt = "array"
		case param.Boolean != nil:
			// @TODO remove after deprecation period
			description = param.Boolean.Description
			prompt = "boolean"
		case param.Dir != nil:
			isSecret = param.Dir.IsSecret
			// @TODO remove after deprecation period
			description = param.Dir.Description
			prompt = "directory"
		case param.File != nil:
			isSecret = param.File.IsSecret
			// @TODO remove after deprecation period
			description = param.File.Description
			prompt = "file"
		case param.Number != nil:
			isSecret = param.Number.IsSecret
			// @TODO remove after deprecation period
			description = param.Number.Description
			prompt = "number"
		case param.Object != nil:
			isSecret = param.Object.IsSecret
			// @TODO remove after deprecation period
			description = param.Object.Description
			prompt = "object"
		case param.Socket != nil:
			isSecret = param.Socket.IsSecret
			// @TODO remove after deprecation period
			description = param.Socket.Description
			prompt = "socket"
		case param.String != nil:
			isSecret = param.String.IsSecret
			// @TODO remove after deprecation period
			description = param.String.Description
			prompt = "string"
		}
		prompt += ": "

		if param.Description != "" {
			// non-deprecated property takes precedence
			description = param.Description
		}

		line := liner.NewLiner()
		defer line.Close()
		line.SetCtrlCAborts(true)

		if description != "" {
			this.cliOutput.Attention(
				fmt.Sprintf("input: '%s'\n%s", inputName, description),
			)
		} else {
			this.cliOutput.Attention(
				fmt.Sprintf("input: '%s'", inputName),
			)
		}

		// liner has inconsistent behavior if non empty prompt arg passed so use ""
		var (
			err    error
			rawArg string
		)
		if isSecret {
			rawArg, err = line.PasswordPrompt(prompt)
		} else {
			rawArg, err = line.Prompt(prompt)
		}
		if err == nil {
			return &rawArg, true
		}
	}

	return nil, false
}
