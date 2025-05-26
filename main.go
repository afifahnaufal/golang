package main

import "fmt"

// Struct User dengan field tambahan sesuai request
type User struct {
    ID          int
    Name        string
    Email       string
    Status      string
    Universitas string
    Jurusan     string
}

// Fungsi getUser mengembalikan data User lengkap
func getUser() User {
    return User{
        ID:          1,
        Name:        "Afifah Naufal Rahmani",
        Email:       "afifah@example.com",
        Status:      "Mahasiswa",
        Universitas: "Logistik Bisnis International",
        Jurusan:     "Informatika",
    }
}

func main() {
    user := getUser()
    fmt.Println("User:", user)
}
