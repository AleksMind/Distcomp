package model

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrInvalidTitle   = errors.New("title must be between 2 and 64 characters")
	ErrInvalidContent = errors.New("content must be between 4 and 2048 characters")
)

type Topic struct {
	ID        int64      `json:"id" db:"id"`
	CreatorID int64      `json:"creatorId" db:"creator_id"`
	Title     string     `json:"title" db:"title"`
	Content   string     `json:"content" db:"content"`
	Created   time.Time  `json:"created" db:"created"`
	Markers   []string   `json:"markers,omitempty" db:"markers"`
	Modified  *time.Time `json:"modified,omitempty" db:"modified"`
}

func (i *Topic) Validate() error {
	if len(strings.TrimSpace(i.Title)) < 2 || len(i.Title) > 64 {
		return ErrInvalidTitle
	}
	if len(strings.TrimSpace(i.Content)) < 4 || len(i.Content) > 2048 {
		return ErrInvalidContent
	}
	return nil
}
