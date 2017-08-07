// +build !linux

package main

import (
	"github.com/codegangsta/cli"
	"github.com/sirupsen/logrus"
)

func getDefaultID() string {
	return ""
}

var (
	checkpointCommand cli.Command
	eventsCommand     cli.Command
	restoreCommand    cli.Command
	specCommand       cli.Command
	killCommand       cli.Command
)

func runAction(*cli.Context) {
	logrus.Fatal("Current OS is not supported yet")
}
