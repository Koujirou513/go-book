package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/koujirou513/go-book/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func main() {
	// MySQLデータベースに接続
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")

	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d): %s", i+1, err)
		time.Sleep(2 * time.Second)
	}
	
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
	e.Use(middleware.CORS())    // CORSミドルウェア

	//ルーティング
	e.GET("/books", api.GetAllBooksHandler(db))
	e.POST("/books", api.CreateBookHandler(db))
	e.PUT("/books/:id", api.UpdateBookHandler(db))
	e.DELETE("/books/:id", api.DeleteBookHandler(db))

	e.Static("/", "static")
	e.GET("/health", healthCheckHandler)

	// サーバーをポート 8080で起動
	e.Logger.Fatal(e.Start(":8080"))

}

// ヘルスチェック用関数
func healthCheckHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func initDB(db *sql.DB) error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS books (
		id INT PRIMARY KEY AUTO_INCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}
