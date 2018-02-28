package api

type OsdLocationMapList []map[string]string

type OsdMetadataList []map[string]interface{}

type CrushMapRuleRequest struct {
	RuleName, BucketName, BucketType, DeviceClass string
}
