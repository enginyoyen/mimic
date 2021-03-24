package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/enginyoyen/mimic/pkg/cmd"
	"github.com/enginyoyen/mimic/pkg/mapping"
	"github.com/gorilla/mux"
)

var mappings *[]mapping.Mapping

func Serve(config *cmd.Config, maps *[]mapping.Mapping) {
	mappings = maps
	r := mux.NewRouter()
	r.HandleFunc(`/{rest:[a-zA-Z0-9=\-\/]+}`, MockHandler)

	addr := config.BindAddress + ":" + strconv.Itoa(config.Port)
	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}

func InitHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You are using mimics\n")
}

func MockHandler(w http.ResponseWriter, r *http.Request) {
	if m, ok := match(r); ok {
		w.WriteHeader(m.Response.Status)
		fmt.Fprintf(w, m.Response.Body)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found\n")
	}
}

func match(r *http.Request) (*mapping.Mapping, bool) {
	for _, v := range *mappings {
		if v.Request.Url == r.URL.Path {
			return &v, true
		}
	}
	return nil, false
}
