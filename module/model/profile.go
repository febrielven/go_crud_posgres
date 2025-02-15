package model

import (
	"time"
)

// Profile struct
type Profile struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Profiles type Profile list
type Profiles []Profile

//NewProfile Profile's Constructor
func NewProfile() *Profile {
	return &Profile{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
