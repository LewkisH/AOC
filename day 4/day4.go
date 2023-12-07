package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	/* 	Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	   	Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
	   	Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
	   	Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
	   	Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
	   	Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
	*/
	fileB, _ := os.ReadFile("test.txt")
	file := strings.Split(string(fileB), "\n")
	//points := 0
	var leftSide string
	//var ans int
	finalAns := 206
	var rightSide string
	var colonIndex int
	var nums [206]int
	for lineIndex, line := range file {
		//	for j := 0; j < nums[lineIndex]+1; j++ {
		counta := 0
		for i, ch := range line {
			if ch == ':' {
				colonIndex = i
			}
			if ch == '|' {
				leftSide = line[colonIndex+1 : i]
				rightSide = line[i+1:]
				r := regexp.MustCompile(`\d+`)
				numbersLeft := r.FindAllString(leftSide, -1)
				numbersRight := r.FindAllString(rightSide, -1)
				for _, num := range numbersLeft {
					for _, numMatch := range numbersRight {
						if num == numMatch {
							counta++
							nums[lineIndex+counta] += 1*nums[lineIndex] + 1
							fmt.Println("match:", num)
							fmt.Println(counta)
						}

					}
				}

				break
			}

		}
		/* if counta > 0 {
			fmt.Println("matches:", counta)
			ans = 1
			for i := 0; i < counta-1; i++ {
				ans *= 2
			}
			finalAns += ans
		} */
		/* 		fmt.Println(leftSide + "X" + rightSide)
		 */

	}
	for _, value := range nums {
		finalAns += value
	}
	fmt.Println(finalAns)
}
