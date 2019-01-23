package child

import (
	"golang.org/x/crypto/bcrypt"

	"home-task-tracker/application/core"
)

const (
	CREATE_CHILD_ACCOUNT           = "INSERT INTO htt.child_accounts(parent_id, first_name, last_name, email, login, password) VALUES($1,$2,$3,$4,$5,$6) returning id;"
	GET_CHILD_ACCOUNT_BY_ID        = "SELECT * FROM htt.child_accounts WHERE id=$1;"
	GET_CHILD_ACCOUNT_BY_PARENT_ID = "SELECT * FROM htt.child_accounts WHERE parent_id=$1;"
	UPDATE_CHILD_ACCOUNT           = "UPDATE htt.child_accounts SET first_name=$1, last_name=$2, email=$3, login=$4, password=$5 WHERE id=$6;"
	UPDATE_CHILD_ACCOUNT_POINTS    = "UPDATE htt.child_accounts SET points=$1 WHERE id=$2;"
	DELETE_CHILD_ACCOUNT           = "DELETE FROM htt.child_accounts WHERE id=$1;"
)

type ChildModel struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Points    string `json:"points"`
}

func Create(parentId string, c ChildModel) error {

	db := core.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(CREATE_CHILD_ACCOUNT)

	defer stmt.Close()

	pwd, _ := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	_, err = stmt.Exec(parentId, c.FirstName, c.LastName, c.Email, c.Login, string(pwd))

	return err
}

func GetById(id string) (ChildModel, error) {

	db := core.GetConnection()
	defer db.Close()

	child := ChildModel{}

	var parentId, firstName, lastName, email, login, password, points string
	err := db.QueryRow(GET_CHILD_ACCOUNT_BY_ID, id).Scan(&id, &parentId, &firstName, &lastName, &email, &login, &password, &points)

	child.FirstName = firstName
	child.LastName = lastName
	child.Email = email
	child.Login = login
	child.Password = password
	child.Points = points

	return child, err
}

func GetByParentId(id string) ([]ChildModel, error) {

	childs := make([]ChildModel, 0)

	db := core.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(GET_CHILD_ACCOUNT_BY_PARENT_ID)

	if err != nil {
		return childs, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return childs, err
	}

	for rows.Next() {

		var row ChildModel
		var id, parentId string
		err := rows.Scan(&id, &parentId, &row.FirstName, &row.LastName, &row.Email, &row.Login, &row.Password, &row.Points)

		if err != nil {
			return nil, err
		}

		childs = append(childs, row)
	}

	return childs, err
}

func Update(id string, c ChildModel) error {

	chld, err := GetById(id)

	if len(c.FirstName) > 0 {
		chld.FirstName = c.FirstName
	}
	if len(c.LastName) > 0 {
		chld.LastName = c.LastName
	}
	if len(c.Email) > 0 {
		chld.Email = c.Email
	}
	if len(c.Login) > 0 {
		chld.Login = c.Login
	}
	if len(c.Password) > 0 {
		pwd, _ := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
		chld.Password = string(pwd)
	}

	db := core.GetConnection()
	defer db.Close()

	if err != nil {
		return err
	}

	stmt, err := db.Prepare(UPDATE_CHILD_ACCOUNT)

	defer stmt.Close()

	_, err = stmt.Exec(chld.FirstName, chld.LastName, chld.Email, chld.Login, chld.Password, id)

	return err
}

func UpdatePoints(id string, c ChildModel) error {

	db := core.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(UPDATE_CHILD_ACCOUNT_POINTS)

	defer stmt.Close()

	_, err = stmt.Exec(c.Points, id)

	return err
}

func Delete(id string) error {

	db := core.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(DELETE_CHILD_ACCOUNT)

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}
