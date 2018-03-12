package parser

import (
	"fmt"
	"io/ioutil"

	"github.com/leonwanghui/opensds-installer/api"
	"gopkg.in/yaml.v2"
)

func ParsePoolTopology(p string) (map[string]api.Location, error) {
	var conf = &api.Config{}

	if err := parse(conf, p); err != nil {
		return nil, err
	}
	return conf.PoolTopology, nil
}

func parse(conf interface{}, p string) error {
	confYaml, err := ioutil.ReadFile(p)
	if err != nil {
		fmt.Errorf("Read config yaml file (%s) failed, reason:(%v)", p, err)
		return err
	}
	if err = yaml.Unmarshal(confYaml, conf); err != nil {
		fmt.Errorf("Parse error: %v", err)
		return err
	}
	return nil
}
