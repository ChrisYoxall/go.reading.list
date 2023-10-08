package data

import (
	"time"
)

type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // omit from JSON
	Title     string    `json:"title"`
	Published int       `json:"published,omitempty"`    // omit from JSON if empty
	Pages     int       `json:"pages,omitempty,string"` // change return data type to string
	Genres    []string  `json:"genres,omitempty"`       // omit from JSON if empty
	Rating    float32   `json:"rating,omitempty"`       // omit from JSON if empty
	Version   int32     `json:"-"`                      // omit from JSON
}
