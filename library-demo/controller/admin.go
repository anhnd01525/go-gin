package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"library-demo/model"
	"log"
	"net/http"
	"time"
	"xid"
)

//CREATE USER
func (c *Controller) CreateUser(ctx *gin.Context) {
	var user model.CreateUser
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var newUser model.User
	copier.Copy(&newUser, &user)
	newUser.Id = xid.New().String()
	err = c.DB.Insert(&newUser)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể tạo user"))
		return
	}

	ctx.String(http.StatusOK, "Tạo user thành công")
}

//ALL USERS
func (c *Controller) GetUsers(ctx *gin.Context) {
	var users []model.GetUsers
	rs, err := c.DB.Query(&users, `SELECT * FROM library.users`)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Lỗi"))
		return
	}

	if rs.RowsReturned() == 0 {
		model.NewError(ctx, http.StatusNotFound, errors.New("Không có user"))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

//CREATE AUTHOR
func (c *Controller) CreateAuthor(ctx *gin.Context) {
	var author model.CreateAuthor
	err := ctx.ShouldBindJSON(&author)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var newAuthor model.Author
	copier.Copy(&newAuthor, &author)
	newAuthor.Id = xid.New().String()
	err = c.DB.Insert(&newAuthor)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể tạo tác giả"))
		return
	}

	ctx.String(http.StatusOK, "Tạo tác giả thành công")
}

//ALL AUTHORS
func (c *Controller) GetAuthors(ctx *gin.Context) {
	var authors []model.GetAuthors
	rs, err := c.DB.Query(&authors, `SELECT * FROM library.authors`)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Lỗi"))
		return
	}

	if rs.RowsReturned() == 0 {
		model.NewError(ctx, http.StatusNotFound, errors.New("Không có tác giả"))
		return
	}

	ctx.JSON(http.StatusOK, authors)
}

//UPDATE AUTHOR
func (c *Controller) UpdateAuthorById(ctx *gin.Context) {
	var author model.UpdateAuthorById
	err := ctx.ShouldBindJSON(&author)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var updateAuthor model.Author
	copier.Copy(&updateAuthor, &author)
	err = c.DB.Update(&updateAuthor)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể cập nhật tác giả"))
		return
	}

	ctx.String(http.StatusOK, "Cập nhật thành công")
}

//CREATE TYPE
func (c *Controller) CreateType(ctx *gin.Context) {
	var types model.CreateType
	err := ctx.ShouldBindJSON(&types)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var newType model.Type
	copier.Copy(&newType, &types)
	newType.Id = xid.New().String()
	err = c.DB.Insert(&newType)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể tạo thể loại"))
		return
	}

	ctx.String(http.StatusOK, "Tạo thể loại thành công")
}

//ALL TYPES
func (c *Controller) GetTypes(ctx *gin.Context) {
	var types []model.GetTypes
	rs, err := c.DB.Query(&types, `SELECT * FROM library.types`)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Lỗi"))
		return
	}

	if rs.RowsReturned() == 0 {
		model.NewError(ctx, http.StatusNotFound, errors.New("Không có thể loại nào"))
		return
	}

	ctx.JSON(http.StatusOK, &types)
}

//UPDATE TYPE
func (c *Controller) UpdateTypeById(ctx *gin.Context) {
	var types model.UpdateTypeById
	err := ctx.ShouldBindJSON(&types)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var updateType model.Type
	copier.Copy(&updateType, &types)
	err = c.DB.Update(&updateType)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể cập nhật thể loại"))
		return
	}

	ctx.String(http.StatusOK, "Cập nhật thành công")
}

//CREATE PUBLISHER
func (c *Controller) CreatePublisher(ctx *gin.Context) {
	var publisher model.CreatePublisher
	err := ctx.ShouldBindJSON(&publisher)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var newPublisher model.Publisher
	copier.Copy(&newPublisher, &publisher)
	newPublisher.Id = xid.New().String()
	err = c.DB.Insert(&newPublisher)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể tạo nhà xuất bản"))
		return
	}

	ctx.String(http.StatusOK, "Tạo nhà xuất bản thành công")
}

//ALL PUBLISHERS
func (c *Controller) GetPublishers(ctx *gin.Context) {
	var publishers []model.GetPublishers
	rs, err := c.DB.Query(&publishers, `SELECT * FROM library.publishers`)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Lỗi"))
		return
	}

	if rs.RowsReturned() == 0 {
		model.NewError(ctx, http.StatusNotFound, errors.New("Không có nhà xuất bản nào"))
		return
	}

	ctx.JSON(http.StatusOK, &publishers)
}

