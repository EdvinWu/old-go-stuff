package task

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var t TaskModel

func CreateTask(w http.ResponseWriter, r *http.Request) {

	mapJsonToTaskModel(r, &t)
	err := Create(mux.Vars(r)["id"], t)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}

}

func GetTask(w http.ResponseWriter, r *http.Request) {

	task, err := GetById(mux.Vars(r)["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
	}

}

func GetTaskByParentId(w http.ResponseWriter, r *http.Request) {

	tasks, err := GetByCreatorId(mux.Vars(r)["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tasks)
	}

}

func GetTaskByChildId(w http.ResponseWriter, r *http.Request) {

	tasks, err := GetByAssigneeId(mux.Vars(r)["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tasks)
	}

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	mapJsonToTaskModel(r, &t)
	err := Update(mux.Vars(r)["id"], t)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	err := Delete(mux.Vars(r)["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

}

func mapJsonToTaskModel(r *http.Request, t *TaskModel) {

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&t)

	if err != nil {
		log.Println("Can't parse json to model")
	}

}
