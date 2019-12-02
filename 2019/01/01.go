package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculateFuelRec(agg, mass int) int {
	fuelReq := mass/3 - 2
	if fuelReq < 0 {
		fuelReq = 0
	}
	if mass == 0 {
		return agg + fuelReq
	}
	return calculateFuelRec(agg+fuelReq, fuelReq)
}

// CalculateFuel computes fuel required per module
func CalculateFuel(mass int) int {
	return calculateFuelRec(0, mass)
}

//SumFuelRequirement reads in module weights
func SumFuelRequirement(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed opening file: ", err)
	}

	scanner := bufio.NewScanner(file)
	var total int
	for scanner.Scan() {
		weight, err := strconv.ParseInt(scanner.Text(), 0, 64)
		if err != nil {
			fmt.Println("Error parsing weight", err)
		}
		total += CalculateFuel(int(weight))
	}

	return total
}

func main() {
	fmt.Println(SumFuelRequirement("input.txt"))
}
