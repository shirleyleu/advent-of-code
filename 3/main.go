package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	s, err := read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	p := parse(s)
	fmt.Println(count(p))
}

// Read file into slice of strings
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

// Square coordinates are defined by the lower left point and represent a square inch
type square struct {
	x int
	y int
}

// Parse through slice to create to create a list of square inches
func parse(s []string) [][]square {
	var result [][]square
	for _, v := range s {
		result = append(result, fabricCoord(v))
	}
	return result
}

func fabricCoord(s string) []square {
	re := regexp.MustCompile(`@ (\d+),(\d+): (\d+)x(\d+)`)
	slice := re.FindStringSubmatch(s)
	var coords []square
	x_i, err := strconv.Atoi(slice[1])
	if err != nil{
		log.Fatal(err)
	}
	y_i, err := strconv.Atoi(slice[2])
	if err != nil{
		log.Fatal(err)
	}
	x_size, err := strconv.Atoi(slice[3])
	if err != nil{
		log.Fatal(err)
	}
	y_size, err := strconv.Atoi(slice[4])
	if err != nil{
		log.Fatal(err)
	}
	for i := x_i; i < x_i+x_size; i++ {
		for j := y_i; j < y_i+y_size; j++ {
			coords = append(coords, square{x: i, y: j})
		}
	}
	return coords
}

// Make a map representing squares and log the used squares in the map
// Iterate over map to get the number of square inches of fabric with 2 or more claims
func count(c [][]square) int {
	var count int
	m := make(map[square]int)
	for _, e := range c {
		for _, coord := range e {
			m[coord] += 1
		}
	}
	for _, v := range m {
		if v > 1 {
			count += 1
		}
	}
	return count
}
