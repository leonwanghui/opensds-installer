package collector

import (
	"fmt"
	"strings"

	"github.com/leonwanghui/opensds-installer/api"
	"github.com/leonwanghui/opensds-installer/pkg/executor"
)

func CollectOsdLocation() (api.OsdLocationMapList, error) {
	metaList, err := executor.GetOsdMetadataList()
	if err != nil {
		return nil, err
	}

	var mapList api.OsdLocationMapList
	for _, meta := range metaList {
		var lMap = make(map[string]string)
		lMap["id"] = fmt.Sprint(meta["id"])
		lMap["device_path"] = strings.TrimRight(fmt.Sprint(meta["backend_filestore_partition_path"]), "p1")
		lMap["hostname"] = fmt.Sprint(meta["hostname"])
		mapList = append(mapList, lMap)
	}

	return mapList, nil
}
