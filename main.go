package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var faceFlag = flag.String("f", "", "goat, demon, skelly, bat")
var inputFlag = flag.String("i", "", "misfortunes/misfortunes.txt")
var helpFlag = flag.Bool("h", false, "display help message")

func readQuotes(file string) [][]string {
	quotes := make([][]string, 0, 0)
	quote := make([]string, 0, 0)

	data, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	words := bufio.NewScanner(data)
	for i := 0; words.Scan(); {
		if words.Text() == "@" {
			if i == 1 {
				newQuoteSize := len(quote) - 1
				quote = quote[:newQuoteSize]
				quote[newQuoteSize-1] = strings.TrimRight(quote[newQuoteSize-1], "\t \n")
				quotes = append(quotes, quote)
				quote = nil
			}
			continue
		}
		quote = append(quote, words.Text()+"\n")
		i = 1
	}

	return quotes
}

func quoteRandomizer(quotes [][]string) []string {
	// function to select a random quote from the file
	rand.Seed(time.Now().UnixNano())
	randomQuote := quotes[(rand.Intn(len(quotes)))]
	return randomQuote
}

func main() {
	const (
		helpMessage = `Usage: misfortune
       -f goat | demon | skelly | bat (default: random)
       -i path to file (default: misfortunes/misfortunes.txt)
    `
		maxWidth = 60
	)

	var (
		file      string
		lineWidth int
	)

	flag.Parse()

	if *helpFlag {
		fmt.Println(helpMessage)
		os.Exit(0)
	}

	if *inputFlag == "" {
		file = "misfortunes/misfortunes.txt"
	} else {
		file = *inputFlag
	}

	quotes := readQuotes(file)
	chosenQuote := quoteRandomizer(quotes)
	chosenFace := avatar(*faceFlag)

	fmt.Println()
	for _, quoteLine := range chosenQuote {
		for _, quoteChar := range quoteLine {
			fmt.Printf("%c", quoteChar)
			switch lineWidth++; {
			case lineWidth > maxWidth && unicode.IsSpace(quoteChar):
				fmt.Println()
				fallthrough
			case quoteChar == '\n':
				lineWidth = 0
			}
		}
	}
	fmt.Printf("\n%s\n", chosenFace)
}
