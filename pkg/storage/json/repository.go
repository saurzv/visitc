package json

import (
	"encoding/json"
	"path"
	"runtime"
	"time"

	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/saurzv/visitc/pkg/adding"
	"github.com/saurzv/visitc/pkg/listing"
	"github.com/saurzv/visitc/pkg/storage"
)

const (
	dir            = "/data/"
	CollectionSite = "sites"
)

type Storage struct {
	db *scribble.Driver
}

func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, file, _, _ := runtime.Caller(0)
	p := path.Dir(file)

	s.db, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Storage) AddSite(newSite adding.Site) error {
	id, err := storage.GetId(newSite.Name)
	if err != nil {
		return err
	}
	newS := Site{
		ID:      id,
		Name:    newSite.Name,
		Created: time.Now(),
	}

	if err := s.db.Write(CollectionSite, newS.ID, &newS); err != nil {
		return err
	}
	return nil
}

func (s *Storage) RemoveSite(id string) error {
	if err := s.db.Delete(CollectionSite, id); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetSite(id string) (listing.Site, error) {
	var site Site
	var gotSite listing.Site

	if err := s.db.Read(CollectionSite, id, &site); err != nil {
		return gotSite, err
	}

	gotSite.ID = site.ID
	gotSite.Name = site.Name
	gotSite.Created = site.Created

	return gotSite, nil
}

func (s *Storage) GetAllSites() []listing.Site {
	list := []listing.Site{}
	records, err := s.db.ReadAll(CollectionSite)
	if err != nil {
		return list
	}

	for _, r := range records {
		var site Site
		var gotSite listing.Site

		if err := json.Unmarshal([]byte(r), &site); err != nil {
			return list
		}

		gotSite.ID = site.ID
		gotSite.Name = site.Name
		gotSite.Created = site.Created

		list = append(list, gotSite)
	}
	return list
}
