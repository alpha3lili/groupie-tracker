package server

import (
	"fmt"
	"log"
	"main/controllers"
	"net/http"
)

func InitServer() {
	err := controllers.Init()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/home", controllers.HomeControlers)
	fileserver := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("http://localhost:8000/home")
	err = http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}