package main

import (
	"fmt"
)

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
	const marbles = 71431 * 10

	var board []int

	var score [players]int

	var max = score[0]
	var maxIdx int

	board = append(board, 0, 2, 1, 3)

	fmt.Println(board)

	var currMarble = 3
	var currPlayer = 3

	for m := 4; m <= marbles; m++ {
		var currMarbleIdx int

		currPlayer = (m % players)

		//Find the index of the current marble
		for i := range board {
			if board[i] == currMarble {
				currMarbleIdx = i
				currMarble = m
				break
			}
		}

		if (m % 23) != 0 {

			idx := (currMarbleIdx + 2) % len(board)

			if idx == 0 {
				board = append(board, m)
			} else {

				board = append(board, 0)

				copy(board[idx+1:], board[idx:])

				board[idx] = m
			}
		} else {

			//actual length
			act := 7 % len(board)

			idx := (currMarbleIdx - act + len(board)) % len(board)

			score[currPlayer] += board[idx] + m

			if score[currPlayer] > max {
				maxIdx = currPlayer
				max = score[currPlayer]
			}

			copy(board[idx:], board[idx+1:])
			board[len(board)-1] = 0
			board = board[:len(board)-1]

			currMarble = board[idx]

		}
	}
	fmt.Printf("Highest Score: %d; Player ID: %d\n", max, maxIdx)
}
