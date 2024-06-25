package main

import (
	"database/sql"
	"fmt"
	"log"
)

var dbCred = Credentials{"localhost", 5432, "hb", "123123", "splitly_db"}

var Db *sql.DB

func connectToDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbCred.Host, dbCred.Port, dbCred.User, dbCred.Password, dbCred.Dbname)

	var err error
	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
}
