package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMainWithFlags(t *testing.T) {
	cleanOutput()

	config := GetConfig("test/versed.yml")

	os.Args = []string{"test", "-f", "test/versed.yml"}
	main()

	output, err := ioutil.ReadFile("test/output/example.json")
	if err != nil {
		t.Errorf("Did not create any output")
	}

	assertOutputToConfig(output, t, config)
	cleanOutput()
}
