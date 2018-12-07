package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	s, err := read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Iterate over the slice and record if a letter appears twice or thrice
	c := count(s)
	// Multiply number of strings with twice letters by number of strings with thrice letters
	fmt.Println(c.twice * c.thrice)
}

// Read the file into a slice of strings
func read(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var out []string
	for scanner.Scan() {
		linestr := scanner.Text()
		out = append(out, linestr)
	}
	return out, nil
}

type counts struct {
	twice  int
	thrice int
}

func count(s []string) (counts) {
	var c counts

	for _, v := range s {
		runes := []rune(v)
		m := map[rune]int{}
		for _, r := range runes {
			m[r] += 1
		}
		for _, co := range m {
			if co == 2 {
				c.twice += 1
				break
			}
		}
		for _, co := range m {
			if co == 3 {
				c.thrice += 1
				break
			}
		}
	}
	return c
}
