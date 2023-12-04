package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
*/
type box [3]string

func main() {
	fileB, _ := os.ReadFile("day3.txt")
	file := strings.Split(string(fileB), "\n")
	symbols := "@#$%+-=/&*"
	var ans int
	for lineIndex, line := range file {
		var box box
		numba := regexp.MustCompile(`\d+`)
		numbers := numba.FindAllString(line, -1)
		numIndexes := numba.FindAllStringIndex(line, -1)

		for j, number := range numbers {
			end := 3
			start := 0
			if lineIndex == 0 {
				start = 1
			}

			for i := start; i < end; i++ {
				if lineIndex == len(file)-1 {
					end--
				}
				if numIndexes[j][0] == 0 {
					numIndexes[j][0]++
				}
				if numIndexes[j][1] > len(line)-1 {
					numIndexes[j][1]--
				}
				box[i] = file[lineIndex-1+i][numIndexes[j][0]-1 : numIndexes[j][1]+1]

			}
			for _, boxline := range box {
				if strings.ContainsAny(symbols, boxline) {
					intNum, _ := strconv.Atoi(number)
					ans += intNum
					break
				}
			}

			//fmt.Println(box)
			//fmt.Println(number)

		}

		//fmt.Println(file[0])
		//fmt.Println(numbers)
		//fmt.Println(numIndexes)
	}
	fmt.Println(ans)
}
