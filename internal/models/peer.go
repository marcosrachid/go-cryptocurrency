package models

type Peer struct {
	Domain string `json:"domain"`
	Port   uint32 `json:"port"`
	Type   string `json:"type"`
}
