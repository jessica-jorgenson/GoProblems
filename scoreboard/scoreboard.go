package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringNum := scanner.Text()
	number, err := strconv.Atoi(stringNum)
	if err != nil {
		fmt.Println("Error in read_integer:", err)
	}
	return number
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
		return
	}
	defer outputFile.Close()

	cases := readInt(scanner)
	scanner.Scan()
	scanner.Scan()
	for i := 1; i <= cases; i++ {
		board := make(map[int]map[string][]int)

		for len(scanner.Text()) > 0 {
			data := strings.Fields(scanner.Text())
			user, err := strconv.Atoi(data[0])
			if err != nil {
				fmt.Println(err)
			}

			if data[len(data)-1] == "I" || data[len(data)-1] == "C" {
				prob, err := strconv.Atoi(data[1])
				if err != nil {
					fmt.Println(err)
				}
				times, err := strconv.Atoi(data[2])
				if err != nil {
					fmt.Println(err)
				}
				if test, inBoard := board[user]; !inBoard {
					fmt.Println(test)
					board[user] = make(map[string][]int)
					board[user]["problems"] = []int{prob}
					board[user]["times"] = []int{times}
				} else {
					board[user]["times"] = append(board[user]["times"], times)
					inArray := false
					for _, problem := range board[user]["problems"] {
						if problem == prob {
							inArray = true
						}
					}
					if !inArray {
						board[user]["problems"] = append(board[user]["problems"], prob)
					}

				}
			}
			fmt.Println(board)
			scanner.Scan()
		}
	}

}
