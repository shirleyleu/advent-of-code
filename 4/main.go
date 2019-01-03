package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	// Read input file into a slice of strings
	s, err := read("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Sort slice in alphabetical order
	sort.Strings(s)

	// Associate each guard id with minutes asleep
	m := mapSleepTimes(s)

	// Count for each guard the total number of minutes asleep
	n := totalAsleep(m)

	fmt.Println(mostAsleep(n))

	// For which minute was he asleep the most often?
	fmt.Println(minuteMostAsleep(timeAsleep(m)))

	j := timeAsleep(m)
	fmt.Println(mostSameMinute(j))
}

func mostSameMinute(m map[ID]map[int]int) (ID, int) {
	most := 0
	var minute int
	var g ID
	for guard, tally := range m {
		for min, count := range tally {
			if count > most {
				most = count
				minute = min
				g = guard
			}
		}
	}
	return g, minute
}

// Read input file into slice of strings
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

type ID int
type sleep struct {
	start int
	end   int
}
type logSleep map[ID][]sleep

func mapSleepTimes(s []string) logSleep {
	reID := regexp.MustCompile(`Guard #(\d+)`)
	reMin := regexp.MustCompile(`\[*00:(\d+)] (\w)`)
	l := logSleep{}
	var guard ID
	var fallMinute int
	for _, line := range s {
		matchID := reID.FindStringSubmatch(line)
		switch {
		case matchID != nil:
			num, err := strconv.Atoi(matchID[1])
			if err != nil {
				log.Fatal(err)
			}
			guard = ID(num)
		case matchID == nil:
			sleepGroups := reMin.FindStringSubmatch(line)
			minute, err := strconv.Atoi(sleepGroups[1])
			if err != nil {
				log.Fatal(err)
			}
			letter := sleepGroups[2]
			switch {
			case letter == "f":
				fallMinute = minute
			case letter == "w":
				l[guard] = append(l[guard], sleep{fallMinute, minute})
			}
		}
	}
	return l
}

func totalAsleep(l logSleep) map[ID]int {
	m := make(map[ID]int)
	for guard, times := range l {
		for _, time := range times {
			n := time.end - time.start
			m[guard] += n
		}
	}
	return m
}

func mostAsleep(m map[ID]int) ID {
	most := 0
	var g ID
	for guard, total := range m {
		if total > most {
			most = total
			g = guard
		}
	}
	return g
}

func timeAsleep(l logSleep) map[ID]map[int]int {
	m := make(map[ID]map[int]int)
	for guard, times := range l {
		for _, time := range times {
			for min := time.start; min < time.end; min++ {
				if _, ok := m[guard]; !ok {
					m[guard] = make(map[int]int)
				}
				m[guard][min] += 1
			}
		}
	}
	return m
}

func minuteMostAsleep(l map[ID]map[int]int) map[ID]int {
	n := make(map[ID]int)
	for guard, tally := range l {
		highest := 0
		for minute, count := range tally {
			if count > highest {
				highest = count
				n[guard] = minute
			}
		}
	}
	return n
}
