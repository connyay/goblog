package main

import (
	. "github.com/connyay/goblog/controllers"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func pong(res http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	res.Write([]byte("pong"))
}

func StartServer() {
	router := httprouter.New()
	router.GET("/ping/", pong)
	router.GET("/posts/", ListPosts)
	router.GET("/posts/:id", ListPost)

	log.Fatal(http.ListenAndServe(":3333", router))
}

func main() {
	StartServer()
}
