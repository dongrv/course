package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the database successfully.")

	// 查询记录
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("User ID: %d, Name: %s\n", id, name)
	}

	// 插入记录
	result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "John Doe", "john.doe@example.com")
	if err != nil {
		log.Fatal(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Last Insert ID: %d, Rows Affected: %d\n", lastInsertId, rowsAffected)

	// 更新记录
	_, err = db.Exec("UPDATE users SET name = ? WHERE id = ?", "Jane Doe", 1)
	if err != nil {
		log.Fatal(err)
	}

	// 删除记录
	_, err = db.Exec("DELETE FROM users WHERE id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	// 事务
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// 执行多个操作
	_, err = tx.Exec("UPDATE users SET balance = balance - 10 WHERE id = ?", 1)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	_, err = tx.Exec("INSERT INTO logs (user_id, action) VALUES (?, 'debit')", 1)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
