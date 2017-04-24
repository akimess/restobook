package controllers

import (
	"net/http"

	"github.com/akimess/restobook/models"

	"encoding/json"

	"fmt"

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

}

//GetRestaraunt retrieves individual restaraunt data
func (rc RestarauntController) GetRestaraunt(w http.ResponseWriter, r *http.Request) {

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

}
