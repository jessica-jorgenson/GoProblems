package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"reflect"
)

func convertToIntArr(stringArr []string) []int {
	var returnArr []int
	for _, item := range stringArr {
		integer, _ := strconv.Atoi(item)
		returnArr = append(returnArr, integer)
	}
	return returnArr
}

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

func alreadyIn(arr [][]int, point []int) bool{
	for _, item := range arr {
		if reflect.DeepEqual(item, point){
			return true
		}
	}
	return false
}

func makeGrid(height int, width int) [][]int{
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int, width)
	}
	return grid
}

func getAdjNodes(currY int, currX int, maxY int) [][]int {
	var adj [][]int
	moveY := []int{0, -1, 1,}
	lookInCol := currX + 1
	var yLocs [3]int 
	for i := 0; i < 3; i++ {
		newY := currY + moveY[i]
		if newY < 0 {
			newY = maxY
		} else if newY == maxY + 1 {
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

func getTotalWeight(grid [][]int, path [][]int) int {
	totalWeight := 0
	for _, location := range path {
		totalWeight += grid[location[0]][location[1]]
	}
	return totalWeight
}

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

func newGetPath(grid [][]int, currY int, currX int, path [][]int) [][]int {
	dontGoPast := len(grid[0]) - 1
	path = append(path, []int{currY, currX})
	if currX == dontGoPast {
		return path
	} else {
		adjNodes := getAdjNodes(currY, currX, len(grid) - 1)
		var comparePaths [][][]int
		for _, node := range adjNodes {
			comparePaths = append(comparePaths, newGetPath(grid, node[0], node[1], path))
		}

		var weights []int
		for _, paths := range comparePaths {
			weights = append(weights, getTotalWeight(grid, paths))
		}
		return comparePaths[getMinIdx(weights)]
	}
	return path
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

	for scanner.Scan() {
		currentLine := strings.Fields(scanner.Text())
		if len(currentLine) == 2 {
			gridDims := convertToIntArr(currentLine)
			grid := makeGrid(gridDims[0], gridDims[1])
			
			scanner.Scan()
			heightIndex, widthIndex := 0, 0
			for ok := true; ok; ok = contains(grid, 0) {
				currentLine := convertToIntArr(strings.Fields(scanner.Text()))
				for _, item := range currentLine {
					grid[heightIndex][widthIndex] = item
					widthIndex++
					if widthIndex == gridDims[1] {
						heightIndex++
						widthIndex = 0
					}
				}
				if heightIndex != gridDims[0] {
					scanner.Scan()
				}
			}
		
			path := [][]int {
			}
			path = newGetPath(grid, 0, 0, path);
			fmt.Println(path);
		}

	}


}
