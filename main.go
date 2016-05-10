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
	"strings"
	"encoding/json"
)

var server_address string
var modules_container map[string] shares.HasStats

func init() {
	modules_container                = make(map[string] shares.HasStats)
	modules_container["facebook"]    = modules.Facebook{&shares.Stats{}}
	modules_container["vk"]          = modules.Vk{&shares.Stats{}}
	modules_container["ok"]          = modules.Ok{&shares.Stats{}}
	modules_container["google_plus"] = modules.Gp{&shares.Stats{}}
	modules_container["my_mail"]     = modules.Mm{&shares.Stats{}}
}

func GetStats(url string) map[string] shares.HasStats {
	for _, module := range modules_container {
		module.SetUrl(url)
		module.CalculateShares()
	}

	return modules_container
}

func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	url := req.URL.Query().Get("url")

	services := req.URL.Query().Get("services")
	services_list := []string{}
	if (len(services) == 0) {
		for k, _ := range modules_container {
			services_list = append(services_list, k)
		}
	} else {
		services_list = strings.Split(services, ",")
	}
	fmt.Println(services_list)

	response := make(map[string]int)
	for _, service := range services_list {
		modules_container[service].SetUrl(url)
		modules_container[service].CalculateShares()
		response[service] = modules_container[service].GetShares()
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