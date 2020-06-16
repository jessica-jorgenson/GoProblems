package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type gridNode struct {
	weight   int
	location []int
	adjNodes []gridNode
	isEnd    bool
}

func (n gridNode) getX() int {
	return n.location[1]
}

func (n gridNode) getY() int {
	return n.location[0]
}

func printGrid(grid [][]gridNode) {
	for _, row := range grid {
		rowString := ""
		for _, item := range row {
			rowString = rowString + " " + strconv.Itoa(item.weight) + " " + strings.Replace(fmt.Sprint(item.location), " ", " ", -1)
		}
		fmt.Println(rowString)
	}
}

func printPath(path []gridNode) {
	rowString := ""
	for _, item := range path {
		rowString = rowString + " " + strconv.Itoa(item.weight) + " " + strings.Replace(fmt.Sprint(item.location), " ", " ", -1)
	}
	fmt.Println(rowString)
}

// This function converts string array to integer array
func convertToIntArr(stringArr []string) []int {
	var returnArr []int
	for _, item := range stringArr {
		integer, _ := strconv.Atoi(item)
		returnArr = append(returnArr, integer)
	}
	return returnArr
}

func getAdjNodes(grid [][]gridNode, currentNode gridNode) []gridNode {
	moveY := []int{0, 1, -1}
	var adjNodes []gridNode
	maxHeight := len(grid)
	nextCol := currentNode.getX() + 1
	currentY := currentNode.getY()

	for _, yDiff := range moveY {
		adjNode := grid[getNewY(maxHeight, currentY+yDiff)][nextCol]
		adjNodes = append(adjNodes, adjNode)
	}

	return adjNodes
}

func getNewY(gridHeight int, maybeY int) int {
	if maybeY < 0 {
		maybeY = gridHeight - 1
	} else if maybeY == gridHeight {
		maybeY = 0
	}
	return maybeY
}

func getPathWeight(path []gridNode) int {
	totalWeight := 0
	for _, node := range path {
		totalWeight += node.weight
	}
	return totalWeight
}

// This function returns if a num integer contains a specified int
func contains(arr [][]gridNode, num int) bool {
	for arrindex := range arr {
		for _, b := range arr[arrindex] {
			if b.weight == num {
				return true
			}
		}
	}
	return false
}

func getRows(path []gridNode) string {
	rowString := ""
	for i, item := range path {
		rowString = rowString + strconv.Itoa(item.getY()+1)
		if i < len(path)-1 {
			rowString = rowString + " "
		}
	}
	return rowString
}

func findPath(grid [][]gridNode, currentNode gridNode, allPaths *[][]gridNode, path []gridNode) {
	path = append(path, currentNode)
	if currentNode.isEnd {
		if len(path) == len(grid[0]) {
			*allPaths = append(*allPaths, path)
		}
	} else {
		currentNode.adjNodes = getAdjNodes(grid, currentNode)
		for _, node := range currentNode.adjNodes {
			findPath(grid, node, allPaths, path)
		}
	}
}

func makeGrid(scanner *bufio.Scanner) [][]gridNode {
	dims := convertToIntArr(strings.Fields(scanner.Text()))
	scanner.Scan()
	grid := make([][]gridNode, dims[0])
	heightIndex, widthIndex := 0, 0
	for i := 0; i < dims[0]; i++ {
		grid[i] = make([]gridNode, dims[1])
	}
	for ok := true; ok; ok = contains(grid, 0) {
		line := convertToIntArr(strings.Fields(scanner.Text()))
		for _, item := range line {
			var currentNode gridNode
			currentNode.weight = item
			currentNode.location = []int{heightIndex, widthIndex}
			if widthIndex < dims[1]-1 {
				currentNode.isEnd = false
			} else {
				currentNode.isEnd = true
			}
			grid[heightIndex][widthIndex] = currentNode
			widthIndex++
			if widthIndex == dims[1] {
				heightIndex++
				widthIndex = 0
			}
		}
		if heightIndex != dims[0] {
			scanner.Scan()
		}

	}
	return grid
}

func findMinWeight(paths [][]gridNode) []gridNode {
	var smallestWeight []gridNode
	for _, path := range paths {
		currentWeight := getPathWeight(path)
		smallest := getPathWeight(smallestWeight)
		if getPathWeight(path) < getPathWeight(smallestWeight) || getPathWeight(smallestWeight) == 0 {
			fmt.Println(smallest)
			smallestWeight = path
			fmt.Println(currentWeight)
			printPath(path)
			fmt.Println("")
		}
	}
	return smallestWeight
}

func main() {
	// After opening/preparing the files..
	inputFile, err := os.Open("tsp.in")
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)

	outputFile, err := os.Create("tsp.out") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer outputFile.Close()

	// We interate through the file
	for scanner.Scan() {
		grid := makeGrid(scanner)
		printGrid(grid)
		var emptyPath []gridNode
		var allPaths [][]gridNode
		for _, row := range grid {
			findPath(grid, row[0], &allPaths, emptyPath)
		}
		minWeightPath := findMinWeight(allPaths)

		pathRowArr := getRows(minWeightPath)
		pathWeight := getPathWeight(minWeightPath)
		fmt.Println(pathRowArr)
		outputFile.WriteString(pathRowArr + "\n")
		outputFile.WriteString(strconv.Itoa(pathWeight) + "\n")
		outputFile.Sync()

	}
}
