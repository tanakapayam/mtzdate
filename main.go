package main

import (
	"os"
)

func main() {
	setTimezones()
	updateFlags()
	setWorkhours()

	if loop, ok := os.LookupEnv("MTZDATE_LOOP"); ok && loop != "" && loop != "0" {
		loopShowTimeTable()
	} else {
		showTimeTable()
	}
}
