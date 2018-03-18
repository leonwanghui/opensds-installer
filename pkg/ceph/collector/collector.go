package collector

import (
	"fmt"

	"github.com/leonwanghui/opensds-installer/api"
	"github.com/leonwanghui/opensds-installer/pkg/ceph/executor"
)

func CollectOsdLocation() (api.OsdLocationMapList, error) {
	metaList, err := executor.GetOsdMetadataList()
	if err != nil {
		return nil, err
	}

	var mapList api.OsdLocationMapList
	for _, meta := range metaList {
		var lMap = make(map[string]string)
		dPath := fmt.Sprint(meta["backend_filestore_partition_path"])

		lMap["id"] = fmt.Sprint(meta["id"])
		lMap["size"] = fmt.Sprint(meta["size"])
		lMap["device_path"] = dPath[:len(dPath)-2]
		lMap["hostname"] = fmt.Sprint(meta["hostname"])
		mapList = append(mapList, lMap)
	}

	if len(mapList) == 0 {
		return nil, fmt.Errorf("Cannot find any osd!")
	}
	return mapList, nil
}
