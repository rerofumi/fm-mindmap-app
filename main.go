package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Load .env
	// 1. Try side-by-side with executable (Production)
	ex, err := os.Executable()
	if err == nil {
		envPath := filepath.Join(filepath.Dir(ex), ".env")
		_ = godotenv.Load(envPath)
	}
	// 2. Try current directory (Dev / Fallback)
	_ = godotenv.Load()

	// Prepare env vars for injection
	envVars := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) == 2 {
			// Pass both VITE_ prefixed and LLM_ prefixed variables
			if strings.HasPrefix(pair[0], "VITE_") || strings.HasPrefix(pair[0], "LLM_") {
				envVars[pair[0]] = pair[1]
			}
		}
	}

	// Create an instance of the app structure
	app := NewApp()

	// Create AssetServer options
	assetServerOptions := &assetserver.Options{
		Assets: assets,
	}

	// Only add Middleware in production (when assets are embedded)
	// We check !isDev to ensure we don't mess with Vite dev server proxy in 'wails dev' mode.
	if !isDev && len(envVars) > 0 {
		assetServerOptions.Middleware = func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Intercept index.html only when loading from embedded assets
				if r.URL.Path == "/" || r.URL.Path == "/index.html" {
					// Try to read from embedded assets
					fileData, err := assets.ReadFile("frontend/dist/index.html")
					if err != nil {
						// Should not happen in production if built correctly
						next.ServeHTTP(w, r)
						return
					}

					// Generate script
					var scriptBuilder strings.Builder
					scriptBuilder.WriteString("<script>")
					for k, v := range envVars {
						b, _ := json.Marshal(v)
						scriptBuilder.WriteString(fmt.Sprintf("window.%s = %s;", k, string(b)))
					}
					scriptBuilder.WriteString("</script>")

					content := string(fileData)
					// Inject script at the beginning of head
					content = strings.Replace(content, "<head>", "<head>"+scriptBuilder.String(), 1)

					w.Header().Set("Content-Type", "text/html")
					w.Write([]byte(content))
					return
				}
				next.ServeHTTP(w, r)
			})
		}
	}
	

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "fm-mindmap",
		Width:  1024,
		Height: 768,
		AssetServer: assetServerOptions,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnDomReady:       app.domReady,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
