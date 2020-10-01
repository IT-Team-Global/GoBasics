package MariaDB

import (
	"GoBasic/StrukturaProjekta/DB"
	"errors"
	"fmt"
)

//Funkcija iz interface-a
func (db *MariaDB) GetUsers() (users []DB.User, err error) {

	//Naredimo query na bazo
	//Za stavke od katerih ne pričakujemo odgovora (UPDATE, INSERT) uporabimo namesto "Query" "Exec"
	rows, err := db.database.Query("SELECT user_id, username, email from user ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Z defer določimo funkcijo, ki se bo poklicala kadar se funkcija nad njo zaključi.
	//V tem primeru bo to klicano po return

	//Pri SQL-u moremo VEDNO zapreti vrstice.
	//V nasprotnem primeru pride do memory leak-a, vrstice ostanejo odprte, povezave ostaneko odprte
	//in čez čas zmanjka povezav in se dostop do baze ustavi in edina rešitev je potem še reboot
	defer func() {
		err2 := rows.Close()
		if err2 != nil {
			fmt.Println(err2.Error())
		}
	}()

	//Ustvarimo objekt User in neomejen array tipa User
	var user DB.User
	users = make([]DB.User, 0)

	//Loop čez vse vrstice
	for rows.Next() {

		//Preberemo podatke vrstice v objekt
		err = rows.Scan(&user.Id, &user.Username, &user.Email)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//Dodamo objekt user na konec arraya users
		users = append(users, user)

	}

	//Vrnemo podatke. Ne rabimo napisat katero spremenjljivko vrnemo saj je ta definirana in inicializirana na vrhu funkcije
	return

}

func (db *MariaDB) GetUserById(userId int) (user DB.User, err error) {

	rows, err := db.database.Query("SELECT user_id, username, email from user  WHERE user_id = ? LIMIT 1", userId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {
		err2 := rows.Close()
		if err2 != nil {
			fmt.Println(err2.Error())
		}
	}()

	//Prestavimo se na naslednjo vrstico (na prvo vrnjeno vrstico), če je ni pomeni da je odgovor prazen in ne obstaja
	if !rows.Next() {
		err = errors.New("user does not exist")
		return
	}

	err = rows.Scan(&user.Id, &user.Username, &user.Email)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return

}
