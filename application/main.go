package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"home-task-tracker/application/child"
	"home-task-tracker/application/goal"
	"home-task-tracker/application/parent"
	"home-task-tracker/application/task"
)

const APPLICATION_PORT = ":8080"

//just leave it here until better times
//err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//return err == nil

func main() {

	r := mux.NewRouter()

	//Parent api
	r.HandleFunc("/api/parent", parent.CreateParent).Methods(http.MethodPost)
	r.HandleFunc("/api/parent/{id}", parent.GetParent).Methods(http.MethodGet)
	r.HandleFunc("/api/parent/{id}", parent.UpdateParent).Methods(http.MethodPut)
	r.HandleFunc("/api/parent/{id}", parent.DeleteParent).Methods(http.MethodDelete)
	r.HandleFunc("/api/parent/{id}/child", child.CreateChild).Methods(http.MethodPost)
	r.HandleFunc("/api/parent/{id}/child", child.GetChildByParentId).Methods(http.MethodGet)
	r.HandleFunc("/api/parent/{id}/task", task.CreateTask).Methods(http.MethodPost)
	r.HandleFunc("/api/parent/{id}/task", task.GetTaskByParentId).Methods(http.MethodGet)
	r.HandleFunc("/api/parent/{id}/goal", goal.CreateGoal).Methods(http.MethodPost)
	r.HandleFunc("/api/parent/{id}/goal", goal.GetGoaldByParentId).Methods(http.MethodGet)

	//Child api
	r.HandleFunc("/api/child/{id}", child.GetChild).Methods(http.MethodGet)
	r.HandleFunc("/api/child/{id}", child.UpdateChild).Methods(http.MethodPut)
	r.HandleFunc("/api/child/{id}/update-points", child.UpdateChildPoints).Methods(http.MethodPut)
	r.HandleFunc("/api/child/{id}", child.DeleteChild).Methods(http.MethodDelete)
	r.HandleFunc("/api/child/{id}/task", task.GetTaskByChildId).Methods(http.MethodGet)

	//Task api
	r.HandleFunc("/api/task/{id}", task.GetTask).Methods(http.MethodGet)
	r.HandleFunc("/api/task/{id}", task.UpdateTask).Methods(http.MethodPut)
	r.HandleFunc("/api/task/{id}", task.DeleteTask).Methods(http.MethodDelete)

	//Goal api
	r.HandleFunc("/api/goal/{id}", goal.GetGoal).Methods(http.MethodGet)
	r.HandleFunc("/api/goal/{id}", goal.UpdateGoal).Methods(http.MethodPut)
	r.HandleFunc("/api/goal/{id}", goal.DeleteGoal).Methods(http.MethodDelete)

	log.Println("Application stared on port", APPLICATION_PORT)
	log.Fatal(http.ListenAndServe(APPLICATION_PORT, r))

}
