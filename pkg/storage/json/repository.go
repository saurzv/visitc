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
	siteDir  = "/data/sites"
	visitDir = "/data/visits"
)

type Storage struct {
	db      *simdb.Driver
	visitDB *simdb.Driver
}

func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, file, _, _ := runtime.Caller(0)
	p := path.Dir(file)

	s.db, err = simdb.New(p + siteDir)
	if err != nil {
		return nil, err
	}

	s.visitDB, err = simdb.New(p + visitDir)
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
	siteFound.Created = site.Analytics.Created.Format("Jan 02 '06")

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

// IP

func (s *Storage) AddVisit(addr string) error {
	var visit Visit
	err := s.visitDB.Open(Visit{}).Where("ip", "=", addr).First().AsEntity(&visit)
	if err == nil {
		return err
	}

	newVist := Visit{
		IP:          addr,
		LastVisited: time.Now(),
	}
	err = s.visitDB.Insert(newVist)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateLastVisit(addr string) error {
	var visit Visit
	err := s.visitDB.Open(Visit{}).Where("ip", "=", addr).First().AsEntity(&visit)
	if err != nil {
		return err
	}

	visit.LastVisited = time.Now()
	err = s.visitDB.Update(visit)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetLastVisit(addr string) (time.Time, error) {
	var visit Visit
	err := s.visitDB.Open(Visit{}).Where("ip", "=", addr).First().AsEntity(&visit)
	if err != nil {
		return time.Now(), err
	}

	return visit.LastVisited, nil
}
