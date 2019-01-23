package parent

import (
	"golang.org/x/crypto/bcrypt"

	"home-task-tracker/application/core"
)

const (
	CREATE_PARENT_ACCOUNT    = "INSERT INTO htt.parent_accounts(first_name, last_name, email, login, password) VALUES($1,$2,$3,$4,$5) returning id;"
	GET_PARENT_ACCOUNT_BY_ID = "SELECT * FROM htt.parent_accounts WHERE id=$1;"
	UPDATE_PARENT_ACCOUNT    = "UPDATE htt.parent_accounts SET first_name=$1, last_name=$2, email=$3, login=$4, password=$5 WHERE id=$6;"
	DELETE_PARENT_ACCOUNT    = "DELETE FROM htt.parent_accounts WHERE id=$1;"
)

type ParentModel struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

func Create(p ParentModel) error {

	db := core.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(CREATE_PARENT_ACCOUNT)

	defer stmt.Close()

	pwd, _ := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	_, err = stmt.Exec(p.FirstName, p.LastName, p.Email, p.Login, string(pwd))

	return err
}

func GetById(id string) (ParentModel, error) {

	db := core.GetConnection()
	defer db.Close()

	parent := ParentModel{}

	var firstName, lastName, email, login, password string
	err := db.QueryRow(GET_PARENT_ACCOUNT_BY_ID, id).Scan(&id, &firstName, &lastName, &email, &login, &password)

	parent.FirstName = firstName
	parent.LastName = lastName
	parent.Email = email
	parent.Login = login
	parent.Password = password

	return parent, err
}

func Update(id string, p ParentModel) error {

	prnt, err := GetById(id)

	if len(p.FirstName) > 0 {
		prnt.FirstName = p.FirstName
	}
	if len(p.LastName) > 0 {
		prnt.LastName = p.LastName
	}
	if len(p.Email) > 0 {
		prnt.Email = p.Email
	}
	if len(p.Login) > 0 {
		prnt.Login = p.Login
	}
	if len(p.Password) > 0 {
		pwd, _ := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
		prnt.Password = string(pwd)
	}

	db := core.GetConnection()
	defer db.Close()

	if err != nil {
		return err
	}

	stmt, err := db.Prepare(UPDATE_PARENT_ACCOUNT)

	defer stmt.Close()

	_, err = stmt.Exec(prnt.FirstName, prnt.LastName, prnt.Email, prnt.Login, prnt.Password, id)

	return err
}

func Delete(id string) error {

	db := core.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(DELETE_PARENT_ACCOUNT)

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}
