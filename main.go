package main

import (
	"fmt"
	"os"
	"unit_tests_stream/internal/business/domain/users"
)

func main() {
	usersDataFile, err := os.Open("data/users.json")
	if err != nil {
		panic(err)
	}

	provider := users.NewProvider(usersDataFile)
	service := users.NewService(provider)

	fmt.Println(service.IsUserExists("ehot_riss"))
	fmt.Println(service.IsUserExists("enot_riss"))
}
