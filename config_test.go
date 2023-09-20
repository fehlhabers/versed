package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestValuesPopulated(t *testing.T) {
	expConf := NewTestConfig("test/output")
	actConf := GetConfig("test/versed.yml")

	for i, src := range actConf.Sources {
		if src != expConf.Sources[i] {
			t.Errorf("Sources did not match. Wanted %s - got %s", expConf.Sources[i], src)
		}
	}

	if expConf.Target != actConf.Target {
		t.Errorf("Target did not match. Wanted %s - got %s", expConf.Target, actConf.Target)
	}

	if expConf.Output != actConf.Output {
		t.Errorf("Output did not match. Wanted %s - got %s", expConf.Output, actConf.Output)
	}
}

func TestNonExistingFile(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		GetConfig("non-existing-file")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestNonExistingFile")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestInvalidFile(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		GetConfig("test/invalid_versed.yml")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestInvalidFile")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func NewTestConfig(output string) Config {
	return Config{
		Target: "test/data",
		Output: output,
		Sources: map[string]Source{
			"testsource1": "v1",
			"testsource2": "v2",
		},
	}
}
