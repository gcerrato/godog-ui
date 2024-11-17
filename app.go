package main

import (
	"context"

	"github.com/gcerrato/godog-backend/pkg/core"
)

// App struct
type App struct {
	ctx  context.Context
	core *core.Core
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	core := core.NewCore()
	core.Init(context.Background())

	a.core = core
}

// Greet returns a greeting for the given name
func (a *App) GetReply(prompt string) string {
	value, _ := a.core.GetReply(a.ctx, prompt)

	return value
}
