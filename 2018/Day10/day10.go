package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type light struct {
	px int
	py int
	vx int
	vy int
}

type arr struct {
	minx int
	miny int
	maxx int
	maxy int
}

func main() {
	var input []string
	var lmap []light
	var limits arr

	file, err := os.Open("input_test_1.txt")
	//file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error: Could not open file!")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	for _, l := range input {
		fmt.Println(l)

		t := strings.SplitAfter(l, ">")

		t[0] = strings.Trim(t[0], "position= velocity= < >")
		t[1] = strings.Trim(t[1], "position= velocity= < >")

		fmt.Println(t[0] + "; " + t[1])

		pos := strings.Split(t[0], ",")

		pos[0] = strings.TrimSpace(pos[0])
		pos[1] = strings.TrimSpace(pos[1])

		x, err := strconv.Atoi(pos[0])

		if err != nil {
			fmt.Println("Error converting x coordinate")
			return
		}

		y, err := strconv.Atoi(pos[1])

		if err != nil {
			fmt.Println("Error converting y coordinate")
			return
		}

		vel := strings.Split(t[1], ",")

		vel[0] = strings.TrimSpace(vel[0])
		vel[1] = strings.TrimSpace(vel[1])

		vx, err := strconv.Atoi(vel[0])

		if err != nil {
			fmt.Println("Error converting x coordinate")
			return
		}

		vy, err := strconv.Atoi(vel[1])

		if err != nil {
			fmt.Println("Error converting y coordinate")
			return
		}

		lmap = append(lmap, light{x, y, vx, vy})

	}

	for !checkForMessage(&lmap) {
		doSim(&lmap)
	}

}

func checkForMessage(lmap *[]light) bool {

	for _, x := range *lmap {
		for _, y := range *lmap {
			if !(x.px-1 == y.px || x.px+1 == y.px || x.py-1 == y.py || x.py+1 == y.py) {
				return false
			}
		}
	}

	return true
}

func doSim(lmap *[]light, a arr) (res arr) {

}
