package main

import (
	"fmt"
	"strconv"
)

func main() {
	const test1input = 9
	const test2input = 5
	const test3input = 18
	const test4input = 2018
	const input = 793031

	//inputArr := []int8{7, 9, 3, 0, 3, 1}
	inputArr := []int8{5, 1, 5, 8, 9}

	lim := 793031

	sb := []int{3, 7}

	pos1 := 0
	pos2 := 1

	for len(sb) < lim+10 {
		var temp []int

		/*
			for i := range sb {
				/*if i == pos1 {
					fmt.Printf("(%d)", sb[i])
				} else if i == pos2 {
					fmt.Printf("[%d]", sb[i])
				} else {

				}

				//fmt.Printf(" %d", sb[i])
			}
			fmt.Println()

			//fmt.Println(sb)

			//fmt.Printf("Pos1: %d; Pos2: %d\n", pos1, pos2)

		*/

		sum := sb[pos1] + sb[pos2]

		if sum == 0 {
			sb = append(sb, 0)
		} else {

			for sum > 0 {
				temp = append(temp, sum%10)
				sum = sum / 10
			}

			for i := len(temp) - 1; i >= 0; i-- {
				sb = append(sb, temp[i])
			}
		}

		pos1 = (pos1 + 1 + sb[pos1]) % len(sb)
		pos2 = (pos2 + 1 + sb[pos2]) % len(sb)
	}

	var res string

	for i := lim; i < lim+10; i++ {
		res += strconv.Itoa(sb[i])
	}

	fmt.Println(res)

	sb2 := [100000000]int8{}

	sb2[0] = 3
	sb2[1] = 7

	var lastIdx = 2
	var itr int64

	pos1 = 0
	pos2 = 1

	for !find(&sb2, &inputArr) {
		var temp []int8

		itr++

		if (itr % 10000) == 0 {
			fmt.Printf("Iterations: %d\n", itr)
		}

		sum := sb2[pos1] + sb2[pos2]

		if sum == 0 {
			sb2[lastIdx] = 0
			lastIdx++
		} else {

			for sum > 0 {
				temp = append(temp, sum%10)
				sum = sum / 10
			}

			for i := len(temp) - 1; i >= 0; i-- {
				sb2[lastIdx] = temp[i]
				lastIdx++
			}
		}

		t1 := pos1 + 1 + int(sb2[pos1])
		t2 := pos2 + 1 + int(sb2[pos2])

		pos1 = t1 % len(sb2)
		pos2 = t2 % len(sb2)
	}

	fmt.Println("DONE")
	fmt.Println(sb2[len(sb2)-6:])
}

func find(src *[100000000]int8, sub *[]int8) bool {
	if len(*src) < len(*sub) {
		return false
	}

	end1 := len(*src) - 1
	end2 := len(*sub) - 1

	for i := 0; i < len(*sub); i++ {
		if (*sub)[end2-i] != (*src)[end1-i] {
			return false
		}
	}
	return true

}
