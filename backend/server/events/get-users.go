package events

import (
	"app-profesor/backend/server/variables"
	"encoding/json"
)

func GetUsers(user *variables.User) {
	var userList []int
	for id := range variables.Users {
		userList = append(userList, id)
	}

	response, _ := json.Marshal(map[string]interface{}{
		"event":              "users",
		"usuariosConectados": userList,
	})
	user.Send <- response
}
