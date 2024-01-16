package db

import (
	"database/sql"
	"errors"
	"fmt"
	"screening/internal/config"

	"github.com/opentracing/opentracing-go/log"
)

var SQLDB *sql.DB

func InitSQLConn() {
	db, err := sql.Open("mysql", config.GlobalConfig.MySQLURI)
	if err != nil {
		panic(fmt.Sprintf("sql open %s", err.Error()))
	}
	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("sql ping error %s", err.Error()))
	}
	SQLDB = db
}

func InsertUser(name, email string) (int64, error) {

	stmt, err := SQLDB.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	defer stmt.Close()
	if err != nil {
		log.Error(err)
		return 0, errors.New("Error creating statement")
	}
	res, err := stmt.Exec(name, email)
	if err != nil {
		log.Error(err)
		return 0, errors.New("Error inserting user")
	}
	insertId, err := res.LastInsertId()
	if err != nil {
		log.Error(err)
		return 0, errors.New("Error fetching inserted userid")
	}
	return insertId, nil

}

func UpdateUser(name, email string, id int) (int64, error) {
	stmt, err := SQLDB.Prepare("UPDATE users SET name=?, email=? WHERE id=?")
	defer stmt.Close()

	if err != nil {
		log.Error(err)
		return 0, errors.New("Error creating statement")

	}
	res, err := stmt.Exec(name, email, id)
	if err != nil {
		log.Error(err)
		return 0, errors.New("Error updating user")

	}
	updatedRowCount, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return 0, errors.New("Error getting updated row count")

	}
	return updatedRowCount, nil
}
