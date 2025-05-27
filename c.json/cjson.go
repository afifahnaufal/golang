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
        Name:        "Afifah Naufal Rahmani",
        Email:       "afifah@example.com",
        Status:      "Mahasiswa",
        Universitas: "Logistik dan Bisnis Internasional",
        Jurusan:     "D4 Teknik Informatika",
    }
}

func main() {
    r := gin.Default()

    // Endpoint GET /user (huruf kecil, sesuai standar REST)
    r.GET("/user", func(c *gin.Context) {
        user := getUser()
        c.JSON(http.StatusOK, gin.H{
            "status": "success",
            "data":   user,
        })
    })

    // Jalankan server di localhost:8080
    r.Run() // default :8080
}
