package Logic

import "GoBasic/StrukturaProjekta/DB"

func (c *Controller) GetUsers() (users []DB.User, err error) {

	//Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetUsers()

}

func (c *Controller) GetUserById(userId int) (user DB.User, err error) {

	//Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetUserById(userId)

}
