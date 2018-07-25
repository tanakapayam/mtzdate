package main

import (
	"strings"
)

func updateFlagMap(array []string) {
	for _, kv := range array {
		_kv := strings.Split(kv, ":")

		if len(_kv) == 1 {
			if countryCode[_kv[0]] != "" {
				flag[_kv[0]] = flag[countryCode[_kv[0]]]
			}
		} else {
			if countryCode[_kv[1]] != "" {
				flag[_kv[0]] = flag[countryCode[_kv[1]]]
			} else {
				flag[_kv[0]] = flag[_kv[1]]
			}
		}

		if flag[_kv[0]] == "" {
			flag[_kv[0]] = "  "
		}
	}
}
