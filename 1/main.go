package main

import(
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
	"sort"
	"strings"
)

func main(){
	fmt.Println("aoc 1. dec")

	//file, err := os.Open("example.txt")
	file, err := os.Open("actual.txt")
	
	if err != nil {
		log.Fatal("could not open file")
	}

	leftCol := []int{}
	rightCol := []int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currline := scanner.Text()
		if currline == "" {
			continue
		}

		splitcolumn := strings.Index(currline, " ")
		
		leftNum, _ := strconv.Atoi(string([]rune(currline)[0:splitcolumn]))
		leftCol = append(leftCol, leftNum)

		rightTemp := strings.TrimSpace(string([]rune(currline)[splitcolumn:]))

		rightNum, _ := strconv.Atoi(rightTemp)
		rightCol = append(rightCol, rightNum)
	}

	sort.Ints(leftCol)
	sort.Ints(rightCol)

	var totalDist int

	for i := range leftCol{
		totalDist += Abs(leftCol[i] - rightCol[i])
		fmt.Println(leftCol[i], "-", rightCol[i], "total:", totalDist)
	}

	fmt.Println(totalDist)
}

func Abs(x int) int{
	if x < 0 {
		return -x
	}
	return x
}
