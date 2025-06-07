package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func StartFileServer() {
	// Servirá los archivos desde ./backend/shared como /files/*
	fs := http.FileServer(http.Dir("./backend/shared"))
	http.Handle("/files/", http.StripPrefix("", fs))

	go func() {
		err := http.ListenAndServe(":8888", nil)
		if err != nil {
			log.Fatalf("Error al iniciar el servidor de archivos: %v", err)
		}
	}()
}

func SaveSharedFile(filename string, data []byte) (string, error) {
	path := filepath.Join("backend", "shared", filename)
	return path, os.WriteFile(path, data, 0644)
}
