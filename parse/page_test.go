package parse

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsTitle(t *testing.T) {
	Convey("isTitle should distinguish between important and unimportant titles", t, func() {
		So(isTitle([]byte("Hello")), ShouldBeTrue)
		So(isTitle([]byte("Help:Contents")), ShouldBeFalse)
		So(isTitle([]byte("#See also")), ShouldBeFalse)
		So(isTitle([]byte("{{TALKPAGENAME}}|Discussion")), ShouldBeFalse)
		So(isTitle([]byte("/example")), ShouldBeFalse)
		So(isTitle([]byte("/example/")), ShouldBeFalse)
		So(isTitle([]byte(":Category:Help")), ShouldBeFalse)
		So(isTitle([]byte("media:example.jpg")), ShouldBeFalse)
		So(isTitle([]byte("Special:MyPage")), ShouldBeFalse)
	})
}
