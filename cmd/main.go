package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Define the routes for UI handlers
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/firewall/rules", firewallRulesHandler)
	router.HandleFunc("/logging", loggingHandler)
	router.HandleFunc("/settings", settingsHandler)

	// Serve static files from the "static" directory
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))

	// Start the HTTP server
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// UI handler for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your home page logic here
	// Example: Render the "index.html" template
	http.ServeFile(w, r, "./ui/templates/index.html")
}

// UI handler for firewall rules
func firewallRulesHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your firewall rules page logic here
	// Example: Render the "firewall_rules.html" template
	http.ServeFile(w, r, "./ui/templates/firewall_rules.html")
}

// UI handler for logging
func loggingHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your logging page logic here
	// Example: Render the "logging.html" template
	http.ServeFile(w, r, "./ui/templates/logging.html")
}

// UI handler for settings
func settingsHandler(w http.ResponseWriter, r *http.Request) {
	// Implement your settings page logic here
	// Example: Render the "settings.html" template
	http.ServeFile(w, r, "./ui/templates/settings.html")
}
