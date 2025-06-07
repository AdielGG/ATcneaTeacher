package server

import (
	"app-profesor/backend/server/variables"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := variables.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error al subir a WebSocket:", err)
		return
	}

	defer conn.Close()

	data, err := json.Marshal(variables.ClasInfo{
		Name:    "testing",
		Teacher: "Ivaldi",
		Color:   "#FF5733FF",
	})

	msg, err := json.Marshal(variables.EventMessage{
		Event:  "class-info",
		Status: 200,
		Data:   string(data),
	})

	fmt.Print(string(msg))
	conn.WriteMessage(websocket.TextMessage, msg)

	ar := conn.LocalAddr()
	fmt.Print(ar)

	// Registrar usuario
	userID, _ := strconv.Atoi(r.Header.Get("user-id")) // Puedes personalizar esto
	user := &variables.User{
		ID:   userID,
		Conn: conn,
		Send: make(chan []byte),
	}

	addUser(user)
	defer removeUser(userID)

	go writer(user)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error al leer:", err)
			removeUser(userID)
			break
		}

		var evt variables.EventRequest

		if err := json.Unmarshal(msg, &evt); err != nil {
			log.Println("Error de JSON:", err)
			continue
		}

		fmt.Print(evt)
		handleEvent(user, evt)
	}
}

func writer(user *variables.User) {
	for msg := range user.Send {
		err := user.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Error escribiendo a", user.ID, ":", err)
			break
		}
	}
}

// === Gestión de usuarios conectados ===

func addUser(user *variables.User) {
	variables.Mutex.Lock()
	defer variables.Mutex.Unlock()
	variables.Users[user.ID] = user
	log.Println("Usuario conectado:", user.ID)
}

func removeUser(userID int) {
	variables.Mutex.Lock()
	defer variables.Mutex.Unlock()

	if user, ok := variables.Users[userID]; ok {
		close(user.Send)
		delete(variables.Users, userID)
		log.Println("Usuario desconectado:", userID)
	}
}

func GetUserList() []int {
	variables.Mutex.Lock()
	defer variables.Mutex.Unlock()

	list := []int{}
	for id := range variables.Users {
		list = append(list, id)
	}
	return list
}

func broadcast(msg []byte) {
	variables.Mutex.Lock()
	defer variables.Mutex.Unlock()

	for _, user := range variables.Users {
		select {
		case user.Send <- msg:
		default:
			log.Println("No se pudo enviar mensaje a", user.ID)
		}
	}
}
