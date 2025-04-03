package main

//define functions for all common sql operations to use as a template

import (
	_ "github.com/lib/pq"
	"strconv"
)

// SelectFromTable Select from table
func SelectFromTable(table string, columns []string, where string) string {
	query := "SELECT "
	for i, column := range columns {
		query += column
		if i < len(columns)-1 {
			query += ", "
		}
	}
	query += " FROM " + table
	if where != "" {
		query += " WHERE " + where
	}
	return query
}

// InsertIntoTable Insert into table
func InsertIntoTable(table string, columns []string) string {
	query := "INSERT INTO " + table + " ("
	for i, column := range columns {
		query += column
		if i < len(columns)-1 {
			query += ", "
		}
	}
	query += ") VALUES ("
	for i := range columns {
		query += "$" + strconv.Itoa(i+1)
		if i < len(columns)-1 {
			query += ", "
		}
	}
	query += ")"
	return query
}

// UpdateTable Update table
func UpdateTable(table string, columns []string, where string) string {
	query := "UPDATE " + table + " SET "
	for i, column := range columns {
		query += column + " = $" + strconv.Itoa(i+1)
		if i < len(columns)-1 {
			query += ", "
		}
	}
	query += " WHERE " + where
	return query
}

// DeleteFromTable Delete from table
func DeleteFromTable(table string, where string) string {
	query := "DELETE FROM " + table
	if where != "" {
		query += " WHERE " + where
	}
	return query
}
