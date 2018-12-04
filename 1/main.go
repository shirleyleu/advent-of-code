package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	freq, err := parse(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(freq)
}

func parse(input io.Reader) (int, error) {
	freq := 0
	var m = make(map[int]int)
	// Initialize for freq of 0
	m[freq] += 1
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lineStr := scanner.Text()
		num, err := strconv.Atoi(lineStr)
		if err != nil {
			return 0, errors.Wrap(err, "parse input file error")
		}
		// Change the freq
		freq += num

		// Increment the number of times that freq has been hit
		m[freq] += 1

		// check how many times this freq has been hit

	}
	return freq, nil
}
