package main

//package main

//package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This function converts string array to integer array
func convertArr(stringArr []string, subType int) []int {
	var returnArr []int
	for _, item := range stringArr[:len(stringArr)-1] {
		integer, _ := strconv.Atoi(item)
		returnArr = append(returnArr, integer)
	}
	returnArr = append(returnArr, subType)
	return returnArr
}

func readInt(stringNum string) int {
	number, err := strconv.Atoi(stringNum)
	if err != nil {
		fmt.Println("Error in read_integer:", err)
	}
	return number
}

func convertSubmission(submissionType string) int {
	if submissionType == "C" {
		return 1
	} else if submissionType == "I" {
		return 0
	}
	return 2
}

func getScore(score int, currentScore int, submissionType int) int {
	if submissionType == 0 {
		score = 10
	} else if submissionType == 1 {
		score = 20
	} else {
		score = 0
	}
	return score
}

//
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

	cases := readInt(scanner.Text())
	currentCase := 1
	//board := make(map[int][]int)
	scanner.Scan()
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			currentCase++
			break
		}
		info := strings.Fields(scanner.Text())
		var subType = convertSubmission(info[len(info)-1])
		if subType < 2 {
			infoArr := convertArr(info, subType)
			fmt.Println(infoArr)
		}
	}
	fmt.Println(cases)
}
