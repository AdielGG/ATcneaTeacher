package events

import (
	"app-profesor/backend/server/variables"
	"encoding/json"
	"fmt"
	"log"
)

func SetScreen(evt variables.EventRequest) {

	data, err := json.Marshal(evt)
	if err != nil {
		log.Panicln(err.Error())
	}

	fmt.Println(string(data))
}
