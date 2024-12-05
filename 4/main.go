package main

import(
	"fmt"
	"os"
	"bufio"
	"log"
)

func main(){
	fmt.Println("aoc 4. 2024")

	file, err := os.Open("example.txt")
	//file, err := os.Open("actual.txt")

	if err != nil {
		log.Fatal("could not read file", err)
	}

	scanner := bufio.NewScanner(file)

	wordMatrix := [][]byte{}

	for scanner.Scan() {
		wordMatrix = append(wordMatrix, []byte(scanner.Text()))
	}

	//fmt.Println(wordMatrix)
	
	//tactic, find all x-s and check around them for xmas strings in all possible directions

	x := byte('X')

	for i := range wordMatrix{
		for j := range wordMatrix[i] {
			if wordMatrix[i][j] == x {
				checkForXmas(i, j, wordMatrix)
			}
		}
	}
}

func checkForXmas(i, j int, wordMatrix [][]byte){
}

func checkVertical(){
}

func checkHorizontal(){
}

func checkDiagonal(){
}
