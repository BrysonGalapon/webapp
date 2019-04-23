package database

const (
	QUERY_TEMPLATE  = "SELECT %v FROM %v"
	INSERT_TEMPLATE = "INSERT INTO %v (%v) VALUES (%v)"
	DELETE_TEMPLATE = "DELETE FROM %v WHERE %v"
)
