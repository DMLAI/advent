package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// CalculateFuel computes fuel required per module
func CalculateFuel(mass int64) int64 {
	return mass/3 - 2
}

//SumFuelRequirement reads in module weights
func SumFuelRequirement(filename string) int64 {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed opening file: ", err)
	}

	scanner := bufio.NewScanner(file)
	var total int64
	for scanner.Scan() {
		weight, err := strconv.ParseInt(scanner.Text(), 0, 64)
		if err != nil {
			fmt.Println("Error parsing weight", err)
		}
		total += CalculateFuel(weight)
	}

	return total
}

func main() {
	fmt.Println(SumFuelRequirement("input.txt"))
}
