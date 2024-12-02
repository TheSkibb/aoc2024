package main

import(
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main(){
	fmt.Println("aoc 2. 2024")

	//file, err := os.Open("example.txt")	
	file, err := os.Open("actual.txt")	

	if err != nil {
		log.Fatal("cant open file", err)
	}

	leftCol := []int{}
	rightCol := []int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currLine := scanner.Text()

		colIndex := strings.Index(currLine, " ")

		leftNum, _ := strconv.Atoi(string([]rune(currLine)[:colIndex]))
		rightNum, _ := strconv.Atoi(string([]rune(currLine)[colIndex+3:]))
		
		leftCol = append(leftCol, leftNum)
		rightCol = append(rightCol, rightNum)
	}

	//fmt.Println(leftCol)
	//fmt.Println(rightCol)

	similarityScore := 0

	for _, leftVal := range leftCol{

		simCount := 0

	 	for _, rightVal := range rightCol {
			if leftVal == rightVal {
				simCount++
			}
		}

		similarityScore += simCount * leftVal
	}

	fmt.Println(similarityScore)
}
