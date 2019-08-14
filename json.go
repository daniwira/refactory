package main

import "encoding/json"
import "fmt"

type User struct {
	User      string `json:"User"`
	Address   string `json: Address`
	Latitude  string `json: Latitude`
	Longitude string `json: Longitude`
	Company   string `json:"Company"`
}

func main() {
	var jsonString = `{"User" : "Dani", "Address" : "jl H Agus Salim rt 6 rw 2 kebayoran baru", "Latitude": "-1.7469322", "Longitude": "140.6689922", "Company": "Kudo"}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User :", data.User)
	fmt.Println("Adress :", data.Address)
	fmt.Println("Latitude :", data.Latitude)
	fmt.Println("Longitude :", data.Longitude)
	fmt.Println("Company :", data.Company)
}
