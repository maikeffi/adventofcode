/* inspired from lizthegrey
 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day01.input", "Relative file path to use as input.")

func main() {
	bytes, err := ioutil.ReadFile(*inputFile)

	if err != nil {
		return
	}

	contents := string(bytes)
	split := strings.Split(contents, "\n")
	//fmt.Println(split)
	split = split[:len(split)-1]
	//fmt.Println(split)
	seen := make([]int, len(split))
	contains := make(map[int]int)
	for i, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Failed to parse %s\n", s)
			break
		}
		if n <= 0 {
			fmt.Printf("Optimization invariant broken: %d <= 0 \n", n)
			break
		}
		seen[i] = n
		contains[n] = i
	}

	//fmt.Println(contains)

partA:

	for i, n := range seen {
		if pos, ok := contains[2020-n]; ok && pos != i {
			fmt.Println(n * (2020 - n))
			break partA
		}
	}

}
