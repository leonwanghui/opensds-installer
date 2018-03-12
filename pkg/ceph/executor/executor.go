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

func CreateHostBucket(name string) error {
	return createBucket(name, "host")
}

func createBucket(name, bType string) error {
	if _, err := execOsdCmd("crush", "add-bucket", name, bType); err != nil {
		fmt.Println("When executing create bucket command:", err)
		return err
	}
	return nil
}

func RemoveBucket(name string) error {
	if _, err := execOsdCmd("crush", "remove", name); err != nil {
		fmt.Println("When executing remove bucket command:", err)
		return err
	}
	return nil
}

func AddOsdInHostBucket(id, size, hostName string) error {
	return addOsdInBucket(id, size, hostName, "host")
}

func addOsdInBucket(id, size, bName, bType string) error {
	bucket := bType + "=" + bName
	if _, err := execOsdCmd("crush", "add", id, size, bucket); err != nil {
		fmt.Println("When executing add osd in bucket command:", err)
		return err
	}
	return nil
}

func MoveHostInRootBucket(hostName, rootName string) error {
	rootBucket := "root=" + rootName
	if _, err := execOsdCmd("crush", "move", hostName, rootBucket); err != nil {
		fmt.Println("When executing move host in root bucket command:", err)
		return err
	}
	return nil
}

func SetPoolRule(name, ruleName string) error {
	if _, err := execOsdCmd("pool", "set", name, "crush_rule", ruleName); err != nil {
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
	if _, err := execOsdCmd("crush", "rule", ruleType, req.RuleName, req.BucketName, req.BucketType, req.DeviceClass); err != nil {
		fmt.Println("When executing create crush map rule command:", err)
		return err
	}
	return nil
}

func GetOsdMetadataList() (api.OsdMetadataList, error) {
	// Get osd location in osd metadata
	ret, err := execOsdCmd("metadata")
	if err != nil {
		fmt.Println("When executing get osd metadata command:", err)
		return nil, err
	}
	metaList := api.OsdMetadataList{}
	if err = json.Unmarshal(ret, &metaList); err != nil {
		return nil, err
	}

	// Get osd capacity in osd tree
	ret, err = execOsdCmd("tree")
	if err != nil {
		fmt.Println("When executing get osd tree command:", err)
		return nil, err
	}

	// Merge them into osd metadata list
	retSlice := strings.Split(string(ret), "\n")
	loc := func() int {
		slice := strings.Fields(retSlice[0])
		for i := range slice {
			if slice[i] == "WEIGHT" {
				return i
			}
		}
		return 1
	}()
	for _, meta := range metaList {
		osdName := "osd." + fmt.Sprint(meta["id"])
		for _, ret := range retSlice {
			if strings.Contains(ret, osdName) {
				slice := strings.Fields(ret)
				meta["size"] = slice[loc]
				break
			}
		}
		if _, ok := meta["size"]; !ok {
			return nil, fmt.Errorf("Cannot find osd %v in tree!", osdName)
		}
	}

	return metaList, nil
}

func execOsdCmd(cmd ...string) ([]byte, error) {
	cmd = append([]string{"osd"}, cmd...)
	return exec.Command("ceph", cmd...).Output()
}
