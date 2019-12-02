package advent

import (
	"fmt"
	"io/ioutil"
)

// CalculateFuel computes fuel required per module
func CalculateFuel(mass int) int {
	return mass/3 - 2
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("I've read:", string(data))
}
