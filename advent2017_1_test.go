package advent2017

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func getLines(filename string) []string {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func TestAdvent2017_1(t *testing.T) {
	sum := Advent2017_1(getLines("/mnt/tools/go/src/advent2017/input_1.txt"), false)
	fmt.Println(sum)
}

func TestAdvent2017_12(t *testing.T) {
	sum := Advent2017_1(getLines("/mnt/tools/go/src/advent2017/input_1.txt"), true)
	fmt.Println(sum)
}

func TestAdvent2017_21(t *testing.T) {
	Advent2017_21(getLines("/mnt/tools/go/src/advent2017/input_2.txt"))
}
