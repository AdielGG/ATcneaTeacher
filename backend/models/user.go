package models

type Usuario struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellidos string `json:"apellidos"`
	Email     string `json:"email"`
	Usuario   string `json:"usuario"`
	Password  string `json:"password"`
}
