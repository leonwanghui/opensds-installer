package main

import (
	"fmt"

	"github.com/leonwanghui/opensds-installer/pkg/collector"
	"github.com/leonwanghui/opensds-installer/pkg/generator"
	"github.com/leonwanghui/opensds-installer/pkg/parser"
)

var config_file = "examples/template.yml"

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	pLocMap, err := parser.ParsePoolTopology(config_file)
	if err != nil {
		return fmt.Errorf("When parsing pool topology:", err)
	}

	mapList, err := collector.CollectOsdLocation()
	if err != nil {
		return fmt.Errorf("When collecting osd location:", err)
	}

	return generator.GenerateCrushMap(pLocMap, mapList)
}
