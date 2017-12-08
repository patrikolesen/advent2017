package advent2017

import (
	"fmt"
	"strconv"
	"strings"
)

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func Advent2017_21(lines []string) {
	var min, max int
	min = 99999
	max = 0
	for _, e := range strings.Split(lines[0], "\t") {
		val, _ := strconv.Atoi(e)
		if val < min {
			min = val
		} else if val > max {
			max = val
		}
	}
}
