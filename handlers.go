package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Status  string `json:"status"`
	Action  string `json:"action"`
	Message string `json:"message"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	val, _ := client.Get("id").Result()
	var response Response
	if val == "" {
		response = Response{Status: "200", Action: "Get", Message: "No ids found"}
	} else {
		response = Response{Status: "200", Action: "Get", Message: "Id: " + val}
	}

	b, _ := json.MarshalIndent(response, "", "  ")
	fmt.Fprintf(w, string(b))
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	client.Set("id", vars["id"], 0)

	var response Response
	response = Response{Status: "200", Action: "Create", Message: "Id created with value: " + vars["id"]}

	b, _ := json.MarshalIndent(response, "", "  ")
	fmt.Fprintf(w, string(b))
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	client.Set("id", vars["id"], 0)

	var response Response
	response = Response{Status: "200", Action: "Update", Message: "Id updated with value: " + vars["id"]}

	b, _ := json.MarshalIndent(response, "", "  ")
	fmt.Fprintf(w, string(b))
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	var response Response
	val, _ := client.Get("id").Result()
	if val == "" {
		response = Response{Status: "500", Action: "Delete", Message: "Id with value: " + vars["id"] + " not found"}
	} else {
		client.Del("id", vars["id"])
		response = Response{Status: "200", Action: "Update", Message: "Id deleted with value: " + vars["id"]}
	}

	b, _ := json.MarshalIndent(response, "", "  ")
	fmt.Fprintf(w, string(b))
}
