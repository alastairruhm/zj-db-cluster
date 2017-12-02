package command

import (
	"bytes"
	"fmt"
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

func TestVersionFlags(t *testing.T) {
	RootCmd := cmd.RootCmd
	Convey("TestVersionCmd", t, func() {
		// some global val
		Convey("version command should work properly", func() {
			buf := new(bytes.Buffer)
			RootCmd.SetArgs([]string{"--version"})
			RootCmd.SetOutput(buf)

			err := RootCmd.Execute()

			if err != nil {
				t.Error(err)
			}

			actual := buf.String()
			expected := "zj-db-cluster version " + cmd.VERSION + "\n"
			So(actual, ShouldEqual, expected)
		})
	})
}

func TestConfigCmd(t *testing.T) {
	rootCmd := cmd.RootCmd
	Convey("TestConfigCmd", t, func() {
		configCmd := cmd.ConfigCmd

		Convey("config command will show usage of config: zj-db-cluster config", func() {
			buf := new(bytes.Buffer)
			rootCmd.SetArgs([]string{"config"})
			rootCmd.SetOutput(buf)

			err := rootCmd.Execute()

			if err != nil {
				t.Error(err)
			}

			actual := buf.String()
			expected := configCmd.UsageString()
			So(actual, ShouldEqual, expected)
		})

		Convey("config help command will show help of config: zj-db-cluster help config", func() {
			buf := new(bytes.Buffer)
			rootCmd.SetArgs([]string{"help", "config"})
			rootCmd.SetOutput(buf)

			err := rootCmd.Execute()

			if err != nil {
				t.Error(err)
			}

			actual := buf.String()
			expected := fmt.Sprintf("%s\n%s", configCmd.Long, configCmd.UsageString())
			So(actual, ShouldEqual, expected)
		})
	})
}
