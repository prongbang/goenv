package main

import (
	"fmt"
	"os"

	_ "github.com/prongbang/goenv"
)

func main() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")

	fmt.Println("user:", user)
	fmt.Println("pass:", pass)
}
