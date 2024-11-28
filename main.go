package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rivo/uniseg"
)

var ErrFileNotFound = errors.New("error, cannot read file")

func main() {
	var pathVar string
	var counterTypeVar string

	flag.StringVar(&pathVar, "path", "", "Path to the file to analyze")
	flag.StringVar(&counterTypeVar, "type", "", "Type of counter to use: words, lines, characters")
	flag.Parse()

	if pathVar == "" {
		fmt.Println("Error: You must provide a valid file path using the -path flag.")
		os.Exit(1)
	}

	switch strings.ToLower(counterTypeVar) {
	case "words":
		words, err := WordCounter(pathVar)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("%v has %v words", pathVar, words)
		}
	case "lines":
		lines, err := LineCounter(pathVar)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("%v has %v lines", pathVar, lines)
		}
	case "characters":
		characters, err := CharacterCounter(pathVar)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("%v has %v characters", pathVar, characters)
		}
	case "":
		AllCounter(pathVar)
	default:
		fmt.Printf("Error: Unknown type %s. Use 'words', 'lines', or 'characters'.\n", counterTypeVar)

	}

}

func AllCounter(pathToFile string) {
	characters, err := CharacterCounter(pathToFile)
	if err != nil {
		panic(err)
	}
	lines, _ := LineCounter(pathToFile)
	words, _ := WordCounter(pathToFile)

	fmt.Printf("%v has %v characters, %v lines, and %v words \n", pathToFile, characters, lines, words)

}

func CharacterCounter(pathToFile string) (int, error) {
	file, err := os.ReadFile(pathToFile)
	if err != nil {
		return 0, ErrFileNotFound
	}
	stringifyContent := string(file)
	stringWithoutSpace := strings.Replace(stringifyContent, " ", "", -1)
	characterCount := uniseg.GraphemeClusterCount(stringWithoutSpace)
	return characterCount, nil
}

func LineCounter(pathToFile string) (int, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return 0, ErrFileNotFound
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var noOfLines int
	for scanner.Scan() {
		if len(scanner.Text()) != 0 {
			noOfLines += 1

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("An error occurred during scanning")
	}
	return noOfLines, nil
}

func WordCounter(pathToFile string) (int, error) {
	file, err := os.Open(pathToFile)

	if err != nil {
		return 0, ErrFileNotFound
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var noOfWords int
	for scanner.Scan() {
		noOfWords += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return noOfWords, nil
}
