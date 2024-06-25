package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func register(c *gin.Context) {
	var user UserCreds
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Insert the new user into the database
	err = Db.QueryRow("INSERT INTO splitly.user_credentials (username, email, password) VALUES ($1, $2, $3) RETURNING id, created_at",
		user.Username, user.Email, string(hashedPassword)).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	user.Password = "" // Don't send the password back
	c.JSON(http.StatusCreated, user)
}

func login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user UserCreds
	var hashedPassword string
	err := Db.QueryRow("SELECT id, username, email, password, created_at FROM splitly.user_credentials WHERE username = $1",
		loginData.Username).Scan(&user.ID, &user.Username, &user.Email, &hashedPassword, &user.CreatedAt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	user.Password = "" // Don't send the password back
	c.JSON(http.StatusOK, user)
}
