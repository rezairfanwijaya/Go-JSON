package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// isi struct
	var user User
	user.Name = "John"
	user.Age = 30

	// masukan struct ke function masrshal
	data, err := json.Marshal(user)

	// cek error
	if err != nil {
		fmt.Println(err.Error())
	}

	// convert to json string
	jsonString := string(data)

	// cetak hasil
	fmt.Println(jsonString)

}
