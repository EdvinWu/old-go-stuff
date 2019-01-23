package goal

import (
	"home-task-tracker/application/core"
)

const (
	CREATE_GOAL           = "INSERT INTO htt.goals(creator_id, name, description, cost, status) VALUES($1,$2,$3,$4,$5) returning id;"
	GET_GOAL_BY_ID        = "SELECT * FROM htt.goals WHERE id=$1;"
	GET_GOAL_BY_PARENT_ID = "SELECT * FROM htt.goals WHERE creator_id=$1;"
	UPDATE_GOAL           = "UPDATE htt.goals SET name=$1, description=$2, cost=$3, status=$4 WHERE id=$5;"
	DELETE_GOAL           = "DELETE FROM htt.goals WHERE id=$1;"
)

type GoalModel struct {
	CreatorId   string `json:"creator_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Cost        string `json:"cost"`
	Status      string `json:"status"`
}

func Create(parentId string, g GoalModel) error {

	db := core.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(CREATE_GOAL)

	defer stmt.Close()

	_, err = stmt.Exec(parentId, g.Name, g.Description, g.Cost, g.Status)

	return err
}

func GetById(id string) (GoalModel, error) {

	db := core.GetConnection()
	defer db.Close()

	goal := GoalModel{}

	var parentId, name, description, cost, status string
	err := db.QueryRow(GET_GOAL_BY_ID, id).Scan(&id, &parentId, &name, &description, &cost, &status)

	goal.CreatorId = parentId
	goal.Name = name
	goal.Description = description
	goal.Cost = cost
	goal.Status = status

	return goal, err
}

func GetByParentId(id string) ([]GoalModel, error) {

	goals := make([]GoalModel, 0)

	db := core.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(GET_GOAL_BY_PARENT_ID)

	if err != nil {
		return goals, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return goals, err
	}

	for rows.Next() {

		var row GoalModel
		var id string
		err := rows.Scan(&id, &row.CreatorId, &row.Name, &row.Description, &row.Cost, &row.Status)

		if err != nil {
			return nil, err
		}

		goals = append(goals, row)
	}

	return goals, err
}

func Update(id string, g GoalModel) error {

	goal, err := GetById(id)

	if len(g.Name) > 0 {
		goal.Name = g.Name
	}
	if len(g.Description) > 0 {
		goal.Description = g.Description
	}
	if len(g.Cost) > 0 {
		goal.Cost = g.Cost
	}
	if len(g.Status) > 0 {
		goal.Status = g.Status
	}

	db := core.GetConnection()
	defer db.Close()

	if err != nil {
		return err
	}

	stmt, err := db.Prepare(UPDATE_GOAL)

	defer stmt.Close()

	_, err = stmt.Exec(goal.Name, goal.Description, goal.Cost, goal.Status, id)

	return err
}

func Delete(id string) error {

	db := core.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(DELETE_GOAL)

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}
