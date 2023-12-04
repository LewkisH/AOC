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
	re := regexp.MustCompile(`Game (\d+)`)
	red := regexp.MustCompile(`(\d+) red`)
	green := regexp.MustCompile(`(\d+) green`)
	blue := regexp.MustCompile(`(\d+) blue`)
	fileB, _ := os.ReadFile("day2.txt")
	file := strings.Split(string(fileB), "\n")
	ans := 0
lineparse:
	for _, line := range file {
		gameNrS := re.FindAllStringSubmatch(line, -1)[0][1]
		//	reds1 := red.FindAllStringSubmatch(line, -1)
		//	fmt.Println(reds1[len(reds1)-1])
		reds := red.FindAllStringSubmatch(line, -1)

		for _, nr := range reds {
			intNr, _ := strconv.Atoi(nr[1])
			if intNr > 12 {
				continue lineparse
			}
		}
		greens := green.FindAllStringSubmatch(line, -1)
		for _, nr := range greens {
			intNr, _ := strconv.Atoi(nr[1])
			if intNr > 13 {
				continue lineparse
			}
		}
		blues := blue.FindAllStringSubmatch(line, -1)
		for _, nr := range blues {
			intNr, _ := strconv.Atoi(nr[1])
			if intNr > 14 {
				continue lineparse
			}
		}
		gameNr, _ := strconv.Atoi(gameNrS)
		ans += gameNr

	}
	fmt.Println(ans)
}
