package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type server struct {
	mongoDB *mongoDBClient
}

func newServer(mongoDB *mongoDBClient) *server {
	return &server{mongoDB}
}

func (h *server) startServer() {
	http.HandleFunc("/instruments", h.instruments)
	http.HandleFunc("/addinstrument", h.addInstrument)
	http.HandleFunc("/deleteinstrument", h.deleteInstrument)
	http.Handle("/", http.FileServer(http.Dir("./ui")))

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(http.ListenAndServe(":8080", nil))
		}
	}()
}

func (h *server) instruments(w http.ResponseWriter, r *http.Request) {
	instruments, err := h.mongoDB.getAllInstruments()
	if err != nil {
		http.Error(w, "failed to get all instruments", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(instruments)
	if err != nil {
		http.Error(w, "marshal error", http.StatusBadRequest)
		return
	}
	s := string(data)
	io.WriteString(w, s)
}

func (h *server) deleteInstrument(w http.ResponseWriter, r *http.Request) {
	id, ok := r.URL.Query()["id"]
	if !ok || len(id[0]) < 1 {
		http.Error(w, "failed to get id", http.StatusBadRequest)
		return
	}
	ID, err := strconv.Atoi(id[0])
	if err != nil {
		http.Error(w, "parsing failed", http.StatusBadRequest)
		return
	}
	if err := h.mongoDB.deleteInstrument(ID); err != nil {
		http.Error(w, "failed to delete the instrument", http.StatusInternalServerError)
		return
	}
}

func (h *server) addInstrument(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var data instrument
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "marshal error", http.StatusBadRequest)
		return
	}

	if err := h.mongoDB.addInstrument(data); err != nil {
		http.Error(w, "failed to add the instrument", http.StatusInternalServerError)
		return
	}
}
