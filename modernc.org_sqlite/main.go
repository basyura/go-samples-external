package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	fmt.Println("start -----")
	if err := doMain(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("end   -----")
}

func doMain() error {
	db, err := sql.Open("sqlite", "file:sample.sqlite?cache=shared")
	if err != nil {
		return err
	}
	defer db.Close()

	if err := createTable(db); err != nil {
		return err
	}

	if err := addUsers(db); err != nil {
		return err
	}

	if err := selectUsers(db); err != nil {
		return err
	}

	return nil
}

func createTable(db *sql.DB) error {
	fmt.Println("User テーブルを作成")
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT NOT NULL,
		Age INTEGER
	)`)

	return err
}

func addUsers(db *sql.DB) error {
	fmt.Println("User レコードを追加")
	res, err := db.Exec(`
        Insert into Users Values (null, "A", 1);
        Insert into Users Values (null, "B", 2);
        Insert into Users Values (null, "C", 3);
        Insert into Users Values (null, "D", 4);
    `)
	if err != nil {
		return err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println(fmt.Printf("  LastInsertId: %d", lastId))

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(fmt.Printf("  RowsAffected: %d", affectedRows))

	return nil
}

func selectUsers(db *sql.DB) error {
	fmt.Println("User レコード検索 : latest 10")
	rows, err := db.Query(`SELECT id, Name, Age FROM Users ORDER BY id DESC LIMIT 10;`)
	if err != nil {
		return err
	}

	// 結果を一行ずつ取得し、fmt.Printlnで出力
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			return err
		}
		fmt.Println("  ID:", id, "Name:", name, "Age:", age)
	}

	return nil
}
