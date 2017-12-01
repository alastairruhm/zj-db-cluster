package cmd

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/alastairruhm/zj-db-cluster/client"
	"github.com/alastairruhm/zj-db-cluster/config"
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

			clientStub := Patch(client.NewChecker, func(_ config.ClusterConfig) client.Checker {
				return mockClient
			})
			defer clientStub.Unpatch()

			// gostub 和 monkey 两种方法都可以注入 mock 对象
			// 但是 gostub 需要改写 func 的定义方式
			// clientStub := gostub.StubFunc(&client.NewChecker, mockClient)
			// defer clientStub.Reset()

			checkArgsStub := Patch(CheckClusterNameArgs, func() error {
				return nil
			})
			defer checkArgsStub.Unpatch()

			loadStub := Patch(LoadConfig, func() error {
				return nil
			})
			defer loadStub.Unpatch()

			c := client.NewChecker(config.Config)
			t.Logf("%+v", reflect.TypeOf(c))

			buf := new(bytes.Buffer)
			CheckConnectionCmd.SetOutput(buf)
			CheckConnection(CheckConnectionCmd, nil)
			t.Logf("%+v", buf)
			actual := buf.String()
			t.Logf("%+v", actual)
			So(actual, ShouldEqual, result)
		})
	})

}
