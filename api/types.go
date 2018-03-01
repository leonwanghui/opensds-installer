package api

type Config struct {
	PoolTopology map[string]Location `yaml:"pool,flow"`
}

type Location struct {
	Disks []map[string]string `yaml:"disks"`
}

type OsdLocationMapList []map[string]string

type OsdMetadataList []map[string]interface{}

type CrushMapRuleRequest struct {
	RuleName, BucketName, BucketType, DeviceClass string
}
