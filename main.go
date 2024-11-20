package main

import (
	"context"
	"fmt"

	"github.com/irtza33/basic-go-template/infrastructure"
)

func main() {
	fmt.Println("Hello, World!")

	_, err := infrastructure.NewSqlClient(context.Background() ,"", "", "", "")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database successfully connected")
}

// func getAllUsers(conn *pgx.Conn) ([]User, error) {
// 	rows, err := conn.Query(context.Background(), "SELECT * FROM users")
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()
// 	for rows.Next() {
// 		err := rows.Scan()
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	}
// }
