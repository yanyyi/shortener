package connect

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGet(t *testing.T) {
	c.Convey("基础用例", t, func() {
		url := "https://www.liwenzhou.com/posts/Go/unit-test-5/"
		got := Get(url)
		// 断言
		c.So(got, c.ShouldEqual, true)
		c.ShouldBeTrue(got)
	})

	c.Convey("url请求不通示例", t, func() {
		url := "posts/Go/unit-test-5/"
		got := Get(url)
		// 断言
		c.ShouldBeFalse(got)
	})
}
