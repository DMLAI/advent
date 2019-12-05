package main

import (
	"bufio"
	"errors"
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
func RunInstructions(intcode []int) (int, error) {
	for i := 0; i < len(intcode); i += 4 {
		opcode := intcode[i]
		switch opcode {
		case 1:
			intcode[intcode[i+3]] = intcode[intcode[i+1]] + intcode[intcode[i+2]]
		case 2:
			intcode[intcode[i+3]] = intcode[intcode[i+1]] * intcode[intcode[i+2]]
		case 99:
			return intcode[0], nil
		default:
			return 0, errors.New("Unknown opcode")
		}
	}
	return 0, errors.New("Out of instructions")
}

// FindInputs finds two inputs that produce desired output
func FindInputs(want int, intcode []int) (int, error) {
	memory := make([]int, len(intcode))
	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			copy(memory, intcode)
			memory[1] = noun
			memory[2] = verb
			output, err := RunInstructions(memory)
			if err != nil {
				fmt.Println("Could not run instructions")
				continue
			}
			if output == want {
				return 100*noun + verb, nil
			}
		}
	}
	return 0, errors.New("Not found")
}

func main() {
	// Part 1
	instructions := readInstructions("input.txt")
	instructions[1] = 12
	instructions[2] = 2
	partOne, err := RunInstructions(instructions)
	if err != nil {
		fmt.Println("Could not run instructions")
	}
	fmt.Println("Part one:", partOne)

	// Part 2
	instructions = readInstructions("input.txt")
	partTwo, err := FindInputs(19690720, instructions)
	if err != nil {
		fmt.Println("Could not find appropriate noun / verb")
	}
	fmt.Println("Part two:", partTwo)
}
