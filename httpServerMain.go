package main

import (
"fmt"
"net/http"
	"github.com/Blizzardx/httpServer/Server"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Form)
	fmt.Fprintf(w, r.Form.Get("name1"))
}

func main() {
	//router := httprouter.New()
	//router.GET("/", Index)
	//router.GET("/hello/:name", Hello)
	//
	//go func() {
	//	log.Fatal(http.ListenAndServe(":8080", router))
	//}()
	//log.Fatal(http.ListenAndServe(":8081", router))

	Server.RegisterHandler(0,"/",Index)
	Server.RegisterHandler(1,"/hello/:name",Hello)

	Server.StartServer([]*Server.HttpStartConfig{ &Server.HttpStartConfig{"8080",0},&Server.HttpStartConfig{"8081",1},})
	select {

	}
}
