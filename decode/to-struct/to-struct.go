package main

import (
	"encoding/json"
	"fmt"
)

// bikin struct
type User struct {
	Name string
	Age  int
}

func main() {
	// siapkan json
	jsonString := `{"name":"John", "age":30}`

	// convert json ke byte
	data := []byte(jsonString)

	// inisiasi struct
	var user User

	// masukan json hasil convert ke struct
	err := json.Unmarshal(data, &user)

	// cek error
	if err != nil {
		fmt.Println(err.Error())
	}

	// cetak hasil
	fmt.Println(user)
}
