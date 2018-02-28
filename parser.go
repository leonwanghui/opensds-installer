package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	PoolTopology map[string]Location `yaml:"pool,flow"`
}

type Location struct {
	Disks []map[string]string `yaml:"disks"`
}

func ParsePoolTopology(p string) (map[string]Location, error) {
	var conf = &Config{}

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
