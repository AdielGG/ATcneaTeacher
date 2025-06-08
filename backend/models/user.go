package models

type Usuario struct {
	ID        int    `json:"id"`
	Nombre    string `json:"name"`
	Apellidos string `json:"last_name"`
	Usuario   string `json:"username"`
	Password  string `json:"password"`
	Img       string `json:"porfile_image"`
}
