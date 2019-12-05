package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readWirePaths(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var paths []string
	for scanner.Scan() {
		paths = append(paths, scanner.Text())
	}
	file.Close()

	return paths
}

func main() {
	fmt.Println(len(readWirePaths("input.txt")))
}
