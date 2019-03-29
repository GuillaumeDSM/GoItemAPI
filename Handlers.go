package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func ItemIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(getRandomItem().Text); err != nil {
		panic(err)
	}
}

func ItemShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ItemId := vars["ItemId"]
	intId, _ := strconv.Atoi(ItemId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(selectItem(intId)); err != nil {
		panic(err)
	}
}

func ItemApprove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ItemId := vars["ItemId"]
	intId, _ := strconv.Atoi(ItemId)
	IncItemRating(intId)
	if err := json.NewEncoder(w).Encode(selectItem(intId)); err != nil {
		panic(err)
	}
}

func ItemCreate(w http.ResponseWriter, r *http.Request) {
	var item Item
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 99999))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &item); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	newItem := addItem(item)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newItem); err != nil {
		panic(err)
	}
}
