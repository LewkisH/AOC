package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//scanner := bufio.NewScanner()
	numbers := "0123456789"
	var ans int
	fileB, _ := os.ReadFile("day1.txt")
	file := strings.Split(string(fileB), "\n")
	for _, line := range file {
		nr := ""
		for _, ch := range line {

			if strings.Contains(numbers, string(ch)) {
				nr += string(ch)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if strings.Contains(numbers, string(line[i])) {
				nr += string(line[i])
				break
			}
		}
		intNr, _ := strconv.Atoi(nr)
		ans += intNr
	}
	fmt.Println(ans)
}
