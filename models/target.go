package models

type TargetAuth struct {

}
type TargetDisk struct {
	LunId int `json:"lun_id"`
}
type TargetPortal struct {
	GatewayIpList []string `json:"gateway_ip_list"`
	InactivePortalIps []string `json:"inactive_portal_ips"`
	PortalIpAddresses []string `json:"portal_ip_addresses"`
	Tpgs int `json:"tpgs"`
	Updated string `json:"updated"`
}
type Target struct {
	AclEnabled bool `json:"acl_enabled"`
	Auth DiscoveryAuth `json:"auth"`
	Clients map[string]Client `json:"clients"`
	Controls interface{} `json:"controls"`
	Created string `json:"created"`
	Disks map[string]TargetDisk `json:"disks"`
	Groups interface{} `json:"groups"`
	IpList []string `json:"ip_list"`
	Portals map[string]TargetPortal `json:"portals"`
	Updated string `json:"updated"`
	Version int `json:"version"`
}
