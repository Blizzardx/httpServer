package Server

import (
	"net/http"
	"sync"
	"log"
)


type  httpRequestHandler struct{
	HandlerPool sync.Map
}
func (self *httpRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request){

	log.Println("recieve request " + request.URL.Path)

	handler,ok := self.HandlerPool.Load(request.URL.Path)
	if !ok{
		// return 404
		http.NotFound(writer, request)
		return
	}
	handlerFunc := handler.(func(r http.ResponseWriter,r2 *http.Request))
	if nil == handler{
		// return 502
		http.Redirect(writer, request, request.URL.String(), 502)
		return
	}
	handlerFunc(writer,request)
}
func (self * httpRequestHandler) HandleFunc(path string,handlerFunc func(r http.ResponseWriter,r2 *http.Request)){
	_,ok := self.HandlerPool.Load(path)
	if ok{
		// return 404
		log.Println("already register path " + path)
		return
	}
	self.HandlerPool.Store(path, handlerFunc)
}
