package main

type Rule struct {
	Name         string
	Ruleset      int
	MinSize      int
	MaxSize      int
	BucketName   string
	DeviceClass  string
	BucketNumber int
	BucketType   string
}

type Device struct {
	ID            int
	Name          string
	DeviceClasses []string
}

type BucketType struct {
	ID   int
	Name string
}

type Bucket struct {
	Type        string
	Name        string
	ID          int
	BucketItems struct {
		Name   string
		Weight string
	}
}
