package main

import (
	"log"
	"net/http"

	"github.com/akimess/restobook/controllers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://*****:*****@ds029565.mlab.com:29565/restobook")

	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	r := mux.NewRouter()
	rc := controllers.NewRestarauntController(getSession())
	r.HandleFunc("/restaraunts", rc.GetRestaraunts).Methods("GET")
	r.HandleFunc("/restaraunts/{id}", rc.GetRestaraunt).Methods("GET")
	r.HandleFunc("/createresto", rc.CreateRestaraunt).Methods("POST")
	r.HandleFunc("/removeresto/{id}", rc.RemoveRestaraunt).Methods("DELETE")
	r.HandleFunc("/updateresto/{id}", rc.UpdateRestaraunt).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
