package events

import (
	"app-profesor/backend/server/variables"
	"encoding/json"
)

type Class_id struct {
	ClassID int `json:"class-id"`
}

func JoinClass(user *variables.User, eve variables.EventRequest) {

	event := "joined-class"

	class := Class_id{
		ClassID: 21,
	}
	data, err := json.Marshal(class)
	if err != nil {
		res, _ := json.Marshal(variables.EventMessage{
			Event:  event,
			Status: 400,
			Data:   err.Error(),
		})

		variables.UsersConected = append(variables.UsersConected, user.ID)
		user.Send <- res
		return
	}

	res, err := json.Marshal(variables.EventMessage{
		Event:  "joined-class",
		Status: 200,
		Data:   string(data),
	})
	if err != nil {

		res, _ := json.Marshal(variables.EventMessage{
			Event:  event,
			Status: 500,
			Data:   err.Error(),
		})
		user.Send <- res
		return
	}

	user.Send <- res
}
