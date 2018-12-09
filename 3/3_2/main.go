package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
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
	sq := listSquares(s)

	id, err := iterateAndCount(sq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
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

type squares struct {
	id           int
	sliceSquares []square
}

// Parse through slice to create to create a list of ids and their squares
func listSquares(s []string) []squares {
	var result []squares

	for _, v := range s {
		id, c := fabricCoord(v)
		result = append(result, squares{id: id, sliceSquares: c})
	}
	return result
}

func fabricCoord(s string) (int, []square) {
	re := regexp.MustCompile(`(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	slice := re.FindStringSubmatch(s)
	var coords []square
	id, err := strconv.Atoi(slice[1])
	if err != nil {
		log.Fatal(err)
	}
	x_i, err := strconv.Atoi(slice[2])
	if err != nil {
		log.Fatal(err)
	}
	y_i, err := strconv.Atoi(slice[3])
	if err != nil {
		log.Fatal(err)
	}
	x_size, err := strconv.Atoi(slice[4])
	if err != nil {
		log.Fatal(err)
	}
	y_size, err := strconv.Atoi(slice[5])
	if err != nil {
		log.Fatal(err)
	}
	for i := x_i; i < x_i+x_size; i++ {
		for j := y_i; j < y_i+y_size; j++ {
			coords = append(coords, square{x: i, y: j})
		}
	}
	return id, coords
}

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

// Make a map representing squares and log the used squares in the map
// Iterate over each id and find which one has square of only count 1 in map
func iterateAndCount(sq []squares) (int, error) {
	m := make(map[square]int)
	for _, e := range sq {
		for _, coord := range e.sliceSquares {
			m[coord] += 1
		}
	}
	for _, e := range sq {
		found := true
		for _, coord := range e.sliceSquares {
			if m[coord] != 1 {
				found = false
			}
		}
		if found == true {
			return e.id, nil
		}
	}
	return 0, errors.New("No id with no overlapping squares")
}
