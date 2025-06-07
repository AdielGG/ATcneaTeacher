package main

import (
	"app-profesor/backend/server"
	"app-profesor/backend/server/database"
	"app-profesor/backend/streaming"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	database.InitDB()
	server.StartFileServer()
	server.StartWebSocketServer()

	// Create an instance of the app structure
	app := NewApp()
	auth := &server.Auth{}
	str := streaming.NewPDFStream()
	fsr := streaming.NewFileStream()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "app-profesor",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			str,
			fsr,
			auth,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
