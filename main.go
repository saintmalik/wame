package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var tpl *template.Template

func init(){
    tpl = template.Must(template.ParseGlob("templates/*.html",))
}

func wame(w http.ResponseWriter, r *http.Request){
    tpl.ExecuteTemplate(w, "index.html", nil)
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
	  port = "8080"
	} // this was for simple deployment on heroku

    http.HandleFunc("/", wame)
    http.HandleFunc("/walink", walink)
    http.ListenAndServe(":"+port, nil)
}

func walink(w http.ResponseWriter, r *http.Request){
    if r.Method != "POST" {
        http.Redirect(w, r, "/", http.StatusSeeOther ) 
        return
}
    countrycode := r.FormValue("cid")
    phonenum :=  r.FormValue("numbs") // takes the value from the form
    
v := strings.Fields(``+phonenum+``) //split numbers into new lines

var WameNum []string //declared a global variable

switch countrycode {
    case "https://wa.me/+234":
        for _, number := range v {
           l := number[1:]
           WameNum = append(WameNum, l)
        }
    case "https://wa.me/+1":
        for _, number := range v {
            l := number[2:]
            WameNum = append(WameNum, l)
     }
    case "https://wa.me/+44":
        for _, number := range v {
            l := number[3:]
            WameNum = append(WameNum, l)
        }
} // a switch case to check the country code to decide the number format stripping, so it runs a range over the numbers and strips the required things off based on the rules

var  inform string  //declared a global variable

switch countrycode {
case "https://wa.me/+234":
    inform = fmt.Sprint("Your Nigerian WaMe Links are ready!!!")
case "https://wa.me/+1":
    inform = fmt.Sprint("Your USA WaMe Links are ready!!!")
case "https://wa.me/+44":
    inform = fmt.Sprint("Your UK WaMe Links are ready!!!")
} //this switch case is to check the country code to decide the message to be display for the user based on the country they have picked

d := struct {
        CountryMessage string
        Country string
        Number [] string
    }{
        CountryMessage:  inform,
        Country: countrycode,  
        Number: WameNum,
}
// a struct to hold the data and pass them to the template to be deisplayed

tpl.ExecuteTemplate(w, "walink.html", d)

}