package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-postgres/config"
	_ "github.com/lib/pq"
)

var (
	db1  *sql.DB
	once sync.Once
)

func init() {
	fmt.Println("Start db connection")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config.Db.Host, 5432, config.Config.Db.User, config.Config.Db.Password, config.Config.Db.Name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Fail to connect.")
		panic(err)
	}
	// defer db.Close()
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(3600 * time.Second)
	err = db.Ping()
	if err != nil {
		fmt.Println("Fail to ping.")
		panic(err)
	}

	fmt.Println("Successfully connected1!")
	db1 = db
}

func new() *sql.DB {
	fmt.Println("Start db connection")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config.Db.Host, 5432, config.Config.Db.User, config.Config.Db.Password, config.Config.Db.Name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Fail to connect.")
		panic(err)
	}
	// defer db.Close()

	// db.DB().SetMaxIdleConns(20)
	// db.DB().SetMaxOpenConns(100)
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(3600 * time.Second)
	err = db.Ping()
	if err != nil {
		fmt.Println("Fail to ping.")
		panic(err)
	}

	fmt.Println("Successfully connected!2")
	return db
}

func getDB() *sql.DB {
	once.Do(func() {
		db1 = new()
	})
	return db1
}

func listIdentityHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path

	fmt.Printf("We have a request at url: %s", message)
	//query user
	msg := ""
	sqlStatement := `SELECT id, phone FROM identity WHERE id<$1;`
	var phone string
	var id int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	db := new()
	row := db.QueryRow(sqlStatement, 100)
	switch err := row.Scan(&id, &phone); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		msg += fmt.Sprintln("No rows were returned!")
	case nil:
		fmt.Println(id, phone)
		msg += fmt.Sprintln(id, phone)
	default:
		panic(err)
	}

	//query multiple rows

	rows, err := db.Query(sqlStatement, 100)

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
		msg += fmt.Sprintln(customerId, phoneNumber)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/", listIdentityHandler)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}
