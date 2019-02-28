package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	//query user

	sqlStatement := `SELECT customer_id, phone FROM customer.identity WHERE is_actived=$1;`
	var phone string
	var id int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := db.QueryRow(sqlStatement, true)
	switch err := row.Scan(&id, &phone); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, phone)
	default:
		panic(err)
	}

	//query multiple rows

	rows, err := db.Query(sqlStatement, true)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var customerId int
		var phoneNumber string
		err = rows.Scan(&customerId, &phoneNumber)
		if err != nil {
			panic(err)
		}
		fmt.Println(customerId, phoneNumber)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

}
