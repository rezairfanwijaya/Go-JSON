package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

type Users struct {
	Users []User
}

func main() {
	// siapkan array atau slice
	var user User
	user.Name = "John"
	user.Age = 30

	var users Users
	users.Users = append(users.Users, user)

	// masukan array/slice ka marshal
	data, err := json.Marshal(users)

	// cek error
	if err != nil {
		fmt.Println(err.Error())
	}

	// convert to string
	jsonString := string(data)

	// cetak hasil
	fmt.Println(jsonString)
}
