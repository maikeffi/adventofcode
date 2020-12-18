package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("inputs/day01.input")
	if err != nil {
		fmt.Println("Error opening file ", err)
		os.Exit(1)
	}
	var nums []int
	var num int
	var found bool

	for {
		_, err := fmt.Fscanf(f, "%d\n", &num)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		nums = append(nums, num)
	}

	fmt.Println(nums)

	for _, i := range nums {
		if !found {
			for _, j := range nums {
				if !found {
					for _, k := range nums {
						if i+j+k == 2020 {
							fmt.Printf("%d , %d , %d, %d \n", i, j, k, i*j*k)
							found = true
						}

					}
				}

			}
		}

	}

}
