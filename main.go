package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)
import "os"
import "bufio"

func main() {
	// https://www.codingame.com/ide/puzzle/ascii-art
	input := readInput()
	asciiArt := calculateAsciiArt(input)
	print(asciiArt)
}

func readInput() (input Input) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &input.LetterWidth)

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &input.LetterHeight)

	scanner.Scan()
	input.Text = scanner.Text()

	for i := 0; i < input.LetterHeight; i++ {
		scanner.Scan()
		ROW := scanner.Text()
		input.Letters = append(input.Letters, ROW)
	}

	return
}

type Input struct {
	LetterWidth  int
	LetterHeight int
	Text         string
	Letters      []string
}

func calculateAsciiArt(input Input) (res []string) {

	fmt.Fprintln(os.Stderr, input.ToString())

	res = initOutput(input.LetterHeight)
	textLetters := strings.Split(input.Text, "")

	for _, l := range textLetters {
		l = sanitizeLetter(l)
		letterIndex := getLetterIndex(l)

		for i := 0; i < input.LetterHeight; i++ {
			res[i] += getLineOfLetter(input, letterIndex, i)
		}
	}

	return
}

func print(asciiArt []string) {
	for _, line := range asciiArt {
		fmt.Println(line)
	}
}

func getLineOfLetter(input Input, letterIndex int, i int) string {
	indexStart := letterIndex * input.LetterWidth
	indexEnd := (letterIndex + 1) * input.LetterWidth

	fmt.Fprintf(os.Stderr, "start index: %d, end index: %d\n", indexStart, indexEnd)

	lineOfLetter := input.Letters[i][indexStart:indexEnd]

	return lineOfLetter
}

func getLetterIndex(l string) int {
	r, _ := utf8.DecodeRune([]byte(l))
	letterIndex := int(r) - 65

	if letterIndex < 0 {
		// question Mark is at the end of the letter list
		letterIndex = 26
	}

	fmt.Fprintf(os.Stderr, "Letter: %s, rune: %v, letter index: %d\n", l, r, letterIndex)

	return letterIndex
}

func sanitizeLetter(l string) string {
	l = toUpperCase(l)
	l = invalidLetterToQuestionMark(l)
	return l
}

func toUpperCase(l string) string {
	if isLowerCaseLetter, _ := regexp.MatchString(`[a-z]`, l); isLowerCaseLetter {
		l = strings.ToUpper(l)
	}
	return l
}

func invalidLetterToQuestionMark(l string) string {
	if isNotLetter, _ := regexp.MatchString(`[^A-Z]`, l); isNotLetter {
		l = "?"
	}
	return l
}

func initOutput(height int) (output []string) {
	for i := 0; i < height; i++ {
		output = append(output, "")
	}
	return
}

func (input Input) ToString() (s string) {
	s = fmt.Sprintf("LetterWidth: %d, LetterHeight: %d\nText: %s\nLetters:\n", input.LetterWidth, input.LetterHeight, input.Text)
	for i := 0; i < input.LetterHeight; i++ {
		s += fmt.Sprintf("%s\n", input.Letters[i])
	}
	return
}
