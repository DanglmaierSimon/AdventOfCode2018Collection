package main

import (
	"container/list"
	"fmt"
	"time"
)

type testcase struct {
	players int
	marbles int
	exp     int64
}

func main() {

	tests := []testcase{
		testcase{9, 25, 32},
		testcase{10, 1618, 8317},
		testcase{13, 7999, 146373},
		testcase{17, 1104, 2764},
		testcase{21, 6111, 54718},
		testcase{30, 5807, 37305},
		testcase{476, 71431, 384205},
		testcase{476, 7143100, 3066307353},
	}

	for i, t := range tests {
		start := time.Now()

		score, id := doCalc(t.players, t.marbles)

		dur := time.Since(start)

		fmt.Printf("Testcase [%d]: ", i)

		if score == t.exp {
			fmt.Printf("Test passed! Score: %d; Player ID: %d\n", score, id+1)
		} else {
			fmt.Printf("Test not passed! Expected Score: %d; Result: %d\n", t.exp, score)
		}

		fmt.Printf("Duration: %s\n=================\n", dur)
	}
}

func doCalc(players int, marbles int) (result int64, id int) {
	var board = list.New()

	var score = make([]int, players)

	var max = score[0]
	var maxIdx int

	board.PushBack(0)
	board.PushBack(2)
	board.PushBack(1)
	board.PushBack(3)

	var currMarbleIdx = board.Back()

	var currPlayer = 3

	for m := 4; m <= marbles; m++ {

		var idx *list.Element
		currPlayer = (m % players)

		idx = currMarbleIdx

		if (m % 23) != 0 {
			//Advance pointer by 2 places, loop to front if needed

			for i := 0; i < 2; i++ {
				if idx.Next() == nil {
					idx = board.Front()
				} else {
					idx = idx.Next()
				}
			}

			if idx == board.Front() {
				board.PushBack(m)
				currMarbleIdx = board.Back()
			} else {
				board.InsertBefore(m, idx)
				currMarbleIdx = idx.Prev()
			}
		} else {

			for i := 0; i < 7; i++ {
				if idx.Prev() == nil {
					idx = board.Back()
				} else {
					idx = idx.Prev()
				}
			}

			score[currPlayer] += idx.Value.(int) + m

			if score[currPlayer] > max {
				maxIdx = currPlayer
				max = score[currPlayer]
			}

			if idx.Next() == nil {
				currMarbleIdx = board.Front()
			} else {
				currMarbleIdx = idx.Next()
			}

			board.Remove(idx)
		}

	}

	return int64(max), maxIdx
}
