package DB

type DB interface {
	Init() (err error)

	GetUsers() (users []DataStructures.User, err error)
	GetUserById(id int) (user DataStructures.User, err error)
}
