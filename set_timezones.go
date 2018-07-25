package main

import (
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	mtzdateTimezones []string
)

// https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
/*
  if no MTZDATE_TIMEZONES,
    set MTZDATE_TIMEZONES to UTC

  set mtzdateTimezones to MTZDATE_TIMEZONES (split on ",")
*/
func setTimezones() {
	re := regexp.MustCompile(".*/")

	if _tz, ok := os.LookupEnv("MTZDATE_TIMEZONES"); !ok || _tz == "" {
		err := os.Setenv("MTZDATE_TIMEZONES", utc)
		die(err)
	}

	mtzdateTimezones = strings.Split(os.Getenv("MTZDATE_TIMEZONES"), ",")

	for i, kv := range mtzdateTimezones {
		// unpack
		tz := strings.Split(kv, ":")

		// Missing tz[1] means tz[0] is actual time zone
		if len(tz) == 1 {
			tz = append(tz, tz[0])
		}

		if _, err := time.LoadLocation(tz[1]); err != nil {
			// Fall back to UTC on bogus time zone
			tz[1] = utc
		}

		tz[0] = re.ReplaceAllString(tz[0], "")

		// repack
		mtzdateTimezones[i] = strings.Join(tz, ":")
	}
}
