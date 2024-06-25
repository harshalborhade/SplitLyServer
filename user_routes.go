package main

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func createUser(context *gin.Context) {
	var user User
	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//insert user info into the database
	_, err := Db.Exec("INSERT INTO splitly.user_dimension (id, first_name, last_name, profile_picture_url) VALUES ($1, $2, $3)", user.Id, user.FirstName, user.LastName, user.ProfilePic)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, user)
}

func getUsers(c *gin.Context) {
	rows, err := Db.Query("SELECT id, first_name, last_name, profile_picture_url FROM splitly.user_dimension")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.ProfilePic); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	err := Db.QueryRow("SELECT id, first_name,last_name,profile_picture_url FROM splitly.user_dimension WHERE id = $1", id).
		Scan(&user.Id, &user.FirstName, &user.LastName, &user.ProfilePic)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// TODO: update this to update user info
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user UserCreds
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the new password if provided
	var hashedPassword []byte
	var err error
	if user.Password != "" {
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
	}

	// Update user in the database
	var result sql.Result
	if user.Password != "" {
		result, err = Db.Exec("UPDATE splitly.user_credentials SET username = $1, email = $2, password = $3 WHERE id = $4",
			user.Username, user.Email, string(hashedPassword), id)
	} else {
		result, err = Db.Exec("UPDATE splitly.user_credentials SET username = $1, email = $2 WHERE id = $3",
			user.Username, user.Email, id)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "UserCreds not found"})
		return
	}

	user.Password = "" // Don't send the password back
	c.JSON(http.StatusOK, user)
}

// TODO: update this to handle credentials and information
func deleteUser(c *gin.Context) {
	id := c.Param("id")

	result, err := Db.Exec("DELETE FROM splitly.user_credentials WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking affected rows"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "UserCreds not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "UserCreds successfully deleted"})
}
