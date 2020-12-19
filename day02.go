package main

// used inspiration from @lizthegray

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative file path to use as input.")

func main() {

	var countA, countB int

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var txtlines []string
	//validCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, str := range txtlines {
		splitString := strings.Split(str, " ")
		char := rune(splitString[1][0])
		charStr := string(splitString[1][0])
		password := splitString[2]
		passSize := strings.Split(splitString[0], "-")
		min, err := strconv.Atoi(passSize[0])
		max, err := strconv.Atoi(passSize[1])
		if err != nil {
			fmt.Printf("Failed to parse %s\n", splitString[0])
			break
		}
		count := strings.Count(password, charStr)
		if count >= min && count <= max {
			countA++
		}
		if (rune(password[min-1]) == char) != (rune(password[max-1]) == char) {
			countB++
		}

	}

	fmt.Println(countA)
	fmt.Println(countB)

}
