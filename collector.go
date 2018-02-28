package main

import (
	"fmt"
	"strings"
)

type OsdLocationMapList []map[string]string

func CollectOsdLocation() (OsdLocationMapList, error) {
	metaList, err := getOsdMetadataList()
	if err != nil {
		return nil, err
	}

	var mapList OsdLocationMapList
	for _, meta := range metaList {
		var lMap = make(map[string]string)
		lMap["id"] = fmt.Sprint(meta["id"])
		lMap["device_path"] = strings.TrimRight(fmt.Sprint(meta["backend_filestore_partition_path"]), "p1")
		lMap["hostname"] = fmt.Sprint(meta["hostname"])
		mapList = append(mapList, lMap)
	}

	return mapList, nil
}
