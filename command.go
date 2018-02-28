package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func SetPoolRule(name, ruleName string) error {
	if _, err := execCmd("pool set", name, "crush_rule", ruleName); err != nil {
		return err
	}
	return nil
}

type CrushMapRuleRequest struct {
	RuleName, BucketName, BucketType, DeviceClass string
}

func CreateReplicatedCrushMapRule(req *CrushMapRuleRequest) error {
	return createCrushMapRule("create-replicated", req)
}

func CreateSimpleCrushMapRule(req *CrushMapRuleRequest) error {
	return createCrushMapRule("create-simple", req)
}

func createCrushMapRule(ruleType string, req *CrushMapRuleRequest) error {
	if _, err := execCmd("crush rule", ruleType, req.RuleName, req.BucketName, req.BucketType, req.DeviceClass); err != nil {
		return err
	}
	return nil
}

type OsdMetadataList []map[string]interface{}

func getOsdMetadataList() (OsdMetadataList, error) {
	ret, err := execCmd("metadata")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	metaList := OsdMetadataList{}
	if err = json.Unmarshal(ret, &metaList); err != nil {
		return nil, err
	}

	return metaList, nil
}

func execCmd(cmd ...string) ([]byte, error) {
	return exec.Command("ceph", "osd", strings.Join(cmd, " ")).Output()
}
