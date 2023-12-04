package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
467..114..	[0 2] [5 7]
...*......	[2 4]
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
	//nums := "1234567890"
	//symbols := "@#$%+-=/&*"
	ans := 0
	for lineIndex, line := range file {
		var box box

		gear := regexp.MustCompile(`\*`)
		gearIndex := gear.FindAllStringIndex(line, -1)
		gears := gear.FindAllString(line, -1)
		for j, _ := range gears {
			end := 3
			start := 0
			if lineIndex == 0 {
				start = 1
			}

			for i := start; i < end; i++ {
				if lineIndex == len(file)-1 {
					end--
				}
				if gearIndex[j][0] == 0 {
					gearIndex[j][0] += 3
				}
				if gearIndex[j][1] > len(line)-1 {
					gearIndex[j][1] -= 3
				}
				box[i] = file[lineIndex-1+i][gearIndex[j][0]-3 : gearIndex[j][1]+3]

			}
			ratio := 1
			count := 0
			for _, boxline := range box {
				numba := regexp.MustCompile(`\d+`)
				numbers := numba.FindAllString(boxline, -1)
				numIndexes := numba.FindAllStringIndex(boxline, -1)

				for i, number := range numbers {
					numInt, _ := strconv.Atoi(number)

					if numIndexes[i][0] >= 2 && numIndexes[i][0] <= 4 {
						fmt.Println(number)
						count++
						ratio *= numInt
					} else if numIndexes[i][1]-1 >= 2 && numIndexes[i][1]-1 <= 4 {
						fmt.Println(number)
						ratio *= numInt
						count++

					}

				}

			}
			if count == 2 {
				ans += ratio
			}
			fmt.Println(count, ratio)
			fmt.Println(ans)

		}

		fmt.Println(box)
		//fmt.Println(number)

	}

	//fmt.Println(file[0])
	//fmt.Println(numbers)
	//fmt.Println(numIndexes)
}

//fmt.Println(ans)}
