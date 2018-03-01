package generator

import (
	"fmt"

	"github.com/leonwanghui/opensds-installer/api"
	"github.com/leonwanghui/opensds-installer/pkg/executor"
)

func GenerateCrushMap(pLocMap map[string]api.Location, mapList api.OsdLocationMapList) error {
	for pName, pLoc := range pLocMap {
		// If user doesn't configure disks in pool, the root bucket would be
		// skipped because it is unneccessary for ceph cluster to create an
		// empty pool.
		if len(pLoc.Disks) == 0 {
			continue
		}

		// Create empty root bucket using user-specified pool name.
		if err := executor.CreateRootBucket(pName); err != nil {
			return err
		}

		// Reconstruct the root bucket with devicepath and host according to
		// user's configuration.
		for _, disk := range pLoc.Disks {
			for _, dMap := range mapList {
				fmt.Println("disk =", disk, "dmap =", dMap)
				if dMap["device_path"] == disk["path"] &&
					dMap["hostname"] == disk["hostname"] {
					if err := executor.AddOsdInRootBucket(
						dMap["id"],
						dMap["size"],
						pName); err != nil {
						return err
					}
				}
			}
		}

		// Create crush map rule using root bucket created.
		ruleName := pName + "-rule"
		req := &api.CrushMapRuleRequest{
			RuleName:   ruleName,
			BucketName: pName,
			BucketType: "root",
		}
		if err := executor.CreateReplicatedCrushMapRule(req); err != nil {
			return err
		}

		// Set crush map rule to the pool.
		if err := executor.SetPoolRule(pName, ruleName); err != nil {
			return err
		}
	}

	return nil
}
