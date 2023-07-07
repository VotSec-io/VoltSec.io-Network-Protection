package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Define a struct to hold the data for the template
type PageData struct {
	Title string
}

// HomeHandler handles the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "VoltSec.io Network Protection",
	}

	renderTemplate(w, "index.html", data)
}

// FirewallRulesHandler handles the firewall rules page
func FirewallRulesHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Firewall Rules - VoltSec.io Network Protection",
	}

	renderTemplate(w, "firewall_rules.html", data)
}

// LoggingHandler handles the logging page
func LoggingHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Logging - VoltSec.io Network Protection",
	}

	renderTemplate(w, "logging.html", data)
}

// SettingsHandler handles the settings page
func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data := PageData{
			Title: "Settings - VoltSec.io Network Protection",
		}

		renderTemplate(w, "settings.html", data)
	} else if r.Method == "POST" {
		// Process the form submission
		err := r.ParseForm()
		if err != nil {
			log.Println("Error parsing form:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Get the form values
		emailNotifications := r.Form.Get("email")
		logLevel := r.Form.Get("log-level")

		// TODO: Perform actions based on the form values (e.g., save settings)

		// Redirect to the settings page
		http.Redirect(w, r, "/settings", http.StatusSeeOther)
	}
}

// renderTemplate is a helper function to render HTML templates
func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