//UPDATE PUBLISHER
func (c *Controller) UpdatePublisherById(ctx *gin.Context) {
	var publisher model.UpdatePublisherById
	err := ctx.ShouldBindJSON(&publisher)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var updatePublisher model.Publisher
	copier.Copy(&updatePublisher, &publisher)
	err = c.DB.Update(&updatePublisher)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể cập nhật nhà xuất bản"))
		return
	}

	ctx.String(http.StatusOK, "Cập nhật thành công")
}

//CREATE BOOK
func (c *Controller) CreateBook(ctx *gin.Context) {
	var book model.CreateBook
	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var newBook model.Book
	copier.Copy(&newBook, &book)
	newBook.Id = xid.New().String()
	err = c.DB.Insert(&newBook)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể thêm sách mới"))
		return
	}

	ctx.String(http.StatusOK, "Thêm sách thành công")
}

//ALL BOOKS
func (c *Controller) GetBooks(ctx *gin.Context) {
	var books []model.GetBooks
	rs, err := c.DB.Query(&books, `SELECT * FROM library.books`)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	if rs.RowsReturned() == 0 {
		model.NewError(ctx, http.StatusNotFound, errors.New("Không có cuốn sách nào"))
		return
	}

	ctx.JSON(http.StatusOK, &books)
}

//UPDATE BOOK
func (c *Controller) UpdateBookById(ctx *gin.Context) {
	var book model.UpdateBookById
	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var updateBook model.Book
	copier.Copy(&updateBook, &book)
	err = c.DB.Update(&updateBook)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể cập nhật sách"))
		return
	}

	ctx.String(http.StatusOK, "Cập nhật thành công")
}

//BORROW BOOK
func (c *Controller) BorrowBooks(ctx *gin.Context) {
	var bill model.Bill
	err := ctx.ShouldBindJSON(&bill)
	borrowID := xid.New().String()

	check := true
	for _, item := range bill.Books {
		var book model.Book

		rs, err := c.DB.Query(&book, `SELECT * FROM library.books WHERE id = ?`, item)
		if err != nil {
			model.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		if rs.RowsReturned() == 0 {
			model.NewError(ctx, http.StatusNotFound, errors.New("Không có sách " + item + " này"))
			check = false
			continue
		}

		if book.Status == 1 {
			model.NewError(ctx, http.StatusNotFound, errors.New("Sách " + book.Title + " đã được cho mượn"))
			check = false
			continue
		}

		if book.Status == 2 {
			model.NewError(ctx, http.StatusNotFound, errors.New("Sách " + book.Title + " này đã bị mất"))
			check = false
			continue
		}
	}

	if !check {
		return
	}

	for _, item := range bill.Books {
		var detail model.Detail
		detail.BorrowId = borrowID
		detail.BookId = item

		err = c.DB.Insert(&detail)
		if err != nil {
			model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể thêm sách này"))
			return
		}

		_, err = c.DB.Exec(`UPDATE library.books SET status = 1 WHERE id = ?`, item)
		if err != nil {
			model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể cập nhật trạng thái sách"))
			return
		}
	}

	var borrow model.Borrow
	borrow.Id = borrowID
	borrow.UserId = bill.UserId
	borrow.Loanday = time.Now()
	borrow.Returnterm = borrow.Loanday.AddDate(0, 1, 0)

	err = c.DB.Insert(&borrow)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể mượn sách"))
		return
	}

	ctx.String(http.StatusOK, "Mượn sách thành công")
}

