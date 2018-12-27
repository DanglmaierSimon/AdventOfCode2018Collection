package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sample struct {
	bef [4]int
	aft [4]int
	cmd [4]int
}

type opcode struct {
	id  int
	opt []string
}

func main() {
	var input []string

	var reg = [4]int{0, 0, 0, 0}

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

		t := strings.Split(l, " ")

		s := convertToSample(t, t, t)

		reg = executeOpCode(reg, s.cmd)

		/*
			fmt.Println("===AFTER TRIMMING===")

			fmt.Printf("Before\t: %v\n", before)
			fmt.Printf("Command\t: %v\n", cmd)
			fmt.Printf("After\t: %v\n", after)
		*/
		//s := convertToSample(before, after, cmd)
		/*
			fmt.Println("===AFTER CONVERTING===")

			fmt.Printf("Before\t: %v\n", s.bef)
			fmt.Printf("Command\t: %v\n", s.cmd)
			fmt.Printf("After\t: %v\n", s.aft)
		*/
		/*
			_, codes := calcOpCodes(s)

			//fmt.Printf("Result: %d\n", res)

			elem, ok := possibleOpCodes[codes.id]

			if ok {
				possibleOpCodes[codes.id] = append(elem, codes.opt...)
			} else {
				possibleOpCodes[codes.id] = codes.opt
			}
		*/
		/*
			fmt.Println("===CYCLE FINISHED===")
			fmt.Println()
			fmt.Println()
		*/
	}

	fmt.Println(reg)

}

//Returns how many opcodes could fit this sample
func calcOpCodes(s sample) (int, opcode) {
	var codes []string
	var opcode opcode

	retVal := 0

	for op := 0; op < 16; op++ {
		//res, code := executeOpCode(s.bef, s.cmd, op)
		/*
			if res == s.aft {
				retVal++
				codes = append(codes, code)
			}
		*/
	}

	opcode.id = s.cmd[0]
	opcode.opt = codes

	return retVal, opcode
}

func executeOpCode(in [4]int, cmd [4]int) [4]int /*, string*/ {
	//var opCode string
	var out = in

	regA := cmd[1]
	regB := cmd[2]
	regC := cmd[3]

	immA := cmd[1]
	immB := cmd[2]

	switch cmd[0] {
	//Add reg
	case 15:
		out[regC] = in[regA] + in[regB]
		//opCode = "addr"

	//Add imm
	case 4:
		out[regC] = in[regA] + immB
		//opCode = "addi"
	//Mul reg
	case 6:
		out[regC] = in[regA] * in[regB]
		//opCode = "mulr"

	//Mul imm
	case 5:
		out[regC] = in[regA] * immB
		//opCode = "muli"

	//And reg
	case 11:
		out[regC] = in[regA] & in[regB]
		//opCode = "andr"

	//And imm
	case 8:
		out[regC] = in[regA] & immB
		//opCode = "andi"

	//Or reg
	case 12:
		out[regC] = in[regA] | in[regB]
		//opCode = "borr"

	//Or imm
	case 10:
		out[regC] = in[regA] | immB
		//opCode = "bori"

	//Set reg (mov)
	case 2:
		out[regC] = in[regA]
		//opCode = "setr"

	//Set imm
	case 0:
		out[regC] = immA
		//opCode = "seti"

	//Greater than imm/reg
	case 3:
		if immA > in[regB] {
			out[regC] = 1
		} else {
			out[regC] = 0
		}
		//opCode = "gtir"

	//Greather than reg/imm
	case 9:
		if in[regA] > immB {
			out[regC] = 1
		} else {
			out[regC] = 0
		}
		//opCode = "gtri"

	//Greather than reg/reg
	case 7:
		if in[regA] > in[regB] {
			out[regC] = 1
		} else {
			out[regC] = 0
		}
		//opCode = "gtrr"

	//Equal imm/reg
	case 1:
		if immA == in[regB] {
			out[regC] = 1
		} else {
			out[regC] = 0
		}
		//opCode = "eqir"

	//Equal reg/imm
	case 13:
		if in[regA] == immB {
			out[regC] = 1
		} else {
			out[regC] = 0
		}
		//opCode = "eqri"

	//Equal reg/reg
	case 14:
		if in[regA] == in[regB] {
			out[regC] = 1
		} else {
			out[regC] = 0
		}
		//opCode = "eqrr"

	default:
		fmt.Printf("Encountered unexpected Op code: %d", cmd[0])
		return [4]int{}
	}

	return out
}

func convertToSample(before []string, after []string, cmd []string) sample {
	var a [4]int
	var b [4]int
	var c [4]int

	if len(before) > 4 {
		fmt.Println("Invalid length of array BEFORE!")
	}

	if len(after) > 4 {
		fmt.Println("Invalid length of array AFTER!")
	}

	if len(cmd) > 4 {
		fmt.Println("Invalid length of array CMD!")
	}

	for i := 0; i < 4; i++ {
		v, err := strconv.Atoi(after[i])

		if err != nil {
			fmt.Printf("Error converting string %s to a number!\n", after[i])
		}

		a[i] = v
	}

	for i := 0; i < 4; i++ {
		v, err := strconv.Atoi(before[i])

		if err != nil {
			fmt.Printf("Error converting string %s to a number!\n", before[i])
		}

		b[i] = v
	}

	for i := 0; i < 4; i++ {
		v, err := strconv.Atoi(cmd[i])

		if err != nil {
			fmt.Printf("Error converting string %s to a number!\n", cmd[i])
		}

		c[i] = v
	}

	return sample{b, a, c}

}

func removeDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}
