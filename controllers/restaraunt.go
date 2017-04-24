package controllers

import (
	"net/http"
)

type (
	RestarauntController struct{}
)

func NewRestarauntController() *RestarauntController {
	return &RestarauntController{}
}

//GetRestaraunts retrieves all restaraunts data
func (rc RestarauntController) GetRestaraunts(w http.ResponseWriter, r *http.Request) {

}

//GetRestaraunt retrieves individual restaraunt data
func (rc RestarauntController) GetRestaraunt(w http.ResponseWriter, r *http.Request) {

}

//CreateRestaraunt creates a new restaraunt data
func (rc RestarauntController) CreateRestaraunt(w http.ResponseWriter, r *http.Request) {

}

//RemoveRestaraunt removes existing restaraunt data
func (rc RestarauntController) RemoveRestaraunt(w http.ResponseWriter, r *http.Request) {

}
