package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
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

func parse(input *os.File) (int, error) {
	freq := 0
	var m = make(map[int]int)
	// Initialize for freq of 0
	m[freq] += 1
	for {
		input.Seek(0,0)
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

			// Check the count for a freq to break
			if m[freq] > 1 {
				return freq, nil
			}
		}
	}
}
