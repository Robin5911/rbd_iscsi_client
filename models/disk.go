package models

type DiskInfo struct {
	AllocationHost string `json:"allocation_host"`
	Backstore string `json:"backstore"`
	BackstoreObjectName string `json:"backstore_object_name"`
	Controls interface{} `json:"controls"`
	Created string `json:"created"`
	Image string `json:"image"`
	Owner string `json:"owner"`
	Pool string `json:"pool"`
	PoolId int `json:"pool_id"`
	Updated string `json:"updated"`
	Wwn string `json:"wwn"`
 }

