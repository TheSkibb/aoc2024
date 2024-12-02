package main

import(
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
	"log"
)

func main(){
	fmt.Println("aoc 2 2024")

	file, err := os.Open("example.txt")
	//file, err := os.Open("actual.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		currLine := scanner.Text()
		count += checkline(currLine)
	}	
	
	fmt.Println(count, "safe reports")
}

// returns 0 or 1
func checkline(line string) int {
	lineSplit := strings.Split(line, " ")

	inc := 1

	for i := range lineSplit{

		currVal, _ := strconv.Atoi(lineSplit[i])

		if i == 0 {
			//determine if increasing or decreasing
			nextVal, _ := strconv.Atoi(lineSplit[i + 1])
			if nextVal < currVal {
				inc = -1
			}
			continue
		}
		prevVal, _ := strconv.Atoi(lineSplit[i - 1])

		if inc == 1 && currVal < prevVal {
			fmt.Println("decrease in increasing function", prevVal, currVal, lineSplit)
			return 0
		}

		if inc == -1 && currVal > prevVal {
			fmt.Println("increase in decreasing function", prevVal, currVal, lineSplit)
			return 0
		}
		
		diff := Abs(currVal - prevVal)
		//fmt.Println("difference", diff)
		
		if 1 > diff || diff > 3 {
			fmt.Println("step wrong size", diff, lineSplit)
			return 0
		}
		
	}

	fmt.Println(lineSplit)
	return 1
}

func Abs(x int) int{
	if x < 0 {
		return -x
	}
	return x
}
