package main

import (
	"os"
	"strconv"
	"strings"
)

func splitEnvIntoArray(env string, def string, array [][]int) [][]int {
	_, set := os.LookupEnv(env)
	if !set {
		err := os.Setenv(env, def)
		die(err)
	}

	for _, r := range strings.Split(os.Getenv(env), ",") {
		var v []int
		for _, _u := range strings.Split(r, "-") {
			_v, err := strconv.Atoi(_u)
			die(err)
			v = append(v, _v)
		}
		array = append(array, v)
	}

	return array
}
