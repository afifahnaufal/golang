package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func getUser(c *gin.Context) {
    user := User{
        ID:    1,
        Name:  "Afifah naufal",
        Email: "afifah@example.com",
    }
    c.JSON(http.StatusOK, user)
}

func main() {
    r := gin.Default()
    r.GET("/user", getUser)
    r.Run() // default di :8080
}
