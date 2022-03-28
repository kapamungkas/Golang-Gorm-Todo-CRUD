package entity

import "time"

type Todo struct {
	ID          int
	Todo        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
