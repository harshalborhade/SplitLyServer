package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {

	connectToDB()
	// Database connection parameters
	defer func(Db *sql.DB) {
		err := Db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)

	err := Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")

	router := gin.Default()

	// Login and register routes
	router.POST("/register", register)
	router.POST("/login", login)

	// User routes
	router.POST("/createUser", createUser)
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
