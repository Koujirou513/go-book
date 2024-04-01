package repository

import (
	"database/sql"
	// "fmt"
	"github.com/koujirou513/go-book/models"

	_ "github.com/mattn/go-sqlite3"
)

// 本の情報を取得
func GetAllBooks(db *sql.DB) ([]models.Book, error) {
	books := []models.Book{}

	rows, err := db.Query("SELECT id, title, author FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

// 本を登録する
func CreateBook(db *sql.DB, title string, author string) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO books(title, author) VALUES(?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(title, author)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
