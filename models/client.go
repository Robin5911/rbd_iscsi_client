package models

type DiscoveryAuth struct {
	MutualPassword string `json:"mutual_password"`
	MutualPasswordEncryptionEnabled bool `json:"mutual_password_encryption_enabled"`
	MutualUsername string `json:"mutual_username"`
	Password string `json:"password"`
	PasswordEncryptionEnabled bool `json:"password_encryption_enabled"`
	Username string `json:"username"`
}

//client auth
type ClientAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Config struct {
	Created string `json:"created"`
	DiscoveryAuth DiscoveryAuth `json:"discovery_auth"`
	Disks interface{} `json:"disks"`
	Epoch int `json:"epoch"`
	Gateways map[string]Gateway `json:"gateways"`
	Targets Target `json:"targets"`
	Updated string `json:"updated"`
	Version int `json:"version"`
}
type Client struct {
	Auth DiscoveryAuth `json:"auth"`
	GroupNmae string `json:"group_nmae"`
	Luns interface{} `json:"luns"`
	Updated string `json:"updated"`
}