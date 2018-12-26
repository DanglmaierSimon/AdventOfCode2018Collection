package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type pwr struct {
	lvl int
	x   int
	y   int
	s   int
}

func main() {
	//Grid Serial number
	const gsn = 5719

	var grid [300][300]int

	var totalMax, totalX, totalY, size int

	start := time.Now()

	for x := 1; x < 299; x++ {
		for y := 1; y < 299; y++ {
			p := calcPowerLevel(x, y, gsn)
			grid[x][y] = p
		}
	}

	fmt.Printf("Time for filling the array: %s\n", time.Since(start))

	ch := make(chan pwr, 300)

	var wg sync.WaitGroup

	for s := 1; s < 300; s++ {

		go func(si int) {
			var max, maxx, maxy, size int

			wg.Add(1)

			for x := 0; x < (300 - si); x++ {
				for y := 0; y < (300 - si); y++ {
					res := calcTotalPower(x, y, &grid, si)

					if res > max {
						max = res
						maxx = x
						maxy = y
						size = si
					}
				}
			}
			ch <- pwr{max, maxx, maxy, size}
			wg.Done()
		}(s)

	}

	wg.Wait()

	close(ch)

	for v := range ch {
		if v.lvl > totalMax {
			totalMax = v.lvl
			totalX = v.x
			totalY = v.y
			size = v.s
		}
	}
	elapsed := time.Since(start)

	fmt.Println("Max Value: " + strconv.Itoa(totalMax))
	fmt.Println("X-Coord: " + strconv.Itoa(totalX))
	fmt.Println("Y-Coord: " + strconv.Itoa(totalY))
	fmt.Println("Size: " + strconv.Itoa(size))
	fmt.Printf("Total Time: %s\n", elapsed)

}

func calcPowerLevel(x int, y int, gsn int) int {
	rackID := x + 10
	pwrLvl := rackID * y

	pwrLvl += gsn

	pwrLvl *= rackID

	pwrLvl /= 100

	pwrLvl %= 10

	pwrLvl -= 5

	return pwrLvl
}

func calcTotalPower(x int, y int, grid *[300][300]int, s int) int {
	var sum int

	for i := x; i < x+s; i++ {
		for j := y; j < y+s; j++ {
			sum += (*grid)[i][j]
		}
	}

	return sum
}
