package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestValuesPopulated(t *testing.T) {
	expConf := newTestConfig()
	actConf := GetConfig("test/versed.yml")
	
	for i,src := range actConf.Sources {
		if src != expConf.Sources[i] {
			t.Fail()
		}
	}

	if expConf.Target != actConf.Target {
		t.Fail()
	}

	if expConf.Output != actConf.Output {
		t.Fail()
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

func newTestConfig() Config {
	return Config{
		Target: "data",
		Output: "output",
		Sources: []Source{
			{
				Name: "testsource1",
				Source: "source1",
				Version: "v1",
			},
			{
				Name: "testsource2",
				Source: "source2",
				Version: "v2",
			},
		},
	}
}
