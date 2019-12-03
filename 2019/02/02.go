package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func applyOpcode(opcode int, instructions []int) int {
	switch opcode {
	case 1:
		fmt.Println("add")
	case 2:
		fmt.Println("multiply")
	case 99:
		return instructions[0]
	default:
		panic("Unknown opcode")
	}
	return 42
}

func readInstructions(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed opening file: ", err)
	}

	dropLF := func(data []byte) []byte {
		if len(data) > 0 && data[len(data)-1] == '\n' {
			return data[0 : len(data)-1]
		}
		return data
	}
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		return 0, dropLF(data), bufio.ErrFinalToken
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(onComma)
	var instructions []int
	for scanner.Scan() {
		code, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			fmt.Println("Error parsing code", err)
		}
		instructions = append(instructions, int(code))
	}

	file.Close()

	return instructions
}

// RunInstructions executes Intcode
func RunInstructions(instructions []int) int {
	return 42
}

func main() {
	instructions := readInstructions("input.txt")
	fmt.Println(RunInstructions(instructions))
}
