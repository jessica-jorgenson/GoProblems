package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// These first three functions are just basic functions to help me finish this problem

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func replaceAtIndex(in string, r rune, idx int) string {
	out := []rune(in)
	out[idx] = r
	return string(out)
}

func allMatch(arr []int) bool {
	firstItem := arr[0]
	for _, item := range arr {
		if item != firstItem {
			return false
		}
	}
	return true
}

// This function returns whether or not the given array goes above 4 (we change into the 4s later since Torsten doesn't know about 4s)
func ignoreAbove3(arr []int) bool {
	badNums := []int{4, 5, 6, 7, 8, 9}
	for _, item := range arr {
		for _, num := range badNums {
			if item == num {
				return true
			}
		}
	}
	return false
}

// This function recursively gets all the different permutations of a sum (e.x. 42 to 24)
func getPermutations(theStr string, index int, allPerms []string) []string {
	if index == len(theStr) {
		return allPerms
	}

	for i := index; i < len(theStr); i++ {
		r := []rune(theStr)
		r[index], r[i] = r[i], r[index]
		newStr := string(r)
		if !contains(allPerms, newStr) {
			allPerms = append(allPerms, newStr)
		}
		allPerms = getPermutations(newStr, index+1, allPerms)
	}
	return allPerms
}

// After we find all the actual correct combinations, we do "Torsten's Rule" to recursively find all the possibilites of how 1s and 4s could fit
func replaceWithFours(theStr string, index int, currentPossibilities []string) []string {
	if index == len(theStr) {
		return currentPossibilities
	}

	if theStr[index] == '1' {
		theStr = replaceAtIndex(theStr, '4', index)
		if !contains(currentPossibilities, theStr) {
			currentPossibilities = append(currentPossibilities, theStr)
		}
		currentPossibilities = replaceWithFours(theStr, index+1, currentPossibilities)

		theStr = replaceAtIndex(theStr, '1', index)
		currentPossibilities = replaceWithFours(theStr, index+1, currentPossibilities)
	} else {
		currentPossibilities = replaceWithFours(theStr, index+1, currentPossibilities)
	}
	return currentPossibilities
}

// This simple function just sends any to-be-summed numbers that contain a 1 into the replaceWithFours function
func findOnes(arr []string) []string {
	var allPoss []string
	for _, item := range arr {
		if strings.Contains(item, "1") {
			var possibilities []string
			possibilities = replaceWithFours(item, 0, possibilities)
			for _, newPoss := range possibilities {
				allPoss = append(allPoss, newPoss)
			}
		}
		allPoss = append(allPoss, item)
	}
	return allPoss
}

// This function recursively finds any number combination that sums to the given number
func findCombos(allSums []string, arr []int, index int, givenNum int, reducedNum int) []string {

	if reducedNum < 0 {
		return allSums
	}

	if reducedNum == 0 && !ignoreAbove3(arr) {
		valuesText := []string{}
		for k := 0; k < index; k++ {
			number := arr[k]
			text := strconv.Itoa(number)
			valuesText = append(valuesText, text)
		}
		result := strings.Join(valuesText, "")
		allSums = append(allSums, result)
		if !allMatch(arr[:index]) {
			var allPerms []string
			allPerms = getPermutations(result, 0, allPerms)
			for _, item := range allPerms {
				if !contains(allSums, item) {
					allSums = append(allSums, item)
				}
			}
		}
		return allSums
	}

	prev := 0
	if index == 0 {
		prev = 1
	} else {
		prev = arr[index-1]
	}

	for i := prev; i < givenNum+1; i++ {
		arr[index] = i

		allSums = findCombos(allSums, arr, index+1, givenNum, reducedNum-i)
	}
	return allSums
}

func main() {
	// After opening/preparing the files..
	inputFile, err := os.Open("counting.in")
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)

	outputFile, err := os.Create("counting.out") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer outputFile.Close()

	for scanner.Scan() {
		toSum, _ := strconv.Atoi(scanner.Text())
		digitsArr := make([]int, toSum)
		var sumsArr []string
		// Combos is all the legal combinations (minus the combinations that use numbers higher than 3)
		combos := findCombos(sumsArr, digitsArr, 0, toSum, toSum)
		// Possibilities is when we replace the 1s with 4s in different places and such
		getPoss := findOnes(combos)

		outputFile.WriteString(fmt.Sprintf("%d\n", len(getPoss)))
		outputFile.Sync()

	}
}
