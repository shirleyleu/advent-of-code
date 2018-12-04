package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	ints, err := read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(search(ints))
}

// Read the file into a slice of ints
func read(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var out []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		out = append(out, i)
	}
	return out, nil
}

// Iterate over the slice and record the frequency into a map
func search(s []int) int {
	i := 0
	f := 0
	l := len(s) - 1
	seen := map[int]bool{0: true}
	for {
		f += s[i]
		if seen[f] {
			return f
		}
		seen[f] = true
		if i == l {
			i = 0
		} else {
			i++
		}
	}
}
