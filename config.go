package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Source struct {
	Source  string `yaml:"source"`
	Version string `yaml:"version"`
}

type Config struct {
	Target  string   `yaml:"target"`
	Configs []Source `yaml:"sources"`
}

func GetConfig() Config {

	yamlFile, err := os.ReadFile(versed)
	if err != nil {
		log.Fatalf("failed to read configuration in file %s due to error: %s", versed, err)
	}

	var config Config
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatalf("failed to parse yaml file due to error: %s", err)
	}

	return config
}
