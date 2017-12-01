package cmd

import (
	"bytes"
	"testing"

	"github.com/prashantv/gostub"

	"github.com/alastairruhm/zj-db-cluster/client"
	mock_client "github.com/alastairruhm/zj-db-cluster/test/mock/mock_client"

	. "github.com/bouk/monkey"
	. "github.com/smartystreets/goconvey/convey"

	. "github.com/golang/mock/gomock"
)

func TestCheckCmd(t *testing.T) {

	Convey("Test check connection command", t, func() {

		Convey("print result normally", func() {
			ctrl := NewController(t)
			defer ctrl.Finish()

			result := "check result"

			mockClient := mock_client.NewMockChecker(ctrl)
			mockClient.EXPECT().CheckConnection().Return(result)

			// clientStub := Patch(client.NewChecker, func(_ config.ClusterConfig) client.Checker {
			// 	return mockClient
			// })
			// defer clientStub.Unpatch()
			clientStub := gostub.StubFunc(&client.NewChecker, mockClient)
			defer clientStub.Reset()

			checkArgsStub := Patch(CheckClusterNameArgs, func() error {
				return nil
			})
			defer checkArgsStub.Unpatch()

			loadStub := Patch(LoadConfig, func() error {
				return nil
			})
			defer loadStub.Unpatch()

			buf := new(bytes.Buffer)
			RootCmd.SetOutput(buf)

			CheckConnection(RootCmd, nil)

			actual := buf.String()

			So(actual, ShouldEqual, result)

		})
	})

}
