package main

import (
	"fmt"
	"os"
	"time"

	"github.com/keypad/clock/src/clock"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "serve" {
		if err := clock.Serve(":4173"); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
	zone := "local"
	if len(os.Args) > 1 {
		zone = os.Args[1]
	}
	text, err := clock.Render(zone, time.Now())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, clock.Help())
		os.Exit(1)
	}
	fmt.Println(text)
}
