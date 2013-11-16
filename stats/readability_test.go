package stats

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestWordCount(t *testing.T) {
	words := Words("the fox jumped over the road. That's pretty cool...	yup?")

	Convey("Words should be correct", t, func() {
		So(words, ShouldResemble, []string{"the", "fox", "jumped", "over", "the", "road", "That's", "pretty", "cool", "yup"})
	})
}

func TestSentences(t *testing.T) {
	sentences := Sentences(
		`Hey this is a unit test. Just a simple unit test...
		nothing to say just making some tests!
		Tests!Yay! This is a test?`)

	Convey("Sentences should be correct", t, func() {
		So(sentences, ShouldResemble, []string{
			"Hey this is a unit test",
			"Just a simple unit test",
			"nothing to say just making some tests",
			"Tests",
			"Yay",
			"This is a test",
		})
	})
}

func TestSyllables(t *testing.T) {
	Convey("Syllable counts should be correct", t, func() {
		So(SyllableCount("logorrhoea"), ShouldEqual, 4)
		So(SyllableCount("used"), ShouldEqual, 1)
		So(SyllableCount("makes"), ShouldEqual, 1)
		So(SyllableCount("themselves"), ShouldEqual, 2)

		Convey("Dipthongs should be considered", func() {
			So(SyllableCount("ion"), ShouldEqual, 2)
		})
	})
}

var testText = `The Flesch/Flesch–Kincaid readability tests are readability tests designed to indicate comprehension difficulty when reading a passage of contemporary academic English. There are two tests, the Flesch Reading Ease, and the Flesch–Kincaid Grade Level. Although they use the same core measures (word length and sentence length), they have different weighting factors. The results of the two tests correlate approximately inversely: a text with a comparatively high score on the Reading Ease test should have a lower score on the Grade Level test. Rudolf Flesch devised both systems while J. Peter Kincaid developed the latter for the United States Navy. Such readability tests suggest that many Wikipedia articles may be "too sophisticated" for their readers.`

func TestFleschKincaidEase(t *testing.T) {
	ease := FleschKincaidEase(testText)

	Convey("FleschKincaidEase should be within 2 points of the control", t, func() {
		So(ease, ShouldBeBetween, 44.3, 48.3)
	})
}
