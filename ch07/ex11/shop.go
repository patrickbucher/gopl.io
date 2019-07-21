package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var db = database{"shoes": 50, "socks": 5}

func main() {
	http.HandleFunc("/list", http.HandlerFunc(db.list))
	http.HandleFunc("/create", http.HandlerFunc(db.create))
	http.HandleFunc("/read", http.HandlerFunc(db.create))
	http.HandleFunc("/update", http.HandlerFunc(db.create))
	http.HandleFunc("/delete", http.HandlerFunc(db.create))
	http.HandleFunc("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fail(405, w)
		return
	}
	item := req.FormValue("item")
	if item == "" {
		http.Error(w, "field 'item' missing", 400)
		return
	}
	if _, exists := db[item]; exists {
		msg := fmt.Sprintf("the item '%s' already exists", item)
		http.Error(w, msg, 409)
		return
	}
	priceStr := req.FormValue("price")
	if priceStr == "" {
		http.Error(w, "field 'price' missing", 400)
		return
	}
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		msg := fmt.Sprintf("parsing price '%s' as float: %v", priceStr, err)
		http.Error(w, msg, 400)
	}
	db[item] = dollars(price)
	w.WriteHeader(http.StatusCreated)
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	fail(501, w)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	fail(501, w)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	fail(501, w)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "n such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func fail(errorCode int, w http.ResponseWriter) {
	message := http.StatusText(errorCode)
	if message == "" {
		message = fmt.Sprintf("unknown error: %d", errorCode)
	} else {
		message = fmt.Sprintf("%d %s", errorCode, message)
	}
	http.Error(w, message, errorCode)
}
