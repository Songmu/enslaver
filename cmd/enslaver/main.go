package main

import (
	"os"

	"github.com/Songmu/enslaver"
)

func main() {
	enslaver.Command(os.Args[1], os.Args[2:]...).Run()
}
