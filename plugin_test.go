package main

import (
  "testing"

  "github.com/cloudfoundry/cli/plugin/fakes"
  . "github.com/smartystreets/goconvey/convey"
)

var (
  cliConn *fakes.FakeCliConnection
)

func TestNoApp(t *testing.T) {
  setup()
  Convey("checkArgs should not return error with tree dora", t, func() {
    err := checkArgs(cliConn, []string{"tree", "dora"})
    So(err, ShouldBeNil)
  })

  Convey("checkArgs should return error with dora", t, func() {
    err := checkArgs(cliConn, []string{"tree"})
    So(err, ShouldNotBeNil)
  })
}

func setup() {
  cliConn = &fakes.FakeCliConnection{}
}
