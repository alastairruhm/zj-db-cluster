package command

import (
	"bytes"
	"testing"

	"github.com/alastairruhm/zj-db-cluster/cmd"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckCmd(t *testing.T) {
	rootCmd := cmd.RootCmd
	Convey("Test check command", t, func() {
		checkCmd := cmd.CheckCmd

		Convey("[zj-db-cluster check]: show usage of check", func() {
			buf := new(bytes.Buffer)
			rootCmd.SetArgs([]string{"check"})
			rootCmd.SetOutput(buf)

			err := rootCmd.Execute()

			if err != nil {
				t.Error(err)
			}

			actual := buf.String()
			expected := checkCmd.UsageString()
			So(actual, ShouldEqual, expected)
		})
	})

}
