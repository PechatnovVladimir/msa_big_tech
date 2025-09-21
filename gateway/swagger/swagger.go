package swagger

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed swagger-ui
var swaggerUI embed.FS

//go:embed api/auth.swagger.json
var swaggerJSON embed.FS

func SetupSwagger(mux *http.ServeMux) {
	// Раздача swagger.json
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		data, err := swaggerJSON.ReadFile("api/auth.swagger.json")
		if err != nil {
			http.Error(w, "Failed to load swagger.json", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})

	// Раздача Swagger UI
	sub, _ := fs.Sub(swaggerUI, "swagger-ui")
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.FS(sub))))
}
