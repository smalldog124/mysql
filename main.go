package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type UserData struct {
	Id              int
	CitizenId       string
	Firstname       string
	Lastname        string
	BirthYear       int
	FirstnameFather string
	LastnameFather  string
	FirstnameMother string
	LastnameMother  string
	SoldierId       int
	AddressId       int
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")

	if err != nil {
		fmt.Println("connect fail")
	}

	defer db.Close()
	
	fmt.Println(remove(db,"3"))
}

func read(db *sql.DB) []UserData{
	results, err := db.Query("SELECT * FROM user")

	var userDataList []UserData
	for results.Next() {
		var userData UserData

		err = results.Scan(
			&userData.Id,
			&userData.CitizenId,
			&userData.Firstname,
			&userData.Lastname,
			&userData.BirthYear,
			&userData.FirstnameFather,
			&userData.LastnameFather,
			&userData.FirstnameMother,
			&userData.LastnameMother,
			&userData.SoldierId,
			&userData.AddressId,
		)
		if err != nil {
			panic(err.Error())
		}

		userDataList = append(userDataList,userData)
		fmt.Println(userDataList)
	}

	return userDataList
	
}

func add(db *sql.DB) bool { //เก็บชุดคำสั่งไว้ในตัวแปร statement เพื่อการป้องกันความปลอดภัย เวลจะ insert ก็ทำการสั่ง exec เข้าไปในชุดคำสั่ง
	statement, _ := db.Prepare(`INSERT INTO user ( 
	 citizen_id,
	 firstname,
	 lastname,
	 birthyear,
	 firstname_father,
	 lastname_father,
	 firstname_mother,
	 lastname_mother,
	 soldier_id,
	 address_id) 
	 VALUES(?,?,?,?,?,?,?,?,?,?)`)

	_, err := statement.Exec("1209700620251",
		"นารีนารถ",
		"เนรัญชร",
		2538,
		"ณัฐพงษ์",
		"ฉิมวัย",
		"กานต์วัฒน์",
		"วงศ์อุดม",
		69,
		1,
	)

	if err != nil{
		panic(err.Error())
		return false
	}
	return true
	
}

func remove (db *sql.DB,id string) bool{
	statement,_ := db.Prepare("DELETE FROM user WHERE user_id =?")

	_, err := statement.Exec(id)
	if err != nil{
		panic(err.Error())
	return false
	}
	return true
}
