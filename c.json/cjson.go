package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Struct User dengan field lengkap
type User struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Email       string `json:"email"`
    Status      string `json:"status"`
    Universitas string `json:"universitas"`
    Jurusan     string `json:"jurusan"`
}

// Fungsi getUser mengembalikan data User
func getUser() User {
    return User{
        ID:          1,
        Name:        "Afifah Naufal Rahma",
        Email:       "afifah@example.com",
        Status:      "Mahasiswa",
        Universitas: "Logistik dan Bisnis International",
        Jurusan:     " D4 Teknik Informatika",
    }
}

func main() {
    r := gin.Default()

    // Route GET /user mengembalikan JSON
    r.GET("/user", func(c *gin.Context) {
        user := getUser()
        c.JSON(http.StatusOK, gin.H{
            "status": "success",
            "data":   user,
        })
    })

    r.Run() // listen di localhost:8080
}
