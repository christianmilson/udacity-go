package controllers

import (
	"net/http"

	"github.com/christianmilson/udacity-go/db"
)

type HomeController struct {
	db *db.Database
}

func NewHomeController(db *db.Database) *HomeController {
	return &HomeController{
		db: db,
	}
}

func (home *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/index.html")
}
