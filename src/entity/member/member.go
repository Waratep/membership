package member

import "time"

type Member struct {
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
