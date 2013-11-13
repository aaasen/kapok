package stats

import (
	"regexp"
	"strings"
)

var wordRegex = regexp.MustCompile("[ \t\r\n\v\f]+")
var wordChars = " \t\r\n\v\f.?!"

var sentenceRegex = regexp.MustCompile("[ \t\r\n\v\f]*[.?!]+[ \t\r\n\v\f]*")

var endingRegex = regexp.MustCompile("([^laeiouy]es|ed|[^laeoiuy]e)$")
var staringYRegex = regexp.MustCompile("^y")
var vowelRegex = regexp.MustCompile("[aeiouy]{1,2}")

func Words(text string) []string {
	return wordRegex.Split(text, -1)
}

func Sentences(text string) []string {
	sentences := sentenceRegex.Split(text, -1)

	if sentences[len(sentences)-1] == "" {
		return sentences[0 : len(sentences)-1]
	}

	return sentences
}

func SyllableCount(word string) int {
	if len(word) <= 3 {
		return 1
	}

	word = endingRegex.ReplaceAllString(word, "")
	word = staringYRegex.ReplaceAllString(word, "")

	return len(vowelRegex.FindAllString(word, -1))
}

func ReadingEase(text string) float64 {
	return 0.0
}

func ReadingGradeLevel(text string) float64 {
	return 0.0
}

func cleanWord(word string) string {
	return strings.Trim(word, wordChars)
}
