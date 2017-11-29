package cmd

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVersionCmd(t *testing.T) {
	Convey("TestVersionCmd", t, func() {
		// some global val
		Convey("version command should work properly", func() {
			buf := new(bytes.Buffer)
			RootCmd.SetArgs([]string{"version"})
			RootCmd.SetOutput(buf)

			err := RootCmd.Execute()

			if err != nil {
				t.Error(err)
			}

			actual := buf.String()
			expected := "zijin database cluster tool version " + VERSION + "\n"
			So(actual, ShouldEqual, expected)
		})
	})
}
