// add additional handlers so that clients can create, read and update , and delete database entries.
// Fpr example, a request of the form /update?item=socks&price=6 will update the price of an item
// in the inventory and report an error if the item does not exist or if the price is invalid. (Warning:
// this changes introduces concurrent variable update.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32 // BAD

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

// helper functions
func strToDollars(s string) (dollars, error) {
	money, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return dollars(0), err
	}
	return dollars(money), nil
}

// add handlers
func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("%s exists in database", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	if p, err := strconv.ParseFloat(price, 32); err != nil {
		msg := fmt.Sprintf("%s cannot be converted to required datatype", price)
		http.Error(w, msg, http.StatusBadRequest)
		return

	} else {
		db[item] = dollars(p)
		fmt.Fprintf(w, "item %s of price: %s added successfully\n", item, price)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("%s does not exist in the database", item)
		http.Error(w, msg, http.StatusUnprocessableEntity)
		return
	}
	p, err := strToDollars(price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	db[item] = p
	fmt.Fprintf(w, "%s updated to %s in the database", item, price)
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		msg := fmt.Sprintf("%s does not exist in the database", item)
		http.Error(w, msg, http.StatusUnprocessableEntity)
		return
	}
	fmt.Fprintf(w, "The price of %s is %v", item, price)
}

// kk
func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/get", db.update)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
