package main

const (
	versed string = "versed.yml"
)

func main() {
	GetConfig(versed).Convert()
}
