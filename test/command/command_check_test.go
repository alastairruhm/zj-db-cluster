package command

import (
	"bytes"
	"testing"

	"github.com/alastairruhm/zj-db-cluster/client"
	"github.com/alastairruhm/zj-db-cluster/config"
	mock_client "github.com/alastairruhm/zj-db-cluster/test/mock/mock_client"

	"github.com/alastairruhm/zj-db-cluster/cmd"
	. "github.com/bouk/monkey"
	. "github.com/smartystreets/goconvey/convey"

	. "github.com/golang/mock/gomock"
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

	Convey("Test check connection command", t, func() {
		checkConnectionCmd := cmd.CheckConnectionCmd
		Convey("print result normally", func() {
			ctrl := NewController(t)
			defer ctrl.Finish()
			result := "check result"
			mockClient := mock_client.NewMockChecker(ctrl)
			mockClient.EXPECT().CheckConnection().Return(result)

			clientStub := Patch(client.NewChecker, func(_ config.ClusterConfig) client.Checker {
				return mockClient
			})
			defer clientStub.Unpatch()

			// checkArgsStub := Patch(cmd.CheckClusterNameArgs, func() error {
			// 	return nil
			// })
			// defer checkArgsStub.Unpatch()

			// loadStub := Patch(cmd.LoadConfig, func() error {
			// 	return nil
			// })
			// defer loadStub.Unpatch()

			buf := new(bytes.Buffer)
			// rootCmd.SetArgs([]string{"check", "connection"})
			checkConnectionCmd.SetOutput(buf)
			err := checkConnectionCmd.Execute()
			if err != nil {
				t.Error(err)
			}
			actual := buf.String()
			expected := result
			So(actual, ShouldEqual, expected)

		})
	})

}
