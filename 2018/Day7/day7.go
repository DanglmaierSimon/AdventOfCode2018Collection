package main

/*
import (
	"bufio"
	"fmt"
	"os"
)

type tNode struct {
	name rune

	remPred uint

	successors []*tNode
}

type graphHead *tNode

func main() {

	var input []string

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

}

func insertNode(head graphHead, currNode rune, succNode rune) bool {
	deleteNode := false

	startNode := findNode(head, currNode)
	endNode := findNode(head, succNode)

	if startNode == nil && endNode == nil {

	}
}

func addSuccessor(node *tNode, name rune) {
	newNode := new(tNode)

	newNode.name = name
	newNode.remPred = 1

	node.successors = append(node.successors, newNode)
}

func findNode(node *tNode, name rune) *tNode {
	if node == nil {
		return nil
	} else if node.name == name {
		return node
	} else {
		for _, e := range node.successors {
			if findNode(e, name) != nil {
				return e
			}
		}
	}

	return nil
}
*/
