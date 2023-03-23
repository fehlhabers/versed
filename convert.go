package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func (c Config) Convert() {
	files, err := ioutil.ReadDir(c.Target)
	if err != nil {
		log.Fatalf("Unable to read files in target dir: %s - error: %s", c.Target, err)
	}

	for _, file := range files {
		if !file.IsDir() {
			c.convertToOutput(file.Name())
		}
	}

}

func (c Config) convertToOutput(file string) {
	copy, err := ioutil.ReadFile(c.tFilePath(file))
	if err != nil {
		log.Fatalf("Unable to read file: %s - error: %s", file, err)
	}

	rows := strings.Split(string(copy), "\n")

	for i, row := range rows {
		updated, nrow := c.replace(row)
		if updated {
			rows[i] = nrow
		}
	}
	ofile := []byte(strings.Join(rows, "\n"))

	err = ioutil.WriteFile(c.oFilePath(file), ofile, 0644)
	if err != nil {
		log.Fatalf("Unable to write file: %s - error: %s", file, err)
	}
}

func (c Config) replace(row string) (bool, string) {
	for sn, s := range c.Sources {
		if strings.Contains(row, replToken(sn)) {
			return true, strings.ReplaceAll(row, replToken(sn),s.replacement())
		}
	}
	return false, ""
}
