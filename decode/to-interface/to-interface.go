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

	// siapkan interface
	var user interface{}

	// masukan json hasil convert ke interface
	err := json.Unmarshal(data, &user)

	// cek error
	if err != nil {
		fmt.Println(err.Error())
	}

	// cetak hasil
	fmt.Println(user)
}
