package main

import (
	"bufio"
	"fmt"
	"math"
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

type info struct {
	minx int
	miny int

	maxx int
	maxy int

	area int64
}

type previt struct {
	i info
	l []light
}

func main() {
	var input []string
	var lmap []light
	var prev info
	var curr info

	var history []previt

	//file, err := os.Open("input_test_1.txt")
	file, err := os.Open("input.txt")

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
	//lmapPrev = lmap
	lmap, prev = doSim(lmap)

	//iteration := 0

	for {
		lmap, curr = doSim(lmap)

		history = append(history, previt{curr, lmap})
		if prev.area < curr.area {
			//printMap(lmap, curr)
			break
		}

		prev = curr
		//lmapPrev = lmap
		/*
			fmt.Println(iteration)
			printMap(lmap, curr)
			time.Sleep(1000000000)
			fmt.Println()
			fmt.Println()

			iteration++
		*/
		/*printMap(lmap, prev)

		fmt.Println()
		fmt.Println()*/

	}

	for i := 1; i <= 20; i++ {
		printMap(history[len(history)-i].l, history[len(history)-i].i)

		fmt.Println()
		fmt.Println()
	}

	//fmt.Println(len(history))

}

func doSim(lmap []light) ([]light, info) {
	var res info

	res.minx = lmap[0].px
	res.miny = lmap[0].py

	res.maxx = lmap[0].px
	res.maxy = lmap[0].py

	for i := range lmap {
		lmap[i].px += lmap[i].vx
		lmap[i].py += lmap[i].vy

		if lmap[i].px < res.minx {
			res.minx = lmap[i].px
		} else if lmap[i].px > res.maxx {
			res.maxx = lmap[i].px
		}

		if lmap[i].py < res.miny {
			res.miny = lmap[i].py
		} else if lmap[i].py > res.maxy {
			res.maxy = lmap[i].py
		}
	}

	height := math.Abs(float64(res.maxx - res.minx))
	width := math.Abs(float64(res.maxy - res.miny))
	res.area = int64(height) * int64(width)

	return lmap, res
}

func printMap(lmap []light, i info) {
	diffx := int((math.Abs(float64(i.maxx - i.minx))))
	diffy := int((math.Abs(float64(i.maxy - i.miny))))

	for y := 0; y <= diffy; y++ {
		for x := 0; x <= diffx; x++ {

			printDot := true
			for _, z := range lmap {
				if z.px == i.minx+x && z.py == i.miny+y {
					printDot = false
					break
				}
			}

			if printDot {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}

		fmt.Println()
	}

}
