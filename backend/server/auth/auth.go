package auth

import (
	"app-profesor/backend/models"
	"app-profesor/backend/server/database"
	"database/sql"
	"errors"
	"log"
)

type Auth struct{}

var a Auth

func (a *Auth) Register(user models.Usuario) (models.Usuario, error) {
	row := database.DB.QueryRow("SELECT id From usuarios order by id Desc limit 1")

	id := 0

	err := row.Scan(&id)

	id++
	stmt, err := database.DB.Prepare(`INSERT INTO usuarios(id,nombre, apellidos, usuario, password) VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		return models.Usuario{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, user.Nombre, user.Apellidos, user.Usuario, user.Password)
	if err != nil {
		return models.Usuario{}, err
	}

	return user, nil
}

func (a *Auth) Login(usuario string, password string) (models.Usuario, error) {
	row := database.DB.QueryRow("SELECT id, nombre, apellidos,  usuario FROM usuarios WHERE usuario=$1 AND password=$2", usuario, password)

	var user models.Usuario
	err := row.Scan(&user.ID, &user.Nombre, &user.Apellidos, &user.Usuario)
	log.Println(user)

	if err == sql.ErrNoRows {
		return models.Usuario{}, errors.New("usuario o contraseña incorrectos")
	} else if err != nil {
		return models.Usuario{}, err
	}
	return user, nil
}
