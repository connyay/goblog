package controllers

import (
	"github.com/connyay/goblog/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ListPosts(res http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(models.GetPosts())
}

func ListPost(res http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(models.GetPost(ps.ByName("id")))
}
