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
	//--User Table
	var user User
	err = createTable(&user, db)
	if err != nil {
		return err
	}

	//--Product Table
	var product Product
	err = createTable(&product, db)
	if err != nil {
		return err
	}

	//--Cart Table
	var cart Cart
	err = createTable(&cart, db)
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
