package main

import (
	"database/sql"
	"log"

	"github.com/koujirou513/go-book/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mattn/go-sqlite3" // SQLite3 driver
)

func main() {
	// SQLiteデータベースに接続
	db, err := sql.Open("sqlite3", "./yourdatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// データベース初期化
	if err := initDB(db); err != nil {
		log.Fatal("Failed to initialize the database", err)
	}

	// Echoインスタンスの作成
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())  //リクエストに関する情報をログに記録する
	e.Use(middleware.Recover()) //エラーハンドリング。ハンドラーで発生したパニックをキャッチし、サーバーのクラッシュを防ぐ

	//ルーティング
	e.GET("/books", api.GetAllBooksHandler(db))
	e.POST("/books", api.CreateBookHandler(db))
	e.PUT("/books/:id", api.UpdateBookHandler(db))
	e.DELETE("/books/:id", api.DeleteBookHandler(db))

	e.Static("/", "static")

	// サーバーをポート 8080で起動
	e.Logger.Fatal(e.Start(":8080"))

}

func initDB(db *sql.DB) error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}
