package main

import (
	"github.com/akimess/restobook/controllers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://restoadmin:restoran@ds029565.mlab.com:29565")

	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	r := mux.NewRouter()
	uc := controllers.NewRestarauntController()
	r.HandleFunc("/restaraunts", uc.GetRestaraunts).Methods("GET")
	r.HandleFunc("/restaraunts/{id}", uc.GetRestaraunt).Methods("GET")
	r.HandleFunc("/createresto", uc.CreateRestaraunt).Methods("POST")
	r.HandleFunc("/removeresto/{id}", uc.RemoveRestaraunt).Methods("GET")
}
