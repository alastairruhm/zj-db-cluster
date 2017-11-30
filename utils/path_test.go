package utils

import (
	"os"
	"testing"

	. "github.com/bouk/monkey"
	. "github.com/smartystreets/goconvey/convey"
)

const any = "any"

func TestExist(t *testing.T) {
	Convey("Test Exist function", t, func() {
		// some global val
		Convey("os.Stat error", func() {
			stat := Patch(os.Stat, func(_ string) (os.FileInfo, error) {
				return nil, &os.PathError{}
			})
			defer stat.Unpatch()
			testify := Patch(os.IsNotExist, func(_ error) bool {
				return true
			})
			defer testify.Unpatch()
			actual, err := Exist(any)
			So(actual, ShouldEqual, false)
			So(err, ShouldBeNil)
		})

		Convey("os.Stat not error", func() {
			stat := Patch(os.Stat, func(_ string) (os.FileInfo, error) {
				return nil, nil
			})
			defer stat.Unpatch()
			testify := Patch(os.IsNotExist, func(_ error) bool {
				return true
			})
			defer testify.Unpatch()
			actual, err := Exist(any)
			So(actual, ShouldEqual, true)
			So(err, ShouldBeNil)
		})

		Convey("os.Stat not error", func() {
			stat := Patch(os.Stat, func(_ string) (os.FileInfo, error) {
				return nil, nil
			})
			defer stat.Unpatch()
			testify := Patch(os.IsNotExist, func(_ error) bool {
				return true
			})
			defer testify.Unpatch()
			actual, err := Exist(any)
			So(actual, ShouldEqual, true)
			So(err, ShouldBeNil)
		})
	})
}
