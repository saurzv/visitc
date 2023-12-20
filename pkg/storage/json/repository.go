package json

import (
	"log"
	"path"
	"runtime"
	"time"

	"github.com/saurzv/visitc/pkg/listing"
	"github.com/saurzv/visitc/pkg/storage"
	"github.com/saurzv/visitc/pkg/updating"
	"github.com/sonyarouje/simdb"
)

const (
	dir = "/data/"
)

type Storage struct {
	db *simdb.Driver
}

func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, file, _, _ := runtime.Caller(0)
	p := path.Dir(file)

	s.db, err = simdb.New(p + dir)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Storage) AddSite(newSite updating.Site) error {
	id, err := storage.GetId(newSite.Name)
	if err != nil {
		return err
	}
	newS := Site{
		SiteID: id,
		Name:   newSite.Name,
		Analytics: Analytics{
			Created: time.Now(),
		},
	}

	if err := s.db.Insert(newS); err != nil {
		return err
	}
	return nil
}

func (s *Storage) RemoveSite(site listing.Site) error {
	siteToRemove := Site{
		SiteID: site.ID,
	}
	if err := s.db.Delete(siteToRemove); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetSite(id string) (listing.Site, error) {
	var site Site
	var siteFound listing.Site

	if err := s.db.Open(Site{}).Where("siteID", "=", id).First().AsEntity(&site); err != nil {
		return siteFound, err
	}

	siteFound.ID = site.SiteID
	siteFound.Name = site.Name
	siteFound.TotalCount = site.Analytics.TotalCount
	siteFound.Created = site.Analytics.Created.Format("Jan 02, '06")

	return siteFound, nil
}

func (s *Storage) GetAllSites() []listing.Site {
	list := []listing.Site{}
	err := s.db.Open(Site{}).Get().AsEntity(&list)
	if err != nil {
		return list
	}
	return list
}

func (s *Storage) IncreaseCount(id string) error {
	var site Site

	err := s.db.Open(Site{}).Where("siteID", "=", id).First().AsEntity(&site)
	if err != nil {
		log.Fatal(err)
		// return err
	}

	site.Analytics.LastVisited = time.Now()
	// TO-DO: Write better algorithm for analytics
	site.Analytics.TotalCount += 1
	err = s.db.Update(site)
	if err != nil {
		return err
	}

	return nil
}
