package main

import (
	"fmt"
	"os"

	"github.com/Songmu/enslaver"
)

func main() {
	l := len(os.Args)
	opts := []string{}
	args := []string{}
	for i, v := range os.Args {
		if v == "--" && i+1 < l {
			args = os.Args[i+1:]
			break
		}
		opts = append(opts, v)
	}
	if len(args) == 0 {
		opts = []string{}
		args = os.Args[1:]
	}
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "No commands are specified")
		os.Exit(1)
	}

	enslaver.Command(args[0], args[1:]...).Run()
}
