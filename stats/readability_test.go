package stats

import (
	"math"
	"testing"
)

func TestWordCount(t *testing.T) {
	expected := 10
	result := WordCount("the fox jumped over the road. That's pretty cool...	yup?")

	if result != expected {
		t.Errorf("expected %v, got %v\n", expected, result)
	}
}

func TestSentenceCount(t *testing.T) {
	expected := 6
	result := SentenceCount(
		`Hey this is a unit test. Just a simple unit test...
		nothing to say just making some tests!
		Tests!Yay! This is a test?`)

	if result != expected {
		t.Errorf("expected %v, got %v\n", expected, result)
	}
}

func TestSyllableCount(t *testing.T) {
	if !(SyllableCount("logorrhoea") == 4 &&
		SyllableCount("used") == 1 &&
		SyllableCount("makes") == 1 &&
		SyllableCount("themselves") == 2 &&
		SyllableCount("ion") == 2) {
		t.Fail()
	}
}

var testText = `The Flesch/Flesch–Kincaid readability tests are readability tests designed to indicate comprehension difficulty when reading a passage of contemporary academic English. There are two tests, the Flesch Reading Ease, and the Flesch–Kincaid Grade Level. Although they use the same core measures (word length and sentence length), they have different weighting factors. The results of the two tests correlate approximately inversely: a text with a comparatively high score on the Reading Ease test should have a lower score on the Grade Level test. Rudolf Flesch devised both systems while J. Peter Kincaid developed the latter for the United States Navy. Such readability tests suggest that many Wikipedia articles may be "too sophisticated" for their readers.`

func TestReadingEase(t *testing.T) {
	expected := 46.3
	result := ReadingEase(testText)

	if math.Abs(expected-result) >= 0.1 {
		t.Errorf("expected %v, got %v\n", expected, result)
	}
}

func TestReadingGradeLevel(t *testing.T) {
	expected := 10.8
	result := ReadingGradeLevel(testText)

	if math.Abs(expected-result) >= 0.1 {
		t.Errorf("expected %v, got %v\n", expected, result)
	}
}
