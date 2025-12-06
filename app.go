package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (a *App) domReady(ctx context.Context) {
	// Inject environment variables
	envVars := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) == 2 {
			if strings.HasPrefix(pair[0], "VITE_") || strings.HasPrefix(pair[0], "LLM_") {
				envVars[pair[0]] = pair[1]
			}
		}
	}

	if len(envVars) > 0 {
		jsonBytes, err := json.Marshal(envVars)
		if err == nil {
			script := fmt.Sprintf("Object.assign(window, %s);", string(jsonBytes))
			runtime.WindowExecJS(ctx, script)
		}
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
