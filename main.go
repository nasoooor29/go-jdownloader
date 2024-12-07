package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	email := os.Getenv("EMAIL")
	pass := os.Getenv("PASS")

	j := NewJdownloader(email, pass)
	r, err := j.Connect()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("r: %#v\n", r)
}
