package generator

import (
	"github.com/leonwanghui/opensds-installer/api"
	"github.com/leonwanghui/opensds-installer/pkg/executor"
)

func GenerateCrushMap(pLocMap map[string]api.Location, mapList api.OsdLocationMapList) error {
	for pName, pLoc := range pLocMap {
		// Create empty root bucket using user-specified pool name.
		if err := executor.CreateRootBucket(pName); err != nil {
			return err
		}

		// If user doesn't configure disks in pool, the root bucket would be
		// deleted because it is unneccessary for ceph cluster to create an
		// empty pool.
		if len(pLoc.Disks) == 0 {
			if err := executor.RemoveBucket(pName); err != nil {
				return err
			}
			continue
		}

		// Reconstruct the root bucket with devicepath and host according to
		// user's configuration.
		for _, disk := range pLoc.Disks {
			for _, dMap := range mapList {
				if dMap["device_path"] == disk["path"] &&
					dMap["hostname"] == disk["hostnames"] {
					if err := executor.AddOsdInRootBucket(
						dMap["id"],
						dMap["size"],
						pName); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
