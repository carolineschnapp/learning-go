package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// What we will use:
	// - strings: func TrimSpace(s string) string
	// - strconv: func Atoi(s string) (int, error)
	// - io: to access the variable os.Stdin
	// - bufio: func NewReader(rd io.Reader) *Reader
	// - bufio: func (b *Reader) ReadString(delim byte) (string, error)
	// - math/rand: func Seed(seed int64)
	// - math/rand: func Intn(n int) int
	// - time: func Now() Time
	// - time: func (t Time) Unix() int64
	randNumber := randomNumber()
	// fmt.Printf("Random number between 1 and 100 is %d.\n", randNumber)
	// => Random number between 1 and 100 is 48.
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		guess := guess(reader)
		if guess == randNumber {
			fmt.Println("You have guessed correctly!")
			return
		} else if guess < randNumber {
			fmt.Println("Higher baby!")
		} else {
			fmt.Println("Lower baby!")
		}
	}
	fmt.Println("You have ran out of guesses. Good day to you, Sir!")
}

func randomNumber() (randNumber int) {
	now := time.Now()
	unixTime := now.Unix() // Number of seconds since Thursday, January 1st, 1970.
	rand.Seed(unixTime)    // To get a different number every time we run the program.
	randNumber = rand.Intn(100) + 1
	return
}

func guess(reader *bufio.Reader) (guess int) {
	fmt.Println("Can you guess which number it is?")
	value, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	trimmedString := strings.TrimSpace(value)
	guess, err = strconv.Atoi(trimmedString)
	if err != nil {
		log.Fatal("You must enter a number!")
	}
	return
}
