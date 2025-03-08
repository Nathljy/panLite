package database

import (
	"database/sql"
	"fmt"
	"log"
	"panLite/resp"
	"reflect"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	DB *sql.DB
}

var (
	InitOnce sync.Once
	DB       *Database
)

func NewDatabase() *Database {
	InitOnce.Do(func() {
		dsn := "root:666666@tcp(127.0.0.1:3306)/pan_lite_database?charset=utf8&parseTime=True&loc=Local"
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
		}
		DB = &Database{db}
	})
	return DB
}

func (db *Database) Close() {
	err := db.DB.Close()
	if err != nil {
		log.Println(err)
	}
}

// Read single data
func Read(obj interface{}, query string, args ...interface{}) *resp.Error {
	argsMarshal := make([]interface{}, len(args))
	for i, v := range args {
		argsMarshal[i] = fmt.Sprintf("%v", v)
	}
	row := DB.DB.QueryRow(query, argsMarshal...)

	val := reflect.ValueOf(obj).Elem()
	vars := make([]interface{}, 0)
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		tag := field.Tag.Get("db")
		if tag != "" {
			vars = append(vars, val.FieldByName(field.Name).Addr().Interface())
		}
	}
	err := row.Scan(vars...)
	if err != nil {
		return resp.DBErr
	}

	return nil
}

// Reads varies data
func Reads[T any](objs *[]*T, query string, args ...interface{}) *resp.Error {
	argsMarshal := make([]interface{}, len(args))
	for i, v := range args {
		argsMarshal[i] = fmt.Sprintf("%v", v)
	}
	rows, err := DB.DB.Query(query, argsMarshal...)
	if err != nil {
		return resp.DBErr
	}

	t := reflect.TypeOf(new(T)).Elem()
	if t.Kind() != reflect.Struct {
		return resp.DBErr
	}

	for rows.Next() {
		obj := new(T)
		val := reflect.ValueOf(obj).Elem()
		vars := make([]interface{}, 0)

		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			tag := field.Tag.Get("db")
			if tag != "" {
				vars = append(vars, val.FieldByName(field.Name).Addr().Interface())
			}
		}

		if err := rows.Scan(vars...); err != nil {
			return resp.DBErr
		}

		*objs = append(*objs, obj)
	}

	if err := rows.Err(); err != nil {
		return resp.DBErr
	}
	return nil
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	argsMarshal := make([]interface{}, len(args))
	for i, v := range args {
		argsMarshal[i] = fmt.Sprintf("%v", v)
	}
	result, err := DB.DB.Exec(query, argsMarshal...)
	if err != nil {
		return nil, resp.DBErr
	}
	return result, nil
}
