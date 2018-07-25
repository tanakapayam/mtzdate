package main

import (
	"fmt"
	"strconv"
	"strings"
)

// nolint
func reformatTime(f []string) []string {
	hms := strings.Split(f[3], ":")
	h, err := strconv.Atoi(hms[0])
	die(err)

	if f[4] != utc {
		if _, ok := workday[f[0]]; ok {
			// MTZDATE_GREEN_HOURS="8-17" -> greenHours=[[8,17]]
			for i := range greenHours {
				if greenHours[i][0] <= h && h < greenHours[i][1] {
					f[3] = green(f[3])
				}
			}

			// MTZDATE_YELLOW_HOURS="7-8,17-18" -> yellowHours=[[7, 8], [17, 18]]
			for i := range yellowHours {
				if yellowHours[i][0] <= h && h < yellowHours[i][1] {
					f[3] = yellow(f[3])
				}
			}

			// MTZDATE_FAINT_HOURS="0-7,22-24" -> faintHours=[[0, 7], [22, 24]]
			for i := range faintHours {
				if faintHours[i][0] <= h && h < faintHours[i][1] {
					f[3] = faint(f[3])
				}
			}
		} else {
			// MTZDATE_FAINT_HOURS="0-7,22-24" -> faintHours=[[0, 7], [22, 24]]
			for i := range faintHours {
				if faintHours[i][0] <= h && h < faintHours[i][1] {
					f[3] = faint(f[3])
				}
			}
		}
	}

	// pad timezone
	f[4] = fmt.Sprintf("%-5s", f[4])

	// drop year
	f = f[:5]

	return f
}