//RETURN BOOK
func (c *Controller) ReturnBooks(ctx *gin.Context) {
	var returnBooks model.ReturnBooks
	err := ctx.ShouldBindJSON(&returnBooks)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	check := true
	for _, item := range returnBooks.Books {
		var detail model.Detail

		rs, err := c.DB.Query(&detail, `SELECT * FROM library.details WHERE borrow_id = ? AND book_id = ?`, returnBooks.BorrowId, item.BookId)
		if err != nil {
			model.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		if rs.RowsReturned() == 0 {
			model.NewError(ctx, http.StatusNotFound, errors.New("Lỗi trả sách " + item.BookId + "!"))
			check = false
			continue
		}
	}

	if !check {
		return
	}

	for _, item := range returnBooks.Books {
		if item.Status == "Ok" {
			var detail model.Detail
			detail.Returnday = time.Now()

			_, err := c.DB.Exec(`UPDATE library.details SET returnday = ? WHERE book_id = ?`, detail.Returnday, item.BookId)
			if err != nil {
				model.NewError(ctx, http.StatusBadRequest, errors.New("Lỗi"))
				return
			}

			type tmp struct {
				Returnterm time.Time `json:"returnterm"`
			}
			var t tmp
			_, err = c.DB.Query(&t, `SELECT returnterm FROM library.borrows WHERE id = ?`, returnBooks.BorrowId)
			if err != nil {
				log.Print(err)
				model.NewError(ctx, http.StatusBadRequest, errors.New("Lỗi ngày trả sách"))
				return
			}

			var book model.Book
			_, err = c.DB.Query(&book, `SELECT * FROM library.books WHERE id = ?`, item.BookId)
			if err != nil {
				model.NewError(ctx, http.StatusBadRequest, err)
				return
			}

			difference := detail.Returnday.Sub(t.Returnterm)
			if difference <= 0 {
				_, err = c.DB.Exec(`UPDATE library.details SET status = 1 WHERE book_id = ?`, item.BookId)
				if err != nil {
					model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể cập nhật trạng thái trả sách"))
					return
				}

				_, err = c.DB.Exec(`UPDATE library.books SET status = 0 WHERE id = ?`, item.BookId)
				if err != nil {
					model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể cập nhật trạng thái sách"))
					return
				}

				ctx.String(http.StatusOK, "Trả sách " + book.Title + " thành công!")

			} else {
				_, err = c.DB.Exec(`UPDATE library.details SET status = 2 WHERE book_id = ?`, item.BookId)
				if err != nil {
					model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể cập nhật trạng thái trả sách"))
					return
				}

				_, err = c.DB.Exec(`UPDATE library.books SET status = 0 WHERE id = ?`, item.BookId)
				if err != nil {
					model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể cập nhật trạng thái sách"))
					return
				}

				outOfDate := int(difference.Hours() / 24)
				money := float64(outOfDate * 5000)

				var fines model.Fines
				fines.Title = book.Title
				fines.OutOfDate = outOfDate
				fines.Money = money

				ctx.JSON(http.StatusOK, &fines)
			}

		} else if item.Status == "Lost" {
			_, err := c.DB.Exec(`UPDATE library.details SET status = 3 WHERE book_id = ?`, item.BookId)
			if err != nil {
				model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể cập nhật trạng thái trả sách"))
				return
			}

			_, err = c.DB.Exec(`UPDATE library.books SET status = 2 WHERE id = ?`, item.BookId)
			if err != nil {
				model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể cập nhật trạng thái sách"))
				return
			}

			var book model.Book
			_, err = c.DB.Query(&book, `SELECT * FROM library.books WHERE id = ?`, item.BookId)
			if err != nil {
				model.NewError(ctx, http.StatusBadRequest, err)
				return
			}

			price := book.Price
			money := price * 2

			var lost model.Lost
			lost.Title = book.Title
			lost.BookLost = "Mất sách"
			lost.Money = money

			ctx.JSON(http.StatusOK, &lost)

		} else {
			model.NewError(ctx, http.StatusBadRequest, errors.New("Lỗi trạng thái trả sách " + item.BookId + "!"))
			return
		}
	}
}

//DELETE BOOK
func (c *Controller) DeleteBookById(ctx *gin.Context) {
	id := ctx.Param("id")

	rs, err := c.DB.Exec(`DELETE FROM library.books WHERE id = ?`, id)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if rs.RowsAffected() == 0 {
		model.NewError(ctx, http.StatusNotFound, errors.New("Không có sách này"))
		return
	}

	ctx.String(http.StatusOK, "Xóa sách thành công")
}

//GET BOOK BY ID
func (c *Controller) GetBookById(ctx *gin.Context) {
	id := ctx.Param("id")

	var book []model.GetBookById
	rs, err := c.DB.Query(&book, `SELECT * FROM library.books WHERE id = ?`, id)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if rs.RowsReturned() == 0 {
		model.NewError(ctx, http.StatusNotFound, errors.New("Không có sách này"))
		return
	}

	ctx.JSON(http.StatusOK, book)
}