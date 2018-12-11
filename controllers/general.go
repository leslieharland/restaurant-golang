package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"restaurant/models"
	"restaurant/session"
)

type viewmodel struct{
	Restaurants []models.Restaurant
	User models.User
}
type Controller struct {
	tpl *template.Template
}

func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

func (c Controller) Index(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)

	data := viewmodel{
		Restaurants: nil,
		User: u,
	}
	c.tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func (c Controller) Restaurants(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	jsonFile, err := os.Open("json/generated_restaurants.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Successfully Opened generated_restaurants.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array

	var restaurants []models.Restaurant

	json.Unmarshal(byteValue, &restaurants)
	data := viewmodel{
		Restaurants: restaurants,
		User: u,
	}


	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example


	c.tpl.ExecuteTemplate(w, "index.gohtml", data)
}
