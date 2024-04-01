package api

import (
	"database/sql"
	"net/http"
	"github.com/koujirou513/go-book/models"
	"github.com/koujirou513/go-book/repository"
	"github.com/labstack/echo/v4"
)

// すべての本を取得するためのハンドラー関数
func GetAllBooksHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// repository から本の一覧を取得
		books, err := repository.GetAllBooks(db)
		if err != nil {
			// サーバーエラーを返す
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		// 取得した本の一覧をJSONとして返す
		return c.JSON(http.StatusOK, books)
	}
}

// 新しい本を追加するためのハンドラー関数
func CreateBookHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// リクエストボディから本の情報を取得
		var newBook models.Book
		if err := c.Bind(&newBook); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// 新しい本をデータベースに追加
		id, err := repository.CreateBook(db, newBook.Title, newBook.Author)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// 追加した本のIDを含むレスポンスを返す
		return c.JSON(http.StatusCreated, map[string]int64{"id": id})
	}
}
