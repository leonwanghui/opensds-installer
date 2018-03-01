package main

import (
	"flag"
	"fmt"

	"github.com/leonwanghui/opensds-installer/pkg/collector"
	"github.com/leonwanghui/opensds-installer/pkg/generator"
	"github.com/leonwanghui/opensds-installer/pkg/parser"
)

var inputFile string

func init() {
	flag.StringVar(&inputFile, "input-file", "examples/template.yml", "the path of storage pool config file")

	flag.Parse()
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	pLocMap, err := parser.ParsePoolTopology(inputFile)
	if err != nil {
		return fmt.Errorf("When parsing pool topology:", err)
	}

	mapList, err := collector.CollectOsdLocation()
	if err != nil {
		return fmt.Errorf("When collecting osd location:", err)
	}

	return generator.GenerateCrushMap(pLocMap, mapList)
}
