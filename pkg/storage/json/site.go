package json

import "time"

type Site struct {
	SiteID    string    `json:"siteID"`
	Name      string    `json:"name"`
	Analytics Analytics `json:"analytics"`
}

type Analytics struct {
	Created     time.Time `json:"created"`
	LastVisited time.Time `json:"lastVisited"`
	TotalCount  int       `json:"totalCount"`
}

func (s Site) ID() (jsonField string, value interface{}) {
	jsonField = "siteID"
	value = s.SiteID
	return
}
