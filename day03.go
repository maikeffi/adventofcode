package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var inputFile = flag.String("inputFile", "inputs/day03.input", "Relative file path to use as input.")

func main() {
	//var tile [][]rune
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
	col := 0
	tree := rune("#"[0])
	treeCount := 0

	for i := 0; i < len(txtlines)-1; i++ {
		r := []rune(txtlines[i+1])
		col = col + 3
		fmt.Printf("Row no %d tile at position %d is %d\n", i, col, r[col])
		if r[col] == tree {
			treeCount++
		}
	}

	fmt.Println(treeCount)
}
