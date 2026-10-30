package models

import "time"

type Contractor struct {
	CreatedAt time.Time
	Mobile    int64
	ID        int64
	Name      string
	Address   string
	Email     string
	Password  string
}
