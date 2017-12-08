package advent2017

func Advent2017_1(lines []string, useHalf bool) int {
	var (
		next int
		sum  int
	)
	sum = 0

	myString := lines[0]
	if useHalf {
		next = len(myString) / 2
	} else {
		next = 1
	}

	for i := 0; i < len(myString); i++ {
		if myString[i] == myString[(i+next)%len(myString)] {
			sum += int(myString[i] - '0')
		}
	}
	return sum
}
