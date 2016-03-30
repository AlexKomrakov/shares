package main

import (
	"net/http"
	"flag"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"github.com/stretchr/graceful"
	"time"
	"fmt"
	"github.com/alexkomrakov/shares/src"
	"github.com/alexkomrakov/shares/src/modules"
)

var server_address string
var modules_container map[string] shares.HasStats

func init() {
	modules_container["facebook"] = modules.Facebook{}
	modules_container["vk"] = modules.Vk{}
}

func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Hello world")
}

func main() {
	flag.StringVar(&server_address, "a", "0.0.0.0:8000", "Server address: host:port")
	flag.Parse()

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(Index)

	n := negroni.Classic()
	n.UseHandler(router)
	graceful.Run(server_address, 30 * time.Second, n)
}