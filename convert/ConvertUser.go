package convert

import (
	"database/sql"

	"1d1hphoto.com-be/db"
	"1d1hphoto.com-be/model"
)

func ParseUsers(rows *sql.Rows) []model.User {
	usersList := make([]model.User, 0)
	for rows.Next() {
		u := model.User{}
		err := rows.Scan(&u.Id, &u.Name, &u.Cognome)
		db.CheckErr(err)
		usersList = append(usersList, u)
	}
	return usersList
}
