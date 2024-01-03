package gui

import (
	"net/http"
	"os"
	"path/filepath"
)

func ServeGUI() {
	// Serve static files
	fs := http.FileServer(http.Dir("gui/static"))
	http.Handle("/gui/static/", http.StripPrefix("/gui/static/", fs))

	// Serve the SPA for any route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("gui/static", "index.html"))
	})

	// Set the PORT from the environment variables or use the default (8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	// Start the server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
