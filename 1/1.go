package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var highest  = 0;

func checkHighest(newScore int){

	if newScore > highest {
		highest = newScore
	}

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
		if line := scanner.Text();line == "" {
			checkHighest(total)
			total = 0
		} else {
			lineNumber,_ := strconv.ParseInt(line, 10, 32)
			total = total + int(lineNumber)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}


	println(highest)
}
