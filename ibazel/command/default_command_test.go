// Copyright 2017 The Bazel Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"os/exec"
	"syscall"
	"testing"

	mock_bazel "github.com/bazelbuild/bazel-watcher/bazel/testing"
)

func TestDefaultCommand(t *testing.T) {
	toKill := exec.Command("sleep", "5s")
	toKill.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	b := &mock_bazel.MockBazel{}
	c := &defaultCommand{
		args:   []string{"moo"},
		b:      b,
		cmd:    toKill,
		target: "//path/to:target",
	}

	if c.IsSubprocessRunning() {
		t.Errorf("New subprocess shouldn't have been started yet. State: %v", toKill.ProcessState)
	}

	toKill.Start()

	if !c.IsSubprocessRunning() {
		t.Errorf("New subprocess was never started. State: %v", toKill.ProcessState)
	}

	// This is synonymous with killing the job so use it to kill the job and test everything.
	c.NotifyOfChanges()
	assertKilled(t, toKill)
}
