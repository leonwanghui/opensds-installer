package main

import (
	"flag"
	"fmt"

	ctrl "github.com/leonwanghui/opensds-installer/pkg/controller"
)

var (
	inputFile, driverName string
)

func init() {
	flag.StringVar(&inputFile, "input-file", "examples/template.yml", "the path of storage pool config file")
	flag.StringVar(&driverName, "driver", "ceph", "the type of opensds backend service")

	flag.Parse()
}

func main() {
	var c ctrl.Controller

	switch driverName {
	case "ceph":
		c = ctrl.NewCephController(inputFile)
	default:
		panic(fmt.Errorf("Backend type (%s) not supported!", driverName))
	}

	if err := c.Run(); err != nil {
		panic(err)
	}
}
