package controller

import (
	"github.com/go-pg/pg"
	"library-demo/config"
)

type Controller struct {
	// DB instance
	DB *pg.DB

	// Configuration
	Config config.Config
}

func NewController() *Controller {
	var c Controller
	return &c
}
