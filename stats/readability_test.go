package stats

import (
	"log"
	"math"
	"testing"
)

func TestWordCount(t *testing.T) {
	expected := 10
	result := len(Words("the fox jumped over the road. That's pretty cool...	yup?"))

	if result != expected {
		t.Errorf("expected %v, got %v\n", expected, result)
	}
}

func TestSentences(t *testing.T) {
	expected := 6
	result := len(Sentences(
		`Hey this is a unit test. Just a simple unit test...
		nothing to say just making some tests!
		Tests!Yay! This is a test?`))

	if result != expected {
		t.Errorf("expected %v, got %v\n", expected, result)
	}
}

func TestSyllables(t *testing.T) {
	log.Println(SyllableCount("logorrhoea"))
	log.Println(SyllableCount("used"))
	log.Println(SyllableCount("makes"))
	log.Println(SyllableCount("themselves"))
	log.Println(SyllableCount("ion"))

	if !(SyllableCount("logorrhoea") == 4 &&
		SyllableCount("used") == 1 &&
		SyllableCount("makes") == 1 &&
		SyllableCount("themselves") == 2 &&
		SyllableCount("ion") == 2) {
		t.Fail()
	}
}

var testText = `The Flesch/Flesch–Kincaid readability tests are readability tests designed to indicate comprehension difficulty when reading a passage of contemporary academic English. There are two tests, the Flesch Reading Ease, and the Flesch–Kincaid Grade Level. Although they use the same core measures (word length and sentence length), they have different weighting factors. The results of the two tests correlate approximately inversely: a text with a comparatively high score on the Reading Ease test should have a lower score on the Grade Level test. Rudolf Flesch devised both systems while J. Peter Kincaid developed the latter for the United States Navy. Such readability tests suggest that many Wikipedia articles may be "too sophisticated" for their readers.`

func TestFleschKincaidEase(t *testing.T) {
	expected := 46.3
	result := FleschKincaidEase(testText)

	if math.Abs(expected-result) >= 0.1 {
		t.Errorf("expected %v, got %v\n", expected, result)
	}
}
