package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var highest = []int{}

// func checkHighest(newScore int){

// 	for _, score := range highest {
//         if newScore > score {
// 			score = newScore
// 			break
// 		}
//     }

// }

func sortArray() {
	sort.Ints(highest[:])
}

func totalHighest() int {
	total := 0
	for index := 1; index < 4; index++ {
		total = total + highest[len(highest)-index]
	}
	return total
}

func main() {
	file, err := os.Open("calories.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		if line := scanner.Text(); line == "" {
			highest = append(highest, total)
			total = 0
		} else {
			lineNumber, _ := strconv.ParseInt(line, 10, 32)
			total = total + int(lineNumber)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	sortArray()
	//print(highest)
	println(totalHighest())
}
