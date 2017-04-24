package controllers

import (
	"net/http"

	"github.com/akimess/restobook/models"

	"encoding/json"

	"fmt"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	RestarauntController struct {
		session *mgo.Session
	}
)

func NewRestarauntController(s *mgo.Session) *RestarauntController {
	return &RestarauntController{s}
}

//GetRestaraunts retrieves all restaraunts data
func (rc RestarauntController) GetRestaraunts(w http.ResponseWriter, r *http.Request) {
	var rt []models.Restaurant

	if err := rc.session.DB("restobook").C("restaraunts").Find(nil).All(&rt); err != nil {
		w.WriteHeader(400)
		return
	}

	rj, _ := json.Marshal(rt)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", rj)
}

//GetRestaraunt retrieves individual restaraunt data
func (rc RestarauntController) GetRestaraunt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	rt := models.Restaurant{}

	if err := rc.session.DB("restobook").C("restaraunts").FindId(oid).One(&rt); err != nil {
		w.WriteHeader(404)
		return
	}

	rj, _ := json.Marshal(rt)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", rj)
}

//CreateRestaraunt creates a new restaraunt data
func (rc RestarauntController) CreateRestaraunt(w http.ResponseWriter, r *http.Request) {
	rt := models.Restaurant{}
	json.NewDecoder(r.Body).Decode(&rt)

	rt.ID = bson.NewObjectId()

	rc.session.DB("restobook").C("restaraunts").Insert(rt)

	rj, _ := json.Marshal(rt)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Print(w, "&s", rj)
}

//RemoveRestaraunt removes existing restaraunt data
func (rc RestarauntController) RemoveRestaraunt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := rc.session.DB("restobook").C("restaraunts").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)
}
