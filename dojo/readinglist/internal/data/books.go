package data

import (
	"time"
)

type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // to remove from json response
	Title     string    `json:"title"`
	Published int       `json:"published,omitempty"` // to make a field optional
	Pages     int       `json:"pages,omitempty"`     // to make a field optional
	Genres    []string  `json:"genres,omitempty"`    // to make a field optional
	Rating    float32   `json:"rating,omitempty"`    // to make a field optional
	Version   int32     `json:"-"`                   // to remove from json response
}
