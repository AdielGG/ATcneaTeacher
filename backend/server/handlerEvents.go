package server

import (
	"app-profesor/backend/server/events"
	"app-profesor/backend/server/variables"
	"encoding/json"
)

func handleEvent(user *variables.User, evt variables.EventRequest) {
	switch evt.Event {
	case "join-class":
		events.JoinClass(user, evt)

	case "get-users":
		events.GetUsers(user)

	case "mensaje":
		// Reenvía el mensaje a todos los usuarios
		msgStr, _ := json.Marshal(map[string]interface{}{
			"event": "mensaje",
			"from":  user.ID,
			"data":  evt.Data,
		})
		broadcast(msgStr)

	case "hand-up":
		events.HandUp(user)

	default:
		user.Send <- []byte(`{"error": "Evento no reconocido"}`)
	}
}
