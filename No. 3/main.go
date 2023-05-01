package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Identity struct {
	Name    string
	Nim     int
	Address string
}

var p []Identity

func getMethod(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func postMethod(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newPerson Identity
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &newPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p = append(p, newPerson)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "New person added: %s", newPerson.Name)
}

func main() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/get", getMethod)
	mux.HandleFunc("/post", postMethod)

	http.ListenAndServe(":8080", mux)
}
