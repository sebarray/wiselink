package model

import (
	"time"
)

type Event struct {
	Id               string    `json:"id"`
	DescriptionShort string    `json:"descriptionshort"`
	Description      string    `json:"description"`
	State            string    `json:"state"`
	Organizer        string    `json:"organizer"`
	Place            string    `json:"place"`
	Date             time.Time `json:"date"`
	Title            string    `json:"title"`
}
