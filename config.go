package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Source struct {
	Source  string
	Version string
}

type Config struct {
	Target  string
	Output  string
	Sources map[string]Source
}

func GetConfig(configFile string) Config {

	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("failed to read configuration in file %s due to error: %s", versed, err)
	}

	var config Config
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatalf("failed to parse yaml file due to error: %s", err)
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
	return s.Source + s.Version
}
