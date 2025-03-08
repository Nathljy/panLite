package dao

import (
	"database/sql"
	"panLite/model"
	"panLite/module/database"
	"panLite/resp"
)

const (
	vars       = " id, type, name, password "
	selectUser = "SELECT " + vars + " FROM user "

	insertUser    = `INSERT INTO user (type, name, password) VALUES (?, ?, ?)`
	queryUserByID = selectUser + ` WHERE id=?`
)

type UserRow struct {
	ID   int    `db:"id"`
	Type string `db:"type"`
	Name string `db:"name"`
	Pass string `db:"password"`
}

// Query

func GetByID(id uint64) (*UserRow, *resp.Error) {
	var u []*UserRow
	err := database.Reads[UserRow](&u, queryUserByID, id)
	if err != nil {
		return nil, resp.DBErr
	}
	return u[0], nil
}

// Execute

type CreateUserDao struct {
	Type        model.UserType
	Name        string
	EncryptPass string
}

func (dao CreateUserDao) Do() (*sql.Result, *resp.Error) {
	result, err := database.Exec(insertUser, dao.Type, dao.Name, dao.EncryptPass)
	if err != nil {
		return nil, resp.DBErr
	}
	return &result, nil
}
