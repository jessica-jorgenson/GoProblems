package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// This function converts string array to integer array
func convertToIntArr(stringArr []string) []int {
	var returnArr []int
	for _, item := range stringArr {
		integer, _ := strconv.Atoi(item)
		returnArr = append(returnArr, integer)
	}
	return returnArr
}

// This function returns if a num integer contains a specified int
func contains(arr [][]int, num int) bool {
	for arrindex := range arr {
		for _, b := range arr[arrindex] {
			if b == num {
				return true
			}
		}
	}
	return false
}

// This function determines if the given point is already in the given array
func alreadyIn(arr [][]int, point []int) bool {
	for _, item := range arr {
		if reflect.DeepEqual(item, point) {
			return true
		}
	}
	return false
}

// This makes an empty grid given height/width
func makeGrid(height int, width int) [][]int {
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int, width)
	}
	return grid
}

// This function finds the adjacent nodes to the current one
func getAdjNodes(currY int, currX int, maxY int) [][]int {
	var adj [][]int
	moveY := []int{0, -1, 1}
	lookInCol := currX + 1
	var yLocs [3]int
	for i := 0; i < 3; i++ {
		newY := currY + moveY[i]
		if newY < 0 {
			newY = maxY
		} else if newY == maxY+1 {
			newY = 0
		}
		yLocs[i] = newY
	}
	for _, item := range yLocs {
		newLocation := []int{
			item, lookInCol,
		}
		if !alreadyIn(adj, newLocation) {
			adj = append(adj, newLocation)
		}
	}
	return adj
}

// Given a path, we return the total weight of the path
func getTotalWeight(grid [][]int, path [][]int) int {
	totalWeight := 0
	for _, location := range path {
		totalWeight += grid[location[0]][location[1]]
	}
	return totalWeight
}

// This function gives the index of the smallest value of the array
func getMinIdx(arr []int) int {
	min := arr[0]
	idx := 0

	for i, val := range arr {
		if val < min {
			min = val
			idx = i
		}
	}
	return idx
}

// This returns the proper formatting for displaying the row path
func getRows(path [][]int) string {
	rowString := ""
	for i, item := range path {
		rowString = rowString + strconv.Itoa(item[0]+1)
		if i < len(path)-1 {
			rowString = rowString + " "
		}
	}
	return rowString
}

// This is the meat of the program
func newGetPath(grid [][]int, currY int, currX int, path [][]int) [][]int {
	// We ensure we're not at the end of our path
	dontGoPast := len(grid[0]) - 1
	path = append(path, []int{currY, currX})
	if currX == dontGoPast {
		return path
	}
	// Otherwise, we go down all the possible paths
	adjNodes := getAdjNodes(currY, currX, len(grid)-1)
	var comparePaths [][][]int
	for _, node := range adjNodes {
		comparePaths = append(comparePaths, newGetPath(grid, node[0], node[1], path))
	}

	// To keep the function from returning every possible path ever,
	// we nix the heavier weight paths
	var weights []int
	for _, paths := range comparePaths {
		weights = append(weights, getTotalWeight(grid, paths))
	}
	return comparePaths[getMinIdx(weights)]
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
		currentLine := strings.Fields(scanner.Text())
		if len(currentLine) == 2 {
			gridDims := convertToIntArr(currentLine)
			grid := makeGrid(gridDims[0], gridDims[1])

			scanner.Scan()
			heightIndex, widthIndex := 0, 0
			// While the grid contains a 0
			for ok := true; ok; ok = contains(grid, 0) {
				currentLine := convertToIntArr(strings.Fields(scanner.Text()))
				// As not every new line is a new row,
				// we have to find the rows ourselves
				for _, item := range currentLine {
					grid[heightIndex][widthIndex] = item
					widthIndex++
					if widthIndex == gridDims[1] {
						heightIndex++
						widthIndex = 0
					}
				}
				// This ensures we don't skip over a possible new grid
				if heightIndex != gridDims[0] {
					scanner.Scan()
				}
			}

			// This part finds all the minimum paths from each beginning node
			// and finds the smallest of them all
			var finalPaths [][][]int
			for i := 0; i < gridDims[0]; i++ {
				var emptyPath [][]int
				finalPaths = append(finalPaths, newGetPath(grid, i, 0, emptyPath))
			}
			var weights []int
			for _, paths := range finalPaths {
				weights = append(weights, getTotalWeight(grid, paths))
			}
			path := finalPaths[getMinIdx(weights)]
			// This formats the output
			pathRowArr := getRows(path)
			pathWeight := getTotalWeight(grid, path)
			outputFile.WriteString(pathRowArr + "\n")
			outputFile.WriteString(strconv.Itoa(pathWeight) + "\n")
			outputFile.Sync()
		}

	}

}
