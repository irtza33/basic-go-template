package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
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
