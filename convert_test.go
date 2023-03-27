package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	cleanOutput()
	config := NewTestConfig("test/output")

	config.Convert()
	output, err := ioutil.ReadFile("test/output/example.json")
	if err != nil {
		t.Errorf("Did not create any output")
	}

	assertOutputToConfig(output, t, config)

	cleanOutput()
}

func TestOutputFolderCreated(t *testing.T) {
	cleanOutput()
	config := NewTestConfig("test/newoutput")

	config.Convert()
	output, err := ioutil.ReadFile("test/newoutput/example.json")
	if err != nil {
		t.Errorf("Did not create any output")
	}

	assertOutputToConfig(output, t, config)

	cleanOutput()
}

func TestOverwriteExisting(t *testing.T) {
	cleanOutput()

	existing := []byte("a lot of crap\ndata being written\nin a file which should be\nreplaced\by the new output")
	ioutil.WriteFile("test/output/example.json", existing, 0664)
	config := NewTestConfig("test/output")

	config.Convert()
	output, err := ioutil.ReadFile("test/output/example.json")
	if err != nil {
		t.Errorf("Did not create any output")
	}

	assertOutputToConfig(output, t, config)

	cleanOutput()
}

func assertOutputToConfig(output []byte, t *testing.T, config Config) {
	var payload []string
	err := json.Unmarshal(output, &payload)
	if err != nil {
		t.Errorf("Output was not properly updated. Should be a proper json. Was: %s", string(output))
	}

	for _, s := range config.Sources {
		found := 0
		for _, act := range payload {
			if act == s.Source+s.Version {
				found++
			}
		}

		if found != 1 {
			t.Errorf("Should be exactly one match in output")
		}
	}
}

func cleanOutput() {
	os.Remove("test/output/example.json")
	os.Remove("test/newoutput/example.json")
	os.Remove("test/newoutput")
}
