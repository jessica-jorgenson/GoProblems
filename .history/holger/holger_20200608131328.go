package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertToIntArr(stringArr []string) []int {
	var returnArr []int
	for _, item := range stringArr {
		integer, _ := strconv.Atoi(item)
		returnArr = append(returnArr, integer)
	}
	return returnArr
}

type characters struct {
	letter   string
	location []int
}

func charJoin(charactersArr []characters) string {
	var returnStr string
	for _, char := range charactersArr {
		returnStr += char.letter
	}
	returnStr = strings.Replace(returnStr, "\n", "", -1)
	return returnStr
}

func reverse(characterArr []characters) []characters {
	returnArr := make([]characters, len(characterArr))
	for i := 0; i < len(characterArr)/2; i++ {
		j := len(characterArr) - i - 1
		returnArr[i], returnArr[j] = characterArr[j], characterArr[i]
	}

	return returnArr
}

func formatCoords(words map[string][][]int, maxHeight int, maxWidth int) string {
	var returnString string
	var location []int
	for _, item := range words {
		if len(item) == 0 {
			returnString += "-1 -1\n"
			continue
		} else if len(item) == 1 {
			returnString += strconv.Itoa(item[0][0]+1) + " " + strconv.Itoa(item[0][1]+1) + "\n"
			continue
		}

		var findLeftMost [][]int
		findTopMost := maxHeight
		for _, coor := range item {
			if coor[0] < findTopMost {
				findTopMost = coor[0]
			}
		}
		for _, coor := range item {
			if coor[0] == findTopMost {
				findLeftMost = append(findLeftMost, coor)
			}
		}
		if len(findLeftMost) > 1 {
			var leftNum = maxWidth
			var leftIdx = len(findLeftMost) - 1
			for idx, coor := range findLeftMost {
				if coor[1] < leftNum {
					leftNum = coor[1]
					leftIdx = idx
				}
			}
			location = findLeftMost[leftIdx]
		} else {
			location = findLeftMost[0]
		}

		returnString += strconv.Itoa(location[0]+1) + " " + strconv.Itoa(location[1]+1) + "\n"

	}
	return returnString
}

func allChars(grid []string) []characters {
	var charactersArr []characters
	for row, gridRow := range grid {
		for col, letter := range gridRow {
			var item characters
			item.letter = string(letter)
			item.location = []int{row, col}
			charactersArr = append(charactersArr, item)
		}

	}
	return charactersArr
}

func fillGrid(scanner *bufio.Scanner, grid []string, maxHeight int) []string {
	for i := 0; i < maxHeight; i++ {
		scanner.Scan()
		grid[i] = strings.ToUpper(scanner.Text()) + "\n"
	}
	return grid
}

func getLocation(charactersArr []characters, word string) []int {
	for i := range charactersArr {
		if charJoin(charactersArr[i:i+len(word)]) == word {
			return charactersArr[i].location
		}
	}
	return []int{-1, -1}
}

func newSearch(scanner *bufio.Scanner, directions map[string]int) ([]string, []characters, map[string][][]int) {
	scanner.Scan()
	dimensions := convertToIntArr(strings.Fields(scanner.Text()))
	letterGrid := fillGrid(scanner, make([]string, dimensions[0]), dimensions[0])
	charactersArr := allChars(letterGrid)
	wordLocations := make(map[string][][]int)
	return letterGrid, charactersArr, wordLocations
}

func getWordsToFind(scanner *bufio.Scanner) []string {
	scanner.Scan()
	numWords, _ := strconv.Atoi(scanner.Text())
	wordsToFind := make([]string, numWords)
	for i := 0; i < numWords; i++ {
		scanner.Scan()
		wordsToFind[i] = strings.ToUpper(strings.TrimSuffix(scanner.Text(), "\n"))
	}

	return wordsToFind
}

func trimNewLines(charactersArr []characters) []characters {
	for i := range searchGrid {
		fmt.Println()
	}
}

func getLooking(grid []string, words []string, charactersArr []characters, directions map[string]int, wordLocations map[string][][]int) map[string][][]int {
	searchGrid := make(map[string][]characters)

	for wordDirection, directionNum := range directions {
		for x := 0; x < len(grid[0]); x++ {
			for i := x; i < len(charactersArr); i += len(grid[0]) + directionNum {
				searchGrid[wordDirection] = append(searchGrid[wordDirection], charactersArr[i])
			}
		}
	}
	searchGrid["right"] = trimNewLines(charactersArr)
	searchGrid["left"] = reverse(charactersArr)
	searchGrid["up"] = reverse(searchGrid["down"])
	searchGrid["up and left diag"] = reverse(searchGrid["down and right diag"])
	searchGrid["up and right diag"] = reverse(searchGrid["down and left diag"])

	for i, chars := range searchGrid {
		fmt.Println(i)
		fmt.Println(chars)
		var theString = charJoin(chars)
		fmt.Println(theString)
		for _, word := range words {
			if strings.Contains(theString, word) {
				wordLocations[word] = append(wordLocations[word], getLocation(chars, word))
			}
		}
	}

	return wordLocations
}

func makeWordLoc(wordsToFind []string) map[string][][]int {
	wordLocations := make(map[string][][]int)
	for _, word := range wordsToFind {
		wordLocations[word] = [][]int{}
	}

	return wordLocations
}

func main() {
	inputFile, err := os.Open("holger.in")
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)

	outputFile, err := os.Create("holger.out") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
	}
	defer outputFile.Close()
	scanner.Scan()

	directions := map[string]int{
		"down":                0,
		"down and left diag":  -1,
		"down and right diag": 1,
	}

	letterGrid := []string{}
	wordsToFind := []string{}
	charactersArr := []characters{}
	wordLocations := make(map[string][][]int)
	toOutput := ""
	cases := 0
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			cases++
			if cases > 1 {
				toOutput = "\n"
			}
			letterGrid, charactersArr, wordLocations = newSearch(scanner, directions)
			wordsToFind = getWordsToFind(scanner)
			wordLocations = makeWordLoc(wordsToFind)
			wordLocations = getLooking(letterGrid, wordsToFind, charactersArr, directions, wordLocations)
			toOutput += formatCoords(wordLocations, len(letterGrid), len(letterGrid[0]))
			outputFile.WriteString(toOutput)
			outputFile.Sync()
		}
	}
}
