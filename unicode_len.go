package main

func unicodeLen(s string) int {
	size := 0
	for _, c := range s {
		if len(string(c)) > 1 {
			size += len(string(c)) - 1
		} else {
			size++
		}
	}
	return size
}
