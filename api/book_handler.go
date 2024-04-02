package api

import (
	"strconv"
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
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()) //IDの変換に失敗したら400BadRequestエラーを返す 
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

// 本の情報を更新するハンドラー関数
func UpdateBookHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// URLパラメータから本のIDを取得
		id, err := strconv.ParseInt(c.Param("id"), 10, 64) // 10進数、64bit
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
		}

		// リクエストボディから更新情報を取得
		var bookUpdate models.Book 
		if err := c.Bind(&bookUpdate); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// 本の情報を更新
		err = repository.UpdateBook(db, id, bookUpdate.Title, bookUpdate.Author)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// 成功レスポンスを返す
		return c.JSON(http.StatusOK, map[string]string{"message": "Book updated successfully"})
	}
}

// 本を削除するハンドラー関数
func DeleteBookHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// URLパラメータから本のIDを取得
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
		}

		// 本を削除
		err = repository.DeleteBook(db, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// 成功レスポンスを返す
		return c.JSON(http.StatusOK, map[string]string{"message": "Book deleted successfully"})
	}
}
