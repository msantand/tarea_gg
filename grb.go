package grb

import (
	"time"
	"github.com/Pallinder/go-randomdata"
	"strings"
)

type Author struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Entry struct {
	Key        string    `json:"key"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreateDate time.Time `json:"create_date"`
	Author     Author    `json:"author"`
}

func NewFakeEntry() *Entry {
	return &Entry{
		Key:        strings.ToLower(randomdata.Letters(12)),
		Title:      randomdata.Title(0),
		Content:    randomdata.Letters(100),
		CreateDate: time.Now(),
		Author: Author{
			Username: randomdata.SillyName(),
			Email:    randomdata.Email(),
		},
	}
}
