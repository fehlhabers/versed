package main

import "fmt"

const (
	versed string = "versed.yml"
)

func main() {
	config  := GetConfig()
	fmt.Println(config.Target)
	for _,val := range config.Configs {
		fmt.Println(val.Source)
		fmt.Println(val.Version)
	}
}
