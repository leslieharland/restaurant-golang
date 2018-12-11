package main

import (
	"html/template"
	"net/http"
	"restaurant/controllers"
	_ "strconv"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("views/templates/*"))
}

func main() {

	c := controllers.NewController(tpl)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/signup", c.SignUp)
	http.HandleFunc("/login", c.Login)
	http.HandleFunc("/logout", c.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":5001", nil)

}
