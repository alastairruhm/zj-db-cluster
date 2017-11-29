package main

import (
	"bytes"
	"testing"

	"github.com/alastairruhm/zj-db-cluster/cmd"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRootCmd(t *testing.T) {
	RootCmd := cmd.RootCmd
	Convey("TestRootCmd", t, func() {
		// some global val
		Convey("default command will show usage string", func() {
			buf := new(bytes.Buffer)
			RootCmd.SetArgs([]string{""})
			RootCmd.SetOutput(buf)

			err := RootCmd.Execute()

			if err != nil {
				t.Error(err)
			}

			actual := buf.String()
			expected := RootCmd.UsageString()
			So(actual, ShouldEqual, expected)
		})
	})
}

func TestVersionCmd(t *testing.T) {
	RootCmd := cmd.RootCmd
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
			expected := "zijin database cluster tool version " + cmd.VERSION + "\n"
			So(actual, ShouldEqual, expected)
		})
	})
}
