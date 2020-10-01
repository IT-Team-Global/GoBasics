package DB

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type DB interface {
	Init() (err error)

	GetUsers() (users []User, err error)
	GetUserById(id int) (user User, err error)
}
