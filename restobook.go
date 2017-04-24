package main

import (
	"log"
	"net/http"

	"github.com/akimess/restobook/controllers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://****:*****@ds029565.mlab.com:29565/restobook")

	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	r := mux.NewRouter()
	uc := controllers.NewRestarauntController(getSession())
	r.HandleFunc("/restaraunts", uc.GetRestaraunts).Methods("GET")
	r.HandleFunc("/restaraunts/{id}", uc.GetRestaraunt).Methods("GET")
	r.HandleFunc("/createresto", uc.CreateRestaraunt).Methods("POST")
	r.HandleFunc("/removeresto/{id}", uc.RemoveRestaraunt).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
