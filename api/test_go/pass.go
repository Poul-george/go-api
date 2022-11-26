package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func main() {
	var str string
    fmt.Scan(&str) // データを格納する変数のアドレスを指定
	hash, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
    fmt.Println(string(hash))

	var pass string
    fmt.Scan(&pass)

	fmt.Println(bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)))
} 