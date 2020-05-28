import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	for scanner.Scan () {
		toSum = int(scanner.Text())
		fmt.Println(toSum)
	}

}