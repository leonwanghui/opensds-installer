package executor

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/leonwanghui/opensds-installer/api"
)

func CreateRootBucket(name string) error {
	return createBucket(name, "root")
}

func createBucket(name, bType string) error {
	if _, err := execCmd("crush add-bucket", name, bType); err != nil {
		fmt.Println("When executing create bucket command:", err)
		return err
	}
	return nil
}

func RemoveBucket(name string) error {
	if _, err := execCmd("crush remove", name); err != nil {
		fmt.Println("When executing remove bucket command:", err)
		return err
	}
	return nil
}

func AddOsdInBucket(id, size, bName, bType string) error {
	bucket := bType + "=" + bName
	if _, err := execCmd("crush add", id, size, bucket); err != nil {
		fmt.Println("When executing add osd in bucket command:", err)
		return err
	}
	return nil
}

func SetPoolRule(name, ruleName string) error {
	if _, err := execCmd("pool set", name, "crush_rule", ruleName); err != nil {
		fmt.Println("When executing set pool rule command:", err)
		return err
	}
	return nil
}

func CreateReplicatedCrushMapRule(req *api.CrushMapRuleRequest) error {
	return createCrushMapRule("create-replicated", req)
}

func CreateSimpleCrushMapRule(req *api.CrushMapRuleRequest) error {
	return createCrushMapRule("create-simple", req)
}

func createCrushMapRule(ruleType string, req *api.CrushMapRuleRequest) error {
	if _, err := execCmd("crush rule", ruleType, req.RuleName, req.BucketName, req.BucketType, req.DeviceClass); err != nil {
		fmt.Println("When executing create crush map rule command:", err)
		return err
	}
	return nil
}

func GetOsdMetadataList() (api.OsdMetadataList, error) {
	// Get osd location in osd metadata
	ret, err := execCmd("metadata")
	if err != nil {
		fmt.Println("When executing get osd metadata command:", err)
		return nil, err
	}
	metaList := api.OsdMetadataList{}
	if err = json.Unmarshal(ret, &metaList); err != nil {
		return nil, err
	}

	// Get osd capacity in osd tree
	ret, err = execCmd("tree")
	if err != nil {
		fmt.Println("When executing get osd tree command:", err)
		return nil, err
	}

	// Merge them into osd metadata list
	retSlice := strings.Split(string(ret), "\n")
	for _, meta := range metaList {
		osdName := "osd." + fmt.Sprint(meta["id"])
		for _, ret := range retSlice {
			if strings.Contains(ret, osdName) {
				slice := strings.Fields(ret)
				meta["size"] = slice[1]
				break
			}
		}
		return nil, fmt.Errorf("Cannot find osd %v in tree!", osdName)
	}

	return metaList, nil
}

func execCmd(cmd ...string) ([]byte, error) {
	return exec.Command("ceph", "osd", strings.Join(cmd, " ")).Output()
}
