package main

import (
	"fmt"
)

const (
	versed string = "versed.yml"
)

func main() {
	config := GetConfig(versed)
	fmt.Println(config.Target)
	for _, val := range config.Sources {
		fmt.Println(val.Source)
		fmt.Println(val.Version)
	}
}
