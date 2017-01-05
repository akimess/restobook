package restobook

import(
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Restaurant struct{
	ID string `json:"id, omitempty"`
	Name string `json:"name, omitempty"`
	Address *Address `json:"address, omitempty"`
}

type Address struct {
	City string `json:"city, omitempty"`
	Street string `json:"street, omitempty"`
}

var restaurants []Restaurant

func GetRestaurantEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for _, item := range restaurants {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Restaurant{})
}

func GetRestarauntsEndpoint(w http.ResponseWriter, req *http.Request){
	json.NewEncoder(w).Encode(restaurants)
}

func CreateRestarauntEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	var restaurant Restaurant
	_ = json.NewDecoder(req.Body).Decode(&restaurant)
	restaurant.ID = params["id"]
	restaurants = append(restaurants, restaurant)
	json.NewEncoder(w).Encode(restaurants)
}
