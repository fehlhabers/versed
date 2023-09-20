package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Source string

type Config struct {
	Target  string
	Output  string
	Sources map[string]Source
}

func GetConfig(versed string) Config {

	yamlFile, err := os.ReadFile(versed)
	if err != nil {
		log.Fatalf("Failed to read configuration in <%s> due to error: %s", versed, err)
	}

	var config Config
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatalf("Failed to parse yaml file due to error: %s", err)
	}

	return config
}

func (c Config) oFilePath(file string) string {
	return c.Output + "/" + file
}

func (c Config) tFilePath(file string) string {
	return c.Target + "/" + file
}

func replToken(name string) string {
	return "&(versed." + name + ")"
}

func (s Source) replacement() string {
	return string(s)
}
