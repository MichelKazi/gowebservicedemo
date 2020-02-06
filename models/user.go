package models

// The User model has an ID, Firstname and LastName
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)
