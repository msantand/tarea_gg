package database

import (
	"github.com/msantand/tarea_gg"
	"errors"
)

type Database struct {
	entries map[string]*grb.Entry
}

func New() *Database {
	return &Database{
		entries: make(map[string]*grb.Entry),
	}
}

func (db *Database) AddEntry(entry *grb.Entry) {
	db.entries[entry.Key] = entry
}

func (db *Database) EntryList() []*grb.Entry {
	entries := make([]*grb.Entry, len(db.entries))

	var i = 0
	for _, entry := range db.entries {
		entries[i] = entry
		i += 1
	}

	return entries
}

func (db *Database) EntryGet(key string) (*grb.Entry, error) {
	value, ok := db.entries[key]
	if !ok {
		return nil, errors.New("entry not found")
	}
	return value, nil
}
