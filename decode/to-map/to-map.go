package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// siapkan json
	jsonString := `{"name":"John", "age":30}`

	// convert json ke byte
	data := []byte(jsonString)

	// siapkan map
	var user map[string]interface{}

	// masukan json hasil convert ke map
	err := json.Unmarshal(data, &user)

	// cek error
	if err != nil {
		fmt.Println(err.Error())
	}

	// cetak hasil
	fmt.Println(user)

}
