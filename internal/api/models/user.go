package models

import "time"

type User struct {
	Id              int
	Email           string
	HashedPassword  []byte
	IsEmailVerified bool
	DateCreated     time.Time
	DateUpdated     time.Time
}
