# Local development with native go tooling

First, ensure dependencies are installed

- [go](https://golang.org/doc/install) (use the version from the [go.mod file](./go.mod#L3))
- [dlv](https://github.com/go-delve/delve) (for debugging) - Because this project uses go modules, install this globally with `go get github.com/go-delve/delve` to run the tool outside of the project directory.

## CLI debugging

Build a binary with

`GOFLAGS=-tags=containers_image_openpgp go build -o opctl ./cli`

Run with [dlv](https://github.com/go-delve/delve), e.g.

`go run github.com/go-delve/delve/cmd/dlv --check-go-version=false --listen=127.0.0.1:40000 --headless=true --api-version=2 exec /PATH/TO/OPCTL/opctl-beta run dev`

This will suspend execution until a client connects on port 40000. If using VSCode, you can use the "Connect to dlv debugger" run configuration to connect.

On macOS, you can debug exit signal handling by sending an interrupt to the underlying PID. When the debugger connects, it will give you the with a message like: `Got a connection, launched process ../opctl/opctl-beta (pid = 79924).` (path to process and PID will be different). Kill with `kill -SIGINT 72958`.

## Node debugging

Ops are run in a "node" running as an external process, which means debugging the CLI directly doesn't catch most of the logic. To debug this external process, run `opctl node kill` to kill the current node (attached to port 42224), then use the VSCode "Run and debug an opctl node" run configuration to spin up a new node with the debugger attached or debug the CLI command "opctl node create".

## Tests

This project uses the [Ginkgo](https://github.com/onsi/ginkgo) test framework, which means some IDE's go tooling won't support normal methods of running tests. Run the VSCode "Run and debug this test suite" run configuration while viewing a test file to run the tests attached to the debugger.
