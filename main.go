package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/saintmalik/wame/model"
	"github.com/saintmalik/wame/views"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := chi.NewRouter()
	r.Get("/", Walink)
	r.Post("/lists", Walink)
	http.ListenAndServe(":"+port, r)
}

func Walink(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		templ.Handler(views.Homepage()).ServeHTTP(w, r)
		return
	}
	// if r.URL.Path != "/lists" {
		if r.Method == http.MethodPost {
			countryCode := r.FormValue("cid")
			phoneNum := r.FormValue("numbs")

			// Split phone numbers by whitespace
			phoneNumbers := strings.Fields(phoneNum)
			// Process phone numbers based on the country code
			var wameNumbers []string
			switch countryCode {
			case "https://wa.me/+234":
				wameNumbers = stripPrefix(phoneNumbers, 1)
			case "https://wa.me/+1":
				wameNumbers = stripPrefix(phoneNumbers, 2)
			case "https://wa.me/+44":
				wameNumbers = stripPrefix(phoneNumbers, 3)
			}

			// Generate informational message based on the country code
			var message string
			switch countryCode {
			case "https://wa.me/+234":
				message = "Your Nigeria WaMe Links are ready!!!"
			case "https://wa.me/+1":
				message = "Your USA WaMe Links are ready!!!"
			case "https://wa.me/+44":
				message = "Your UK WaMe Links are ready!!!"
			}

			// Struct to hold data for the template
			data := model.WaMeData{
				CountryMessage: message,
				Country:        countryCode,
				Numbers:        wameNumbers,
			}

			fmt.Println(data)
			templ.Handler(views.Walink(data)).ServeHTTP(w, r)
		}
}

func stripPrefix(numbers []string, prefixLength int) []string {
	var strippedNumbers []string
	for _, number := range numbers {
		if len(number) > prefixLength {
			strippedNumbers = append(strippedNumbers, number[prefixLength:])
		}
	}
	return strippedNumbers
}
