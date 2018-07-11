package server

import (
	"net/http"
	"strconv"
	//"github.com/goinggo/mapstructure"
	"log"
)

type GoKVServer struct {
	ip     string
	port   int
}

func (server *GoKVServer) start() {
	//server.register()
	//http.HandleFunc("/", server.dispatcher)
	log.Println("server ready to start on ", server.ip, ":", server.port)
	http.ListenAndServe(server.ip +":"+ strconv.Itoa(server.port), newHttpMux())

}

//func (this *GoKVServer) dispatcher(w http.ResponseWriter, req *http.Request) {
//	header := w.Header()
//	header.Add("Content-Type", "application/json;charset=utf-8")
//
//	result := this.exec(req)
//	template := model.HttpResponseTemplate{Code: 200, Msg: "ok", Data: result}
//
//	fmt.Println("request url=[", req.URL, "]\trespond=", template)
//	enc := json.NewEncoder(w)
//	enc.Encode(template)
//}

func Server(port int) *GoKVServer {
	server := &GoKVServer{
		ip:"127.0.0.1",
		port:port,
	}
	server.start()
	return server
}

//func (this *GoKVServer) AddRouter(url string, handler handler.Op, method string) {
//	this.router.Register(url, handler, method)
//}

//func (this *GoKVServer) exec(req *http.Request) interface{} {
//	// get method from router to invoke
//
//	// map -> json string
//	//json.Marshal(req.Form)
//
//	map1 := make(map[string]interface{})
//	map1["1"] = "hello"
//	map1["2"] = "world"
//	//return []byte
//	//str, err := json.Marshal(map1)
//	mapstructure.Decode(map1, &Router{})
//
//	req.ParseForm()
//	//req.PostForm
//
//	return this.router.handlerMap[req.URL.String()].Invoke(req.PostForm)
//}
//
//func (server *GoKVServer) register() {
//	server.router.Register("/set", &handler.Op{}, "Set")
//	server.router.Register("/get", &handler.Op{}, "Get")
//	server.router.Register("/exist", &handler.Op{}, "Exist")
//
//}
