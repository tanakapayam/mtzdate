package main

import (
	"os"
	"strings"
)

func setWorkhours() {
	workday = make(map[string]bool)

	_, set := os.LookupEnv("MTZDATE_WORKDAYS")
	if !set {
		err := os.Setenv("MTZDATE_WORKDAYS", mtzdateWorkdays)
		die(err)
	}

	if os.Getenv("MTZDATE_WORKDAYS") != "" {
		// MTZDATE_GREEN_HOURS="8-17" -> greenHours=[[8,17]]
		greenHours = splitEnvIntoArray("MTZDATE_GREEN_HOURS", mtzdateGreenHours, greenHours)

		// MTZDATE_YELLOW_HOURS="7-8,17-18" -> yellowHours=[[7, 8], [17, 18]]
		yellowHours = splitEnvIntoArray("MTZDATE_YELLOW_HOURS", mtzdateYellowHours, yellowHours)

		// MTZDATE_FAINT_HOURS="0-7,22-24" -> faintHours=[[0, 7], [22, 24]]
		faintHours = splitEnvIntoArray("MTZDATE_FAINT_HOURS", mtzdateFaintHours, faintHours)
	}

	for _, v := range strings.Split(os.Getenv("MTZDATE_WORKDAYS"), ",") {
		workday[v] = true
	}
}
