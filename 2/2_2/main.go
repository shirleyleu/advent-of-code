package main

import (
	"bufio"
	"errors"
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
	s1, s2, err := iterate(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s1, s2)
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

func iterate(s []string) (string, string, error) {
	for i, _ := range s {
		for j := i + 1; j < len(s)-1; j++ {
			diff := compare(s[i], s[j])
			if diff == 1 {
				return s[i], s[j], nil
			}
		}
	}
	return "", "", errors.New("no string with only 1 character difference")
}

func compare(s1, s2 string) int {
	var c int
	for i, l := range s2 {
		if rune(s1[i]) != rune(l) {
			c += 1
		}
	}
	return c
}
