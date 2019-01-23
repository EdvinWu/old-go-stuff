package parent

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"home-task-tracker/application/core"
	"fmt"
)

func CreateParent(w http.ResponseWriter, r *http.Request) {
	var p ParentModel

	mapJsonToParentModel(r, &p)
	err := Create(p)

	if err != nil {
		panic(err)
	}

	redis := core.WriteToRedis(p.Login, p, core.PARENT_CACHE)
	fmt.Printf("redis: %v ", redis)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}

}

func GetParent(w http.ResponseWriter, r *http.Request) {

	prnt, err := GetById(mux.Vars(r)["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(prnt)
	}

}

func UpdateParent(w http.ResponseWriter, r *http.Request) {
	var p ParentModel

	mapJsonToParentModel(r, &p)
	err := Update(mux.Vars(r)["id"], p)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

}

func DeleteParent(w http.ResponseWriter, r *http.Request) {

	err := Delete(mux.Vars(r)["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

}

func mapJsonToParentModel(r *http.Request, p *ParentModel) {

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&p)

	if err != nil {
		log.Println("Can't parse json to model")
	}

}
