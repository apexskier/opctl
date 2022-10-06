package git

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

// Clone clones a git repository referenced by an opctl dataRef
func (gp *_git) Clone(
	ctx context.Context,
	dataRef string,
) error {
	parsedPkgRef, err := parseRef(dataRef)
	if err != nil {
		return fmt.Errorf("invalid git ref: %w", err)
	}

	// This method of git cloning is based on https://cs.opensource.google/go/go/+/refs/tags/go1.17.5:src/cmd/go/internal/get/get.go
	// Using a sub-command instead of go-git for clones makes credential helpers
	// and ssh keys to work without a bunch of extra work in opctl code.

	// Disable any prompting for passwords by Git itself.
	// Only has an effect for 2.3.0 or later, but avoiding
	// the prompt in earlier versions is just too hard.
	// If user has explicitly set GIT_TERMINAL_PROMPT=1, keep
	// prompting.
	// See golang.org/issue/9341 and golang.org/issue/12706.
	if os.Getenv("GIT_TERMINAL_PROMPT") == "" {
		if err := os.Setenv("GIT_TERMINAL_PROMPT", "0"); err != nil {
			return err
		}
	}

	// Also disable prompting for passwords by the 'ssh' subprocess spawned by
	// Git, because apparently GIT_TERMINAL_PROMPT isn't sufficient to do that.
	// Adding '-o BatchMode=yes' should do the trick.
	//
	// If a Git subprocess forks a child into the background to cache a new connection,
	// that child keeps stdout/stderr open. After the Git subprocess exits,
	// os /exec expects to be able to read from the stdout/stderr pipe
	// until EOF to get all the data that the Git subprocess wrote before exiting.
	// The EOF doesn't come until the child exits too, because the child
	// is holding the write end of the pipe.
	// This is unfortunate, but it has come up at least twice
	// (see golang.org/issue/13453 and golang.org/issue/16104)
	// and confuses users when it does.
	// If the user has explicitly set GIT_SSH or GIT_SSH_COMMAND,
	// assume they know what they are doing and don't step on it.
	// But default to turning off ControlMaster.
	if os.Getenv("GIT_SSH") == "" && os.Getenv("GIT_SSH_COMMAND") == "" {
		if err := os.Setenv("GIT_SSH_COMMAND", "ssh -o ControlMaster=no -o BatchMode=yes"); err != nil {
			return err
		}
	}

	// And one more source of Git prompts: the Git Credential Manager Core for Windows.
	//
	// See https://github.com/microsoft/Git-Credential-Manager-Core/blob/master/docs/environment.md#gcm_interactive.
	if os.Getenv("GCM_INTERACTIVE") == "" {
		if err := os.Setenv("GCM_INTERACTIVE", "never"); err != nil {
			return err
		}
	}

	destinationPath := parsedPkgRef.ToPath(gp.basePath)

	if err := os.RemoveAll(destinationPath); err != nil {
		return err
	}

	cmd := exec.CommandContext(
		ctx,
		"git",
		"clone",
		"--quiet",
		"--depth=1",
		"--single-branch",
		fmt.Sprintf("--branch=%s", parsedPkgRef.Version),
		fmt.Sprintf("https://%s", parsedPkgRef.Name),
		destinationPath,
	)
	if output, err := cmd.CombinedOutput(); err != nil {
		if string(output) != "" {
			return fmt.Errorf("%w: %s", err, output)
		}
		return err
	}

	return nil
}
