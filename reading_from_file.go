package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var lines, err = readLinesfromFile("/Users/carolineschnapp/caro.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(lines); i++ {
		fmt.Printf("Line %d: %s\n", i, lines[i])
	}
}

func readLinesfromFile(filepath string) (lines []string, err error) {
	var file *os.File
	file, err = os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		err = scanner.Err()
	}
	return
}
