package listing

type Site struct {
	ID         string `json:"siteId"`
	Name       string `json:"name"`
	TotalCount int    `json:"total"`
	// Analytics Analytics `json:"analytics"`
}

// type Analytics struct {
// 	Created     time.Time `json:"created"`
// 	LastVisited time.Time `json:"lastVisited"`
// 	TotalCount  int       `json:"total"`
// }
