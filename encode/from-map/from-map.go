package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// siapkan map
	var user map[string]interface{}
	user = map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	// masukan map ke function marshal
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
