package child

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var c ChildModel

func CreateChild(w http.ResponseWriter, r *http.Request) {

	mapJsonToChildModel(r, &c)
	err := Create(mux.Vars(r)["id"], c)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}

}

func GetChild(w http.ResponseWriter, r *http.Request) {

	chld, err := GetById(mux.Vars(r)["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(chld)
	}

}

func GetChildByParentId(w http.ResponseWriter, r *http.Request) {

	chlds, err := GetByParentId(mux.Vars(r)["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(chlds)
	}

}

func UpdateChild(w http.ResponseWriter, r *http.Request) {

	mapJsonToChildModel(r, &c)
	err := Update(mux.Vars(r)["id"], c)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

}

func UpdateChildPoints(w http.ResponseWriter, r *http.Request) {

	mapJsonToChildModel(r, &c)
	err := UpdatePoints(mux.Vars(r)["id"], c)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

}

func DeleteChild(w http.ResponseWriter, r *http.Request) {

	err := Delete(mux.Vars(r)["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

}

func mapJsonToChildModel(r *http.Request, c *ChildModel) {

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&c)

	if err != nil {
		log.Println("Can't parse json to model")
	}

}
