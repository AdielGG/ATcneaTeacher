package events

import (
	"app-profesor/backend/server/variables"
	"app-profesor/context"

	"encoding/json"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Data struct {
	Time int `json:"time"`
}

func HandUp(user *variables.User) {

	data, _ := json.Marshal(Data{
		Time: 10,
	})

	res, _ := json.Marshal(variables.EventMessage{
		Event:  "hand-up",
		Status: 200,
		Data:   string(data),
	})
	user.Send <- res
	runtime.EventsEmit(context.GetContext(), "hand-up", user.ID)

}
