package goal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var g GoalModel

func CreateGoal(w http.ResponseWriter, r *http.Request) {

	mapJsonToGoalModel(r, &g)
	err := Create(mux.Vars(r)["id"], g)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}

}

func GetGoal(w http.ResponseWriter, r *http.Request) {

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

func GetGoaldByParentId(w http.ResponseWriter, r *http.Request) {

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

func UpdateGoal(w http.ResponseWriter, r *http.Request) {

	mapJsonToGoalModel(r, &g)
	err := Update(mux.Vars(r)["id"], g)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

}

func DeleteGoal(w http.ResponseWriter, r *http.Request) {

	err := Delete(mux.Vars(r)["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

}

func mapJsonToGoalModel(r *http.Request, g *GoalModel) {

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&g)

	if err != nil {
		log.Println("Can't parse json to model")
	}

}
