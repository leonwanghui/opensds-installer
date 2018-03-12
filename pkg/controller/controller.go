package controller

import (
	"fmt"

	"github.com/leonwanghui/opensds-installer/pkg/ceph/collector"
	"github.com/leonwanghui/opensds-installer/pkg/ceph/generator"
	"github.com/leonwanghui/opensds-installer/pkg/ceph/parser"
)

type Controller interface {
	Run() error
}

func NewCephController(in string) Controller {
	return &CephController{inputFile: in}
}

type CephController struct {
	inputFile string
}

func (c *CephController) Run() error {
	pLocMap, err := parser.ParsePoolTopology(c.inputFile)
	if err != nil {
		return fmt.Errorf("When parsing pool topology:", err)
	}

	mapList, err := collector.CollectOsdLocation()
	if err != nil {
		return fmt.Errorf("When collecting osd location:", err)
	}

	return generator.GenerateCrushMap(pLocMap, mapList)
}
