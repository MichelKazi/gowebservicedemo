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

// GetUsers returns Users
func GetUsers() []*User {
	return users
}

// AddUser adds a user to the users slice
// returns the user that was added
func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
