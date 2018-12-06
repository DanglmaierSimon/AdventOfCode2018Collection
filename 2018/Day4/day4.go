package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type action struct {
	date time.Time
	act  string
}

type guardData struct {
	sleepTime  uint
	bestMinute uint
}

func main() {
	var input []string
	var sortedInput []action
	//var guards map[int]guardData

	//file, err := os.Open("input_test_1.txt")
	file, err := os.Open("input_test_1.txt")

	if err != nil {
		fmt.Println("Error: Could not open file!")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	for _, val := range input {
		fmt.Println(val)

		line := strings.SplitAfter(val, "]")

		timeStr := strings.Trim(line[0], "[]")

		layout := "2006-01-02 15:04"

		var timeObj time.Time
		timeObj, err := time.Parse(layout, timeStr)

		if err != nil {
			fmt.Println("Error converting time")
			return
		}

		sortedInput = append(sortedInput, action{timeObj, line[1]})
	}

	sort.Slice(sortedInput[:], func(i, j int) bool {
		return sortedInput[i].date.Before(sortedInput[j].date)
	})
	i := 0

	for {
		atEnd, res := getGuardActivity(sortedInput, 0)

		if atEnd {
			break
		}

		i += len(res)

	}
}

func extractGuardID(l string) int {
	idx1 := strings.Index(l, "#")

	if idx1 == -1 {
		return -1
	}

	idx2 := strings.Index(l, "begins")

	if idx2 == -1 {
		return -1
	}

	numStr := l[idx1+1 : idx2-1]

	id, err := strconv.Atoi(numStr)

	if err != nil {
		fmt.Println("Error extracting ID")
		return -1
	}

	return id
}

func getGuardActivity(in []action, idx int) (bool, []action) {
	var retVal []action

	if idx >= len(in) {
		return true, nil
	}

	retVal = append(retVal, in[idx])

	idx++

	for i := idx; i < len(in); i++ {

		if strings.Contains(in[i].act, "begins") {
			return false, retVal
		}

		retVal = append(retVal, in[i])
	}

	return true, retVal
}
