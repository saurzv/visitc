package json

import "time"

type Visit struct {
	IP          string    `json:"ip"`
	LastVisited time.Time `json:"last"`
}

func (v Visit) ID() (jsonField string, value interface{}) {
	jsonField = "ip"
	value = v.IP
	return
}
