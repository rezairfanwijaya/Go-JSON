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
	// siapkan json
	var jsonString = `[
		{"name":"John", "age":30},
		{"name":"Bruce", "age":10}
		]`

	// convert json ke byte
	var data = []byte(jsonString)

	// siapkan array/slice
	var pengguna []User

	// masukan json hasil convert ke array/slice
	err := json.Unmarshal(data, &pengguna)

	// cek error
	if err != nil {
		fmt.Println(err.Error())
	}

	// cetak hasil
	fmt.Println(pengguna)

}
