package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	inputFile, err := os.Open("spam.in")
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)

	outputFile, err := os.Create("spam.out") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer outputFile.Close()

	/*//re := regexp.MustCompile(`(?!\.)(?!.*\.$)(?!.*?\.\.)[\w.-]+@(?!\.)(?!.*\.$)(?!.*?\.\.)[\w_.]+`)
	re := regexp.MustCompile(`@`)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 && strings.Contains(scanner.Text(), "@") {
			matches := re.FindAllStringIndex(scanner.Text(), -1)
			for _, atLocation := range matches {
				line := string(scanner.Text())
				myIndex := int(atLocation[0])
				beginningPortion := ""
				for i := myIndex; i >= 0; i -- {

				}
			}
		}
	}*/

	outputFile.WriteString("Start of Spam Program\n")

	re := regexp.MustCompile(`((\w|\d|\-|\_)+(.?))+@((\w|\d|\-|\_)+(.?))+`)
	for scanner.Scan() {
		matches := re.FindAllString(scanner.Text(), -1)
		for _, match := range matches {
			if len(match) > 0 {
				outputFile.WriteString(match + "\n")
			}
		}
	}

	outputFile.WriteString("End of Spam Program")

	outputFile.Sync()

}
