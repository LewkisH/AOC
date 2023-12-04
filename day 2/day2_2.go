package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//12 red cubes, 13 green cubes, and 14 blue cubes
	red := regexp.MustCompile(`(\d+) red`)
	green := regexp.MustCompile(`(\d+) green`)
	blue := regexp.MustCompile(`(\d+) blue`)
	fileB, _ := os.ReadFile("day2.txt")
	file := strings.Split(string(fileB), "\n")
	ans := 0

	for _, line := range file {
		reds := red.FindAllStringSubmatch(line, -1)
		maxRed := 0
		for _, nr := range reds {
			intNr, _ := strconv.Atoi(nr[1])
			if intNr > maxRed {
				maxRed = intNr
			}
		}
		maxGreen := 0
		greens := green.FindAllStringSubmatch(line, -1)
		for _, nr := range greens {
			intNr, _ := strconv.Atoi(nr[1])
			if intNr > maxGreen {
				maxGreen = intNr
			}
		}
		maxBlue := 0
		blues := blue.FindAllStringSubmatch(line, -1)
		for _, nr := range blues {
			intNr, _ := strconv.Atoi(nr[1])
			if intNr > maxBlue {
				maxBlue = intNr
			}
		}
		ans += maxBlue * maxGreen * maxRed

	}
	fmt.Println(ans)
}
