package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("aoc 3 2024")

	//file, err := os.Open("./example.txt")
	file, err := os.Open("./actual.txt")

	if err != nil {
		log.Fatal(err.Error())
	}

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		currLine := scanner.Text()
		total += extractMults(currLine)
	}

	fmt.Println("total", total)
}

func extractMults(s string) int {
	r, _ := regexp.Compile(`mul\(\d*,\d*\)`)
	matches := r.FindAllString(s, -1)

	total := 0

	for _, mul := range matches {
		total += parseMul(mul)
	}

	return total
}

func parseMul(s string) int {

	commaIndex := strings.Index(s, ",")

	leftNum, _ := strconv.Atoi(string([]byte(s)[4:commaIndex]))

	rightNum, _ := strconv.Atoi(string([]byte(s)[commaIndex+1 : len(s)-1]))

	mul := leftNum * rightNum

	fmt.Println(s, leftNum, rightNum, mul)

	return mul
}
