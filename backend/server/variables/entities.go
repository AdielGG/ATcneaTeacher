package variables

import (
	"time"

	"github.com/gorilla/websocket"
)

type UsuarioConectado struct {
	ID        int
	Usuario   string
	Nombre    string
	Conectado time.Time
}

type EventMessage struct {
	Event  string `json:"event"`
	Data   string `json:"data"`
	Status int    `json:"status"`
}

type EventRequest struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

type Class struct {
	ClassID int `json:"class-id"`
}

type User struct {
	ID        int
	Conectado bool
	Conn      *websocket.Conn
	Send      chan []byte
}

type ClasInfo struct {
	Name    string `json:"name"`
	Teacher string `json:"teacher"`
	Color   string `json:"color"`
}
