package main

import (
	"encoding/csv"
	_ "encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("problem.cvs")
	if err != nil {
		fmt.Println("Problem opening csv file")
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Could not close file")
		}
	}(file)
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	var totalQuestion float32
	var correctAns float32

	for _, row := range records {
		fmt.Println("what is the sum of ", row[0], "?")
		var ans int

		for {
			_, err := fmt.Scanln(&ans)

			if err != nil {
				fmt.Println("Could not read input try again")
				_, _ = fmt.Scanln()
				continue
			}
			break
		}

		number, _ := strconv.Atoi(row[1])

		if ans == number {
			correctAns++
		}
		totalQuestion++
	}
	fmt.Println("Total question", totalQuestion, "Correct ans", correctAns)

	fmt.Println("Your accuracy is", (correctAns/totalQuestion)*100)

}
