package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"log"
	"net/http"
	"online-shop/model"
	_ "online-shop/model"
	"time"
)
// REGISTER USER
// @Tags admin
// @Description Register User
// @Param user body model.CreateUser true "Thông tin Register"
// @Success 200 {string} string
// @Failure 500 {object} model.HTTPError
// @Router /shop/create-user [post]
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
		log.Println((err.Error()))
		model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể tạo user"))
		return
	}

	ctx.String(http.StatusOK, "Tạo user thành công")
}


// ALL USERS
// @Tags admin
// @Description Lấy danh sách User
// @Success 200 {string} string
// @Failure 500 {object} model.HTTPError
// @Router /shop/get-users [get]
func (c *Controller) GetUsers(ctx *gin.Context) {
	var users []model.GetUsers
	rs, err := c.DB.Query(&users, `SELECT * FROM shop.users`)
	if err != nil {
		log.Println(err)
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	if rs.RowsReturned() == 0 {
		model.NewError(ctx, http.StatusNotFound, errors.New("Không có user"))
		return
	}

	ctx.JSON(http.StatusOK, users)
}


// USER - username
// @Tags admin
// @Description Login User theo username
// @Param user body model.UserLogIn true "Đăng nhập"
// @Success 200 {string} string
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /shop/login-user [post]
func (c *Controller) UserLogIn(ctx *gin.Context) {
	var login model.UserLogIn
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var user model.User
	rs, err := c.DB.Query(&user, `SELECT * FROM shop.users WHERE username = ?`, login.Username)
	if err != nil {
		log.Println(err)
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	log.Println(login)

	if rs.RowsReturned() == 0 {
		model.NewError(ctx, http.StatusNotFound, errors.New("Tài khoản không tồn tại"))
		return
	}

	if login.Password != user.Password {
		model.NewError(ctx, http.StatusNotFound, errors.New("Tài khoản không đúng"))
		return
	}

	ctx.String(http.StatusOK, "Đăng nhập thành công")
}

//----------------

// CREATE PRODUCT
// @Tags admin
// @Description Tạo Product
// @Param user body model.CreateProduct true "Thông tin tạo sản phẩm"
// @Success 200 {string} string
// @Failure 500 {object} model.HTTPError
// @Router /shop/create-product [post]
func (c *Controller) CreateProduct(ctx *gin.Context) {
	var product model.CreateProduct
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var newProduct model.Product
	copier.Copy(&newProduct, &product)
	newProduct.Id = xid.New().String()
	err = c.DB.Insert(&newProduct)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể thêm sản phẩm"))
		return
	}

	ctx.String(http.StatusOK, "Thêm sản phẩm thành công")
}


// UPDATE PRODUCT - ID
// @Tags admin
// @Description Cập nhật Product theo ID
// @Param product body model.UpdateProductById true "Thông tin sản phẩm"
// @Success 200 {string} string
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /shop/update-product/{id} [put]
func (c *Controller) UpdateProductById(ctx *gin.Context) {
	var product model.UpdateProductById
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var updateProduct model.Product
	copier.Copy(&updateProduct, &product)
	err = c.DB.Update(&updateProduct)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể cập nhật sản phẩm"))
		return
	}

	ctx.String(http.StatusOK, "Cập nhật thành công")
}


// SELECT PRODUCTS
// @Tags admin
// @Description Lấy danh sách Products
// @Success 200 {string} string
// @Failure 500 {object} model.HTTPError
// @Router /shop/get-products [get]
func (c *Controller) GetProducts(ctx *gin.Context) {
	var products []model.GetProducts
	rs, err := c.DB.Query(&products, `SELECT * FROM shop.products`)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	if rs.RowsReturned() == 0 {
		model.NewError(ctx, http.StatusNotFound, errors.New("Không có sản phẩm"))
		return
	}

	ctx.JSON(http.StatusOK, products)
}


// SELECT PRODUCT - ID
// @Tags admin
// @Description Lấy thông tin Product theo ID
// @Param product body model.GetProductById true "Sản phẩm"
// @Success 200 {string} string
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /shop/get-product/{id} [get]
func (c *Controller) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	var products []model.GetProductById
	rs, err := c.DB.Query(&products, `SELECT * FROM shop.products WHERE id = ?`, id)
	if err != nil {
		model.NewError(ctx, http.StatusNotFound, err)
		return
	}

	if rs.RowsReturned() == 0 {
		model.NewError(ctx, http.StatusNotFound, errors.New("Không có sản phẩm"))
		return
	}

	ctx.JSON(http.StatusOK, products)
}


// DELETE PRODUCT - ID
// @Tags admin
// @Description Xóa Product theo ID
// @Success 200 {string} string
// @Failure 404 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Router /shop/delete-product/{id} [delete]
func (c *Controller) DeleteProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	rs, err := c.DB.Exec(`DELETE FROM shop.products WHERE id = ?`, id)
	if err != nil {
		model.NewError(ctx, http.StatusNotFound, err)
		return
	}

	if rs.RowsAffected() == 0 {
		model.NewError(ctx, http.StatusNotFound, errors.New("Không có sản phẩm"))
		return
	}

	ctx.String(http.StatusOK, "Xóa sản phẩm thành công")
}

// CREATE CART
func (c *Controller) CreateCart(ctx *gin.Context) {
	var bill model.Bill
	err := ctx.ShouldBindJSON(&bill)
	cartID := xid.New().String()
	var total float64

	for _, item := range bill.Products {
		log.Println(item)
		var price float64

		rs, err := c.DB.Query(&price, `SELECT Price FROM shop.products WHERE id = ?`, item.ProductId)
		if err != nil {
			model.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		if rs.RowsReturned() == 0 {
			model.NewError(ctx, http.StatusNotFound, errors.New("Không có sản phẩm này"))
			return
		}

		// Check số lượng sản phẩm trong bảng product
		var quantity int16
		_, err = c.DB.Query(&quantity, `SELECT Quantity From shop.products Where id = ?`, item.ProductId)
		if err != nil {
			model.NewError(ctx, http.StatusBadRequest, err)
			return
		}

		if quantity < item.Quantity {
			model.NewError(ctx, http.StatusNotFound, errors.New("Sản phẩm không đủ số lượng"))
			return
		}

		var detail model.Detail
		detail.CartId = cartID
		detail.ProductId = item.ProductId
		detail.Quantity = item.Quantity
		detail.Total = price * float64(detail.Quantity)
		err = c.DB.Insert(&detail)
		if err != nil {
			model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể thêm sản phẩm"))
			return
		}

		total += detail.Total

		// Update số lượng sản phẩm trong bảng product sau khi mua
		_, err = c.DB.Exec(`UPDATE shop.products SET quantity = ? WHERE id = ?`, quantity - item.Quantity, item.ProductId)
		if err != nil {
			model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể cập nhật sản phẩm"))
			return
		}
	}

	var cart model.Cart
	cart.Id = cartID
	cart.UserId = bill.UserId
	cart.Total = total
	cart.Time = time.Now()
	err = c.DB.Insert(&cart)
	log.Println(cart)
	if err != nil {
		model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể thêm cart"))
		log.Println(err)
		return
	}

	ctx.String(http.StatusOK, "Thêm cart thành công")
}

// GET PURCHASE HISTORY
func (c *Controller) GetPurchaseHistory(ctx *gin.Context) {
	id := ctx.Param("id")

	var history []*model.GetPurchaseHistory

	//_, err := c.DB.Query(&history, `SELECT shop.carts.id as cart_id, username, total, time
	//									  FROM shop.carts INNER JOIN shop.users
	//									  ON shop.carts.user_id = shop.users.id
	//									  WHERE shop.users.id = ?`, id)

	rs, err := c.DB.Query(&history,`
		SELECT shop.carts.id as cart_id, username, total, time,
		(
			SELECT json_agg(d)
			FROM (
				SELECT product_name, manufacturer, price, shop.details.quantity, shop.details.total
				FROM shop.details INNER JOIN shop.products
				ON shop.details.product_id = shop.products.id
				WHERE shop.details.cart_id = carts.id
			) as d
		) as detail
		FROM shop.carts INNER JOIN shop.users
		ON shop.carts.user_id = shop.users.id
		WHERE shop.users.id = ?`, id)

	if err != nil {
		log.Println(err)
		model.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	if rs.RowsReturned() == 0 {
		model.NewError(ctx, http.StatusNotFound, errors.New("Không có lịch sử mua hàng với user " + id + " này!"))
		return
	}

	//for _, item := range history {
	//	_, err := c.DB.Query(&item.Detail, `SELECT product_name, manufacturer, price, shop.details.quantity, shop.details.total
	//									   FROM shop.details INNER JOIN shop.products
	//									   ON shop.details.product_id = shop.products.id
	//									   WHERE shop.details.cart_id = ?`, item.CartId)
	//
	//	if err != nil {
	//		model.NewError(ctx, http.StatusNotFound, errors.New("Không có lịch sử mua hàng"))
	//		log.Println(err)
	//		return
	//	}
	//}

	ctx.JSON(http.StatusOK, history)
}