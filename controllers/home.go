package controllers

import (
	"html/template"
	"log"
	"main/utils"
	"net/http"
)

var Templates *template.Template
var err error

type DataHome struct {
	Data utils.Artists
}

func Init() error{
	Templates, err = template.New("").ParseGlob("templates/*.html")
	return err
}

func HomeControlers(w http.ResponseWriter, r *http.Request) {
	Data := DataHome{
		Data: utils.GetArtist(),
	}
	err = Templates.ExecuteTemplate(w, "home", Data)
	if err != nil {
		log.Fatal(err)
	}
}
