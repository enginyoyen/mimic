package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/enginyoyen/mimic/pkg/cmd"
	"github.com/enginyoyen/mimic/pkg/mapping"
	"github.com/gorilla/mux"
)

var mappings *[]mapping.Mapping

func Serve(config *cmd.Config, maps *[]mapping.Mapping) {
	mappings = maps
	r := mux.NewRouter()
	r.HandleFunc("/mimic", InitHandler)

	for _, v := range *mappings {
		m := v
		route := r.HandleFunc(m.Request.Url, func(w http.ResponseWriter, r *http.Request) {
			// Set headers in the
			for header, value := range m.Response.Headers {
				w.Header().Set(header, value)
			}
			w.WriteHeader(m.Response.Status)
			body := string(m.Response.Body)
			body = strings.TrimLeft(body, "\"")
			body = strings.TrimRight(body, "\"")
			fmt.Fprintf(w, "%v", body)
			if m.Response.WithDelay > 0 {
				time.Sleep(time.Millisecond * m.Response.WithDelay)
			}
		}).Methods(m.Request.Method)
		//Set request headers
		for header, value := range m.Request.Headers {
			route.Headers(header, value)
		}
	}

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
