package variables

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Permitir conexiones desde cualquier origen
		},
	}

	Users = make(map[int]*User) // Mapa de usuarios conectados
	Mutex = sync.Mutex{}

	UsersConected []int
)
