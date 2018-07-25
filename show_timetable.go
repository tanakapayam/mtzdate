package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func showTimeTable() {
	maxLen := 0
	now := time.Now()

	for _, kv := range mtzdateTimezones {
		// tz[0]: desired city/time zone
		// tz[1]: actual time zone
		tz := strings.Split(kv, ":")

		label := func(z string) string {
			if z == utc {
				return ""
			}
			return z
		}(tz[0])

		if unicodeLen(label) > maxLen {
			maxLen = unicodeLen(label)
		}
	}
	maxLen++

	for _, kv := range mtzdateTimezones {
		// tz[0]: desired city/time zone
		// tz[1]: actual time zone
		tz := strings.Split(kv, ":")

		z, err := time.LoadLocation(tz[1])
		die(err)

		// Fri Jul 27 03:32:04 UTC 2018
		f := strings.Fields(
			now.In(z).Format(time.UnixDate),
		)

		// color workhours; pad timezone; drop year
		f = reformatTime(f)

		label := func(z string) string {
			if z == utc {
				return ""
			}
			return z
		}(tz[0])

		for _, c := range os.Getenv("MTZDATE_FORMAT") {
			switch string(c) {
			case "d":
				// datetime
				fmt.Printf("%s %s %2s %s %s ", f[0], f[1], f[2], f[3], f[4])

			case "f":
				// flag
				fmt.Printf("%s ", flag[tz[0]])

			case "c":
				// city/time zone
				fmt.Printf("%s%*s",
					label,
					maxLen-unicodeLen(label),
					" ",
				)
			}
		}
		fmt.Println()
	}
}
