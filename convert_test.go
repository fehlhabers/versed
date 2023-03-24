package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	cleanOutput()
	config := NewTestConfig()

	config.Convert()
	output, err := ioutil.ReadFile("test/output/example.json")
	if err != nil {
		t.Errorf("Did not create any output")
	}

	var payload []string
	err = json.Unmarshal(output, &payload)
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

	cleanOutput()
}

func cleanOutput() {
	os.Remove("test/output/example.json")
}
