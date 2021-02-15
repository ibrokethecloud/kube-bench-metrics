package wrapper

import (
	"encoding/json"
	"os/exec"

	"github.com/aquasecurity/kube-bench/check"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Wrapper is the place holder for storing the parse results //
type Wrapper struct {
	Results       check.OverallControls
	CommandOutput []byte
	Error         error
	NodeType      []string
	NodeName      string
}

// NewWrapper is the initialisation function
// It returns the wrapper object which can then be used to generate the metrics
func NewWrapper(nodeType string, hostname string) (w *Wrapper) {
	nodeChecks := []string{}
	if nodeType == "master" {
		nodeChecks = []string{"master", "node"}
	} else {
		nodeChecks = []string{"node"}
	}
	w = &Wrapper{
		NodeType: nodeChecks,
		NodeName: hostname,
	}
	return w
}

// ParseResults will populate the Results in the
// Wrapper object which can then be exposed as prometheus metrics
func (w *Wrapper) ParseResults() error {

	err := json.Unmarshal(w.CommandOutput, &w.Results)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

// RunBenchMarking will run kube-bench using os.Exec
// and send the byte buffer to ParseResults.
func (w *Wrapper) RunBenchMarking(ctx *cli.Context, node string) (err error) {
	args := buildCommand(ctx, node)
	err = w.runCommand(args)

	if err != nil {
		logrus.Error(err)
	}

	return err
}

// runCommand is a wrapper to run the command and return an
// output []byte array.
func (w *Wrapper) runCommand(args []string) (err error) {
	logrus.Debug("Args: ", args)
	cmd := exec.Command("kube-bench", args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Error(string(output))
		return err
	}

	w.CommandOutput = output

	err = w.ParseResults()

	if err != nil {
		logrus.Error(err)
	}

	return err
}

// buildCommand prepares the command for execution
// it will parse over cli context and also add
// a few defaults to the list of arguments
func buildCommand(ctx *cli.Context, nodeType string) (args []string) {
	checks := ctx.String("check")
	version := ctx.String("versionOverride")

	logrus.Debug(checks, " ", version)
	if checks != "" {
		args = append(args, "--check", checks)
	}

	if version != "" {
		args = append(args, "--version", version)
	}

	// Append the defaults //
	// Want Json output only, want noremediation and nosummary
	args = append(args, nodeType, "--json", "--noremediations", "--nosummary")

	return args
}
