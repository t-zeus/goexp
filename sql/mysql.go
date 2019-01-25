package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init mysql driver
)

type user struct {
	Id   int
	Name sql.NullString
	Age  sql.NullInt64
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// query data
	rows, err := db.Query("select * from user")
	defer rows.Close() // 把数据库连接放回连接池
	if err != nil {
		panic(err)
	}

	user := user{}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Println(err)
			continue
		}
		/*
			if !user.Name.Valid {
				user.Name.String = "xxx"
			}
		*/
		fmt.Println(user, user.Name.String, user.Age.Int64)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	// query one row
	var name string
	err = db.QueryRow("select name from user where id = ?", 3).Scan(&name)
	if err != nil {
		panic(err)
	}
	fmt.Println(name)

	// insert
	ret, err := db.Exec(`INSERT INTO user (name, age) VALUES ("zhangsan", 22)`)
	if err != nil {
		panic(err)
	}
	if lastInsertId, err := ret.LastInsertId(); err == nil {
		fmt.Println("last_insert_id =", lastInsertId)
	}
	if rowsAffected, err := ret.RowsAffected(); err == nil {
		fmt.Println("rows_affected =", rowsAffected)
	}
}
