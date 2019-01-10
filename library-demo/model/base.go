package model

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func ConnectDb(user string, password string, database string, address string) (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr:     address,
	})

	return db
}

func MigrationDb(db *pg.DB, schema string) error {
	// Tạo schema theo tên service
	_, err := db.Exec("CREATE SCHEMA IF NOT EXISTS " + schema + ";")
	if err != nil {
		return err
	}

	// Tạo bảng
	//--Author Table
	var author Author
	err = createTable(&author, db)
	if err != nil {
		return err
	}

	//--Type Table
	var types Type
	err = createTable(&types, db)
	if err != nil {
		return err
	}

	//--Publisher Table
	var publisher Publisher
	err = createTable(&publisher, db)
	if err != nil {
		return err
	}

	//--Book Table
	var book Book
	err = createTable(&book, db)
	if err != nil {
		return err
	}

	//--User Table
	var user User
	err = createTable(&user, db)
	if err != nil {
		return err
	}

	//--Borrow Table
	var borrow Borrow
	err = createTable(&borrow, db)
	if err != nil {
		return err
	}

	//--Detail Table
	var detail Detail
	err = createTable(&detail, db)
	if err != nil {
		return err
	}

	return nil
}

func createTable(model interface{}, db *pg.DB) error {
	err := db.CreateTable(model, &orm.CreateTableOptions{
		Temp:          false,
		FKConstraints: true,
		IfNotExists:   true,
	})

	return err
}
