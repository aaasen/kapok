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

func FleschKincaidEase(text string) float64 {
	words := Words(text)
	numWords := float64(len(words))
	numSentences := float64(len(Sentences(text)))
	numSyllables := 0

	for _, word := range words {
		numSyllables += SyllableCount(word)
	}

	return 206.835 - 1.015*(numWords/numSentences) - 84.6*(float64(numSyllables)/numWords)

}

func cleanWord(word string) string {
	return strings.Trim(word, wordChars)
}
