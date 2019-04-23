package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

type DBHandler struct {
	Driver *sql.DB
}

var (
	db  *sql.DB
	err error
)

func LaunchDB() *DBHandler {
	db, err = sql.Open("mysql", "root:password@/webapp")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Launched MySQL database")
	return &DBHandler{Driver: db}
}

func (dbHandler *DBHandler) View(fields []string, table string) *sql.Rows {
	fieldStr := strings.Join(fields, ", ")
	queryString := fmt.Sprintf(QUERY_TEMPLATE, fieldStr, table)
	results, err := db.Query(queryString)

	if err != nil {
		log.Fatal(err)
	}

	return results
}

func (dbHandler *DBHandler) Insert(fields []string, values []string, table string) {
	fieldStr := strings.Join(fields, ", ")
	valueStr := strings.Join(values, ", ")
	queryString := fmt.Sprintf(INSERT_TEMPLATE, table, fieldStr, valueStr)
	_, err := db.Query(queryString)

	if err != nil {
		log.Fatal(err)
	}
}

func (dbHandler *DBHandler) Delete(fields []string, values []string, table string) {
	conditions := buildConditions(fields, values)
	whereStr := strings.Join(conditions, " AND ")
	queryString := fmt.Sprintf(DELETE_TEMPLATE, table, whereStr)
	_, err := db.Query(queryString)

	if err != nil {
		log.Fatal(err)
	}
}

func buildConditions(fields []string, values []string) []string {
	if len(fields) != len(values) {
		log.Fatal("field length and value length must be equal")
	}

	conditions := []string{}

	for i, _ := range fields {
		conditions = append(conditions, fields[i]+"="+values[i])
	}

	return conditions
}

func CloseDB() {
	db.Close()
}
