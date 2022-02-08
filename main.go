package main

import (
    // "fmt"
    // "strings"
    "html/template"
    "net/http"
)

var tpl *template.Template

func init(){
    tpl = template.Must(template.ParseGlob("templates/*.html",))
}

func wame(w http.ResponseWriter, r *http.Request){
    tpl.ExecuteTemplate(w, "index.html", nil)
}

func main() {
    http.HandleFunc("/", wame)
    http.HandleFunc("/walink", walink)
    http.ListenAndServe(":8080", nil)
}

func walink(w http.ResponseWriter, r *http.Request){
    if r.Method != "POST" {
        http.Redirect(w, r, "/", http.StatusSeeOther )
        return
}
    countrycode := r.FormValue("cid")
    numbers :=  r.FormValue("numbs")

    d := struct {
        Country string
        Num string
    }{
        Country:  countrycode,
        Num: numbers[1:],
    }

tpl.ExecuteTemplate(w, "walink.html", d)
}