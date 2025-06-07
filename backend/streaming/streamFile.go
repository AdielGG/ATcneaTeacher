package streaming

import (
	"app-profesor/backend/server"
	"encoding/json"
	"sync"

	"github.com/google/uuid"
)

type FileMessage struct {
	Filename string
	Data     []byte
}

type FileStream struct {
	mu      sync.Mutex
	clients map[string]func(FileMessage)
}

func NewFileStream() *FileStream {
	return &FileStream{
		clients: make(map[string]func(FileMessage)),
	}
}

func (fs *FileStream) SendFile(targetIDs []string, typ string, data []byte) {

	url, err := server.SaveSharedFile(typ, data)
	if err != nil {
		return
	}

	msg := map[string]interface{}{
		"type": "file",
		"url":  url,
	}

	res, _ := json.Marshal(msg)

	server.SendToAll(res)
}

func (fs *FileStream) RegisterClient(cb func(FileMessage)) string {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	id := uuid.New().String()
	fs.clients[id] = cb
	return id
}

func (fs *FileStream) UnregisterClient(id string) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	delete(fs.clients, id)
}
