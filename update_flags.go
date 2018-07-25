package main

import (
	"os"
	"strings"
)

// https://emojipedia.org/flags/
// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
func updateFlags() {
	// Check mtzdateTimezones for countries and country codes
	updateFlagMap(mtzdateTimezones)

	// Check MTZDATE_FLAGS for countries and country codes
	if _, ok := os.LookupEnv("MTZDATE_FLAGS"); ok {
		updateFlagMap(
			strings.Split(os.Getenv("MTZDATE_FLAGS"), ","),
		)
	}
}
