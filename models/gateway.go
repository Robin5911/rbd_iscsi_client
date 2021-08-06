package models

type Gateway struct {
	ActiveLuns int `json:"active_luns"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}
