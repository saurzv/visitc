package listing

import (
	"fmt"
)

type Site struct {
	ID         string `json:"siteId"`
	Name       string `json:"name"`
	TotalCount int    `json:"total"`
	Created    string `json:"created"`
	// Analytics Analytics `json:"analytics"`
}

func (s *Site) Human() string {
	return fmt.Sprintf("Name:\t\t%s\nID:\t\t%s\nTotalVisit:\t%d\n", s.Name, s.ID, s.TotalCount)

}

// type Analytics struct {
// 	Created     time.Time `json:"created"`
// 	LastVisited time.Time `json:"lastVisited"`
// 	TotalCount  int       `json:"total"`
// }
