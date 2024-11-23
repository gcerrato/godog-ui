package main

import (
	"context"
	"time"

	"github.com/gcerrato/godog-backend/pkg/core"
	"github.com/gcerrato/godog-backend/pkg/fs"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	currentDirectory string
	ctx              context.Context
	core             *core.Core
	fs               *fs.FileScanner
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
	core.Init(ctx)

	a.core = core
	a.currentDirectory = "."
	a.fs = a.InitScanner(a.currentDirectory)
}

// Greet returns a greeting for the given name
func (a *App) GetReply(prompt string) string {
	value, _ := a.core.GetReply(a.ctx, prompt)

	return value
}

func (a *App) InitScanner(dirPath string) *fs.FileScanner {
	scannerHelper := fs.NewScannerHelper()
	opts := fs.FileScannerOpts{Interval: 10 * time.Millisecond, FilePath: dirPath, Rescursive: false}
	filescanner := fs.GetFileScannerInstance(opts)
	filescanner.AddObserver(scannerHelper)

	go handleFiles(scannerHelper, a.core)
	return filescanner
}

func handleFiles(sh *fs.ScannerHelper, core *core.Core) {
	for {
		files := <-sh.Notifications

		for _, file := range files {
			core.ProcessFile(context.Background(), file, sh)
		}
	}
}

func (a *App) SetDirectory() string {
	path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		// do something
	}

	if a.currentDirectory != path {
		a.currentDirectory = path
	}

	return a.fs.UpdatePath(a.currentDirectory)
}

func (a *App) UploadDocuments() {
	a.fs.Scan()
}

func (a *App) ConvertPDFToImages() {
	// a.core.ConvertPDFToImages(a.fs.)
}
