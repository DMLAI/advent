package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
func RunInstructions(intcode []int) int {
	for i := 0; i < len(intcode); i += 4 {
		opcode := intcode[i]
		switch opcode {
		case 1:
			intcode[intcode[i+3]] = intcode[intcode[i+1]] + intcode[intcode[i+2]]
		case 2:
			intcode[intcode[i+3]] = intcode[intcode[i+1]] * intcode[intcode[i+2]]
		case 99:
			return intcode[0]
		default:
			fmt.Println(intcode[i])
			panic("Unknown opcode")
		}
	}
	return intcode[0]
}

func main() {
	instructions := readInstructions("input.txt")
	instructions[1] = 12
	instructions[2] = 2
	fmt.Println(RunInstructions(instructions))
}
