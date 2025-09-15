package swagger

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed swagger-ui
var swaggerUI embed.FS

//go:embed api/auth.swagger.json
var authSwaggerJSON embed.FS

//go:embed api/chat.swagger.json
var chatSwaggerJSON embed.FS

//go:embed api/social.swagger.json
var socialSwaggerJSON embed.FS

//go:embed api/users.swagger.json
var usersSwaggerJSON embed.FS

func SetupSwagger(mux *http.ServeMux) {
	// Раздача swagger.json
	mux.HandleFunc("/swagger/auth.json", func(w http.ResponseWriter, r *http.Request) {
		data, err := authSwaggerJSON.ReadFile("api/auth.swagger.json")
		if err != nil {
			http.Error(w, "Failed to load swagger.json", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})

	mux.HandleFunc("/swagger/chat.json", func(w http.ResponseWriter, r *http.Request) {
		data, err := chatSwaggerJSON.ReadFile("api/chat.swagger.json")
		if err != nil {
			http.Error(w, "Failed to load swagger.json", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})

	mux.HandleFunc("/swagger/social.json", func(w http.ResponseWriter, r *http.Request) {
		data, err := socialSwaggerJSON.ReadFile("api/social.swagger.json")
		if err != nil {
			http.Error(w, "Failed to load swagger.json", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})

	mux.HandleFunc("/swagger/users.json", func(w http.ResponseWriter, r *http.Request) {
		data, err := usersSwaggerJSON.ReadFile("api/users.swagger.json")
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
