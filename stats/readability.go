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

// Words returns a slice of the words in a given string.
func Words(text string) []string {
	return wordRegex.Split(text, -1)
}

// Sentences returns a slice of the sentences in a given string.
func Sentences(text string) []string {
	sentences := sentenceRegex.Split(text, -1)

	if sentences[len(sentences)-1] == "" {
		return sentences[0 : len(sentences)-1]
	}

	return sentences
}

// SyllableCount returns the number of syllables in a given string.
// Dipthongs are not taken into account.
func SyllableCount(word string) int {
	if len(word) <= 3 {
		return 1
	}

	word = endingRegex.ReplaceAllString(word, "")
	word = staringYRegex.ReplaceAllString(word, "")

	return len(vowelRegex.FindAllString(word, -1))
}

// FleschKincaidEase computes the ease of reading a given text.
// The algorithm is explained in detail here:
// http://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_tests
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
