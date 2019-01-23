package task

import (
	"home-task-tracker/application/core"
	"strconv"
)

const (
	CREATE_TASK             = "INSERT INTO htt.tasks(creator_id, assignee_id, name, description, points, status) VALUES($1,$2,$3,$4,$5,$6) returning id;"
	GET_TASK_BY_ID          = "SELECT * FROM htt.tasks WHERE id=$1;"
	GET_TASK_BY_CREATOR_ID  = "SELECT * FROM htt.tasks WHERE creator_id=$1;"
	GET_TASK_BY_ASSIGNEE_ID = "SELECT * FROM htt.tasks WHERE assignee_id=$1;"
	UPDATE_TASK             = "UPDATE htt.tasks SET assignee_id=$1, name=$2, description=$3, points=$4, status=$5 WHERE id=$6;"
	DELETE_TASK             = "DELETE FROM htt.tasks WHERE id=$1;"
)

type TaskModel struct {
	AssigneeId  string `json:"assignee_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Points      string `json:"points"`
	Status      string `json:"status"`
}

func Create(creatorId string, t TaskModel) error {

	db := core.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(CREATE_TASK)

	defer stmt.Close()

	assignedId, _ := strconv.Atoi(t.AssigneeId)
	_, err = stmt.Exec(creatorId, assignedId, t.Name, t.Description, t.Points, t.Status)

	return err
}

func GetById(id string) (TaskModel, error) {

	db := core.GetConnection()
	defer db.Close()

	task := TaskModel{}

	var parentId, assigneeId, name, description, points, status string
	err := db.QueryRow(GET_TASK_BY_ID, id).Scan(&id, &parentId, &assigneeId, &name, &description, &points, &status)

	task.AssigneeId = assigneeId
	task.Name = name
	task.Description = description
	task.Points = points
	task.Status = status

	return task, err
}

func GetByCreatorId(id string) ([]TaskModel, error) {

	tasks := make([]TaskModel, 0)

	db := core.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(GET_TASK_BY_CREATOR_ID)

	if err != nil {
		return tasks, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return tasks, err
	}

	for rows.Next() {

		var row TaskModel
		var id, parentId string
		err := rows.Scan(&id, &parentId, &row.AssigneeId, &row.Name, &row.Description, &row.Points, &row.Status)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, row)
	}

	return tasks, err
}

func GetByAssigneeId(id string) ([]TaskModel, error) {

	tasks := make([]TaskModel, 0)

	db := core.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(GET_TASK_BY_ASSIGNEE_ID)

	if err != nil {
		return tasks, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return tasks, err
	}

	for rows.Next() {

		var row TaskModel
		var id, parentId string
		err := rows.Scan(&id, &parentId, &row.AssigneeId, &row.Name, &row.Description, &row.Points, &row.Status)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, row)
	}

	return tasks, err
}

func Update(id string, t TaskModel) error {

	task, err := GetById(id)

	if len(t.AssigneeId) > 0 {
		task.AssigneeId = t.AssigneeId
	}
	if len(t.Name) > 0 {
		task.Name = t.Name
	}
	if len(t.Description) > 0 {
		task.Description = t.Description
	}
	if len(t.Points) > 0 {
		task.Points = t.Points
	}
	if len(t.Status) > 0 {
		task.Points = t.Points
	}

	db := core.GetConnection()
	defer db.Close()

	if err != nil {
		return err
	}

	stmt, err := db.Prepare(UPDATE_TASK)

	defer stmt.Close()

	_, err = stmt.Exec(task.AssigneeId, task.Name, task.Description, task.Points, task.Status, id)

	return err
}

func Delete(id string) error {

	db := core.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(DELETE_TASK)

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}
