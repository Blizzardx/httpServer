package Server

import (
	"sync"
	"time"
	"log"
	"net/http"
	"github.com/Blizzardx/httpServer/Common"
	"strconv"
)

var mutix sync.Mutex
var handlerPool = map[int32]*httpRequestHandler{}
var httpServerList []*http.Server

type HttpStartConfig struct{
	Port string
	GroupId int32
}

//这个方法是线程安全的 double check
func getHandler(groupId int32)*httpRequestHandler{
	if v,ok := handlerPool[groupId];ok{
		return v
	}
	mutix.Lock()
	defer mutix.Unlock()

	if v,ok := handlerPool[groupId];ok{
		return v
	}
	newInstance := &httpRequestHandler{
		HandlerPool: sync.Map{},
	}

	handlerPool[groupId] = newInstance

	return newInstance
}

func RegisterHandler(groupId int32,path string,handlerFunc func(r http.ResponseWriter,r2 *http.Request)){
	handler := getHandler(groupId)
	handler.HandleFunc(path,handlerFunc)
}

func StartServer(httpConfig []*HttpStartConfig){
	for groupId,handler := range handlerPool{

		var configInfo *HttpStartConfig = nil
		for _,cfgElem := range httpConfig{
			if cfgElem.GroupId == groupId{
				configInfo  = cfgElem
				break
			}
		}
		if nil == configInfo{
			log.Println("error on start server,config not found, group id " + strconv.Itoa(int(groupId)))

			continue
		}
		server := startServer(configInfo.Port,handler)
		httpServerList = append(httpServerList,server)
	}
}
func StopServer(groupId int32){
	//for _,server := range httpServerList{
	//
	//}
}
func StopAllServer(){
	//for _,server := range httpServerList{
	//
	//}
}
func startServer(port string,handler *httpRequestHandler)*http.Server{
	httpServer := &http.Server{
		Addr:           ":"+port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        handler,
	}

	go common.SafeCall(func() {
		err := httpServer.ListenAndServe()
		log.Fatal(err)
	})
	return httpServer
}