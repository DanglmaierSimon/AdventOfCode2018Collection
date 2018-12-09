package main

import (
	"container/list"
	"fmt"
)

type testcase struct {
	players int
	marbles int
	score   int64
}

func main() {

	//Test 1: Result: 32
	//const players = 9
	//const marbles = 25

	//Test 2: Result: 8317
	//const players = 10
	//const marbles = 1618

	//Test 3: Result: 146373
	//const players = 13
	//const marbles = 7999

	//Test 4: Result: 2764
	//const players = 17
	//const marbles = 1104

	//Test 5: Result: 54718
	//const players = 21
	//const marbles = 6111

	//Test 6: Result: 37305
	//const players = 30
	//const marbles = 5807

	//The real deal: Result: 384205
	const players = 476
	const marbles = 71431 * 100

	var board = list.New()

	var score [players]int

	var max = score[0]
	var maxIdx int

	board.PushBack(0)
	board.PushBack(2)
	board.PushBack(1)
	board.PushBack(3)

	var currMarbleIdx = board.Back()

	//fmt.Println(board)

	//var currMarble = 3
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
	fmt.Printf("Highest Score: %d; Player ID: %d\n", max, maxIdx)
}
