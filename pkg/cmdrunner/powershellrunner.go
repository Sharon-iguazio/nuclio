/*
Copyright 2017 The Nuclio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmdrunner

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	"github.com/nuclio/nuclio/pkg/common"

	"github.com/nuclio/errors"
	"github.com/nuclio/logger"
)

type PowerShellRunner struct {
	logger logger.Logger
	shell  string
}

func NewPowerShellRunner(parentLogger logger.Logger) (*ShellRunner, error) {
	ps, err := exec.LookPath("powershell.exe")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to lookup for powershell path")
	}
	return &ShellRunner{
		logger: parentLogger.GetChild("powershell_runner"),
		shell:  ps,
	}, nil
}

func (psr *PowerShellRunner) Run(runOptions *RunOptions, format string, vars ...interface{}) (RunResult, error) {

	// support missing runOptions for tests that send nil
	if runOptions == nil {
		runOptions = &RunOptions{}
	}

	// format the command
	formattedCommand := fmt.Sprintf(format, vars...)
	redactedCommand := common.Redact(runOptions.LogRedactions, formattedCommand)
	psr.logger.DebugWith("Executing", "command", redactedCommand)

	// create a command
	cmd := exec.Command(psr.shell, formattedCommand)

	// if there are runOptions, set them
	if runOptions.WorkingDir != nil {
		cmd.Dir = *runOptions.WorkingDir
	}

	// get environment variables if any
	if runOptions.Env != nil {
		cmd.Env = getEnvFromOptions(runOptions)
	}

	if runOptions.Stdin != nil {
		cmd.Stdin = strings.NewReader(*runOptions.Stdin)
	}

	runResult := RunResult{
		ExitCode: 0,
	}

	if err := runAndCaptureOutput(cmd, runOptions, &runResult); err != nil {
		var exitCode int

		// Did the command fail because of an unsuccessful exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.Sys().(syscall.WaitStatus).ExitStatus()
		}

		runResult.ExitCode = exitCode

		psr.logger.DebugWith("Failed to execute command",
			"output", runResult.Output,
			"stderr", runResult.Stderr,
			"exitCode", runResult.ExitCode,
			"err", err)

		return runResult, errors.Wrapf(err, "stdout:\n%s\nstderr:\n%s", runResult.Output, runResult.Stderr)
	}

	psr.logger.DebugWith("Command executed successfully",
		"output", runResult.Output,
		"stderr", runResult.Stderr,
		"exitCode", runResult.ExitCode)

	return runResult, nil
}

func (psr *PowerShellRunner) SetShell(shell string) {
	psr.shell = shell
}
