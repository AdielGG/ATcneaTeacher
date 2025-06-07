package main

import (
	"app-profesor/backend/application"
	app_context "app-profesor/context"

	"context"
)

// App struct
type App struct {
	ctx    context.Context
	dg     *application.DialogImp
	Emiter *application.EventEmitterImp
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.dg = application.NewDialog(ctx)
	a.Emiter = application.NewEventEmitter(ctx)
	app_context.SetContext(a.ctx)
}

// Greet returns a greeting for the given name

func (a *App) OpenFile() string {

	path, err := a.dg.Open("Seleccionar Archivo", "All|*.*")
	if err != nil {
		a.dg.Error(err)
		return ""
	}
	return path
}
