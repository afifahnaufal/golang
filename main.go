package main

import "fmt"

// Struct
type User struct {
    ID    int
    Name  string
    Email string
}

// Fungsi
func getUser() User {
    return User{
        ID:    1,
        Name:  "Afifah",
        Email: "afifah@example.com",
    }
}

func main() {
    user := getUser()
    fmt.Println("User:", user)
}
