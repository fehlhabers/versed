package main

import "flag"

func main() {
	var versed string
	flag.StringVar(&versed, "f", "versed.yml", "Specify config location")
	flag.Parse()
	GetConfig(versed).Convert()
}
