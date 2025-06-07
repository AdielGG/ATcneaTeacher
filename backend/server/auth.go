package server

import (
	"app-profesor/backend/models"
	"app-profesor/backend/server/database"
	"database/sql"
	"errors"
)

type Auth struct{}

func (a *Auth) Register(user models.Usuario) (models.Usuario, error) {
	stmt, err := database.DB.Prepare(`INSERT INTO usuarios(nombre, apellidos, email, usuario, password) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return models.Usuario{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Nombre, user.Apellidos, user.Email, user.Usuario, user.Password)
	if err != nil {
		return models.Usuario{}, err
	}

	id, _ := result.LastInsertId()
	user.ID = int(id)
	user.Password = "" // no retornar contraseña
	return user, nil
}

func (a *Auth) Login(usuario string, password string) (models.Usuario, error) {
	row := database.DB.QueryRow("SELECT id, nombre, apellidos, email, usuario FROM usuarios WHERE usuario = ? AND password = ?", usuario, password)

	var user models.Usuario
	err := row.Scan(&user.ID, &user.Nombre, &user.Apellidos, &user.Email, &user.Usuario)
	if err == sql.ErrNoRows {
		return models.Usuario{}, errors.New("usuario o contraseña incorrectos")
	} else if err != nil {
		return models.Usuario{}, err
	}
	return user, nil
}
