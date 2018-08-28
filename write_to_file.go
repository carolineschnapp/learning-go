package main

import (
	"io"
	"os"
)

func main() {
	writeToFile("Hey Caro, comment ça va?")
}

func writeToFile(text string) (err error) {
	var file *os.File
	file, err = os.Create("/Users/carolineschnapp/caro.txt")
	if err != nil {
		return
	}

	defer file.Close()

	_, err = io.WriteString(file, text)
	return
}
