package main

import (
	"net/http"
	"flag"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"github.com/stretchr/graceful"
	"time"
	"fmt"
	"shares/modules"
	"strings"
	"encoding/json"
)

var server_address string
var modules_container map[string] modules.HasStats

func init() {
	modules_container                  = make(map[string] modules.HasStats)
	modules_container["facebook"]      = modules.Facebook{&modules.Stats{}}
	modules_container["vk"]            = modules.Vk{&modules.Stats{}}
	modules_container["odnoklassniki"] = modules.Ok{&modules.Stats{}}
	modules_container["google_plus"]   = modules.Gp{&modules.Stats{}}
	modules_container["my_mail"]       = modules.Mm{&modules.Stats{}}
}

func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	url := req.URL.Query().Get("url")
	if url == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	services := req.URL.Query().Get("services")
	services_list := []string{}
	if (len(services) == 0) {
		for k, _ := range modules_container {
			services_list = append(services_list, k)
		}
	} else {
		services_list = strings.Split(services, ",")
	}

	response := make(map[string]int)
	for _, service := range services_list {
		modules_container[service].SetUrl(url)
		modules_container[service].CalculateShares()
		response[service] = modules_container[service].GetShares()
		response["total"] += response[service]
	}


	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Fprint(res, string(b))
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