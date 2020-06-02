package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Team struct {
	submittedProblems map[int][]int
}

func makeScoreBoard() string {

}

func main() {
	inputFile, err := os.Open("scoreboard.in")
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)

	outputFile, err := os.Create("scoreboard.out") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
	}
	defer outputFile.Close()
	scanner.Scan()

	cases, _ := strconv.Atoi(scanner.Text())
	currentCase := 1
	board := make(map[int]Team)
	scanner.Scan()
	for scanner.Scan() {
		info := strings.Fields(scanner.Text())
		if info[len(info)-1] == "I" || info[len(info)-1] == "C" {

		}
	}

}
