package main

import (
	// "fmt"
	"html/template"
	"net/http"
	"os"
	// "reflect"
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
	}

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
    numbers :=  r.FormValue("numbs") // takes the value from the form
    
v := strings.Fields(``+numbers+``) //split numbers into new lines

 var codes []string
 
 for _, s := range v{
       k := s[1:]

codes = append(codes, k)
 }

d := struct {
        Country string
        Num [] string 
    }{
        Country: countrycode,  
        Num: codes,
}
// fmt.Println(reflect.ValueOf(d.Num).Kind() )
// fmt.Println(d.Num)// print the numbers to cli to see the outcomes

tpl.ExecuteTemplate(w, "walink.html", d)

}