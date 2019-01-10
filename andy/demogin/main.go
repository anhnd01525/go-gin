package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"xid"
	"net/http"
	"strconv"
)

func getDemo(c *gin.Context) {
	c.String(http.StatusOK, "Hello Techmaster")
}

func getUser(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "My teacher is %s", name)
}

func getAge(c *gin.Context) {
	born := c.Query("born")
	i, _ := strconv.Atoi(born)
	age := 2018 - i
	c.String(http.StatusOK, "I'm %d years old", age)
}

func addUser(c *gin.Context) {
	log.Println(c.Request.Body)
	var createUser CreateUser
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "Request sai!")
		return
	}

	name := createUser.Name
	born := createUser.Born
	//i, _ := strconv.Atoi(born)
	age := int32(2018 - born)
	id := xid.New().String()

	newUser := User{id, name, age}
	users = append(users, newUser)
	log.Println(users)
	c.String(http.StatusOK, "Thêm user thành công")
}

func updateProfile(c *gin.Context) {
	var updateUser UpdateUser
	id := c.Param("id")
	err := c.ShouldBindJSON(&updateUser)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "Request sai!")
		return
	}

	log.Println(updateUser)

	nameUpd := updateUser.User.Name
	ageUpd := updateUser.User.Age


	for i, _ := range users {
		if users[i].id == id {
			users[i].name = nameUpd
			users[i].age = ageUpd
			log.Println(users[i])
			var rsp  = UpdateUser {
				User:Thuy{
					Name:users[i].name,
					Age: users[i].age,
				},
			}
			c.JSON(http.StatusOK, rsp)
			return
		}
	}
}

func deleteUser (c *gin.Context) {
	id := c.Param("id")

	for i, _ := range users {
		if users[i].id == id {
			users = append(users[:i], users[i+1:]...)
			log.Println(users)
			c.String(http.StatusOK, "Xoa thanh cong")
			return
		}
	}
	c.JSON(404, gin.H {
		"message": "Không tồn tại user",
	})
}

type CreateUser struct {
	Name string `json:"name"`
	Born int32 `json:"born"`
}

type UpdateUser struct {
	User Thuy `json:"user"`
}

type Thuy struct {
	Name string `json:"name"`
	Age int32 `json:"age"`
}

type User struct {
	id string `json:"id"`
	name string `json:"name"`
	age int32 `json:"age"`
}

var users = []User {
	{
		id: "bg1q3c2on8cepnj3b88g",
		name: "Nguyen Van A",
		age: 30,
	},
	{
		id: "bg1q3mqon8ceofhthtn0",
		name: "Tran Thi B",
		age: 21,
	},
	{
		id: "bg1q48aon8cf77ekifbg",
		name: "Mai Tinh C",
		age: 51,
	},
	{
		id: "bg1q4cion8cf7npl7lgg",
		name: "Phung Thanh D",
		age: 9,
	},
}

func main() {
	r := gin.Default()

	//r.GET("/", func(c *gin.Context) {
	//	//c.JSON(200, gin.H {
	//	//	"message": "Hello Techmaster",
	//	//})
	//	c.String(http.StatusOK, "Hello Techmaster")
	//})

	r.GET("/", getDemo)

	r.GET("/introduce/:name", getUser)

	r.GET("/age", getAge)

	r.POST("/api/add-user", addUser)

	r.PUT("/api/update-profile/:id", updateProfile)

	r.DELETE("/api/delete-user/:id", deleteUser)

	r.Run(":8090")
}