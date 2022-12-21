package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DevBeast3800/GolangAPI/home"
	"github.com/DevBeast3800/GolangAPI/items"
	"github.com/DevBeast3800/GolangAPI/promotions"
	"github.com/DevBeast3800/GolangAPI/users"
	"github.com/DevBeast3800/GolangAPI/orders"
	"github.com/DevBeast3800/GolangAPI/itemsorders"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Here we are instantiating the gorilla/mux router
	r := mux.NewRouter()

	// On the default page we will simply serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./views/")))

	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// healtcheck routes
	r.HandleFunc("/healthCheck", home.HealthCheck).Methods("GET")

	// items routes
	r.HandleFunc("/items", items.ItemCreate).Methods("POST")
	r.HandleFunc("/items", items.ItemIndex).Methods("GET")
	r.HandleFunc("/items/{id}", items.ItemFind).Methods("GET")
	r.HandleFunc("/items/{id}", items.ItemUpdate).Methods("PUT")
	r.HandleFunc("/items/{id}", items.ItemDelete).Methods("DELETE")
	
	// promotions routes
	r.HandleFunc("/promotions", promotions.PromotionCreate).Methods("POST")
	r.HandleFunc("/promotions", promotions.PromotionIndex).Methods("GET")
	r.HandleFunc("/promotions/{id}", promotions.PromotionFind).Methods("GET")
	r.HandleFunc("/promotions/{id}", promotions.PromotionUpdate).Methods("PUT")
	r.HandleFunc("/promotions/{id}", promotions.PromotionDelete).Methods("DELETE")
	
	// users routes
	r.HandleFunc("/users", users.UserCreate).Methods("POST")
	r.HandleFunc("/users", users.UserIndex).Methods("GET")
	r.HandleFunc("/users/{id}", users.UserFind).Methods("GET")
	r.HandleFunc("/users/{id}", users.UserUpdate).Methods("PUT")
	r.HandleFunc("/users/{id}", users.UserDelete).Methods("DELETE")
	
	// orders routes
	r.HandleFunc("/orders", orders.OrderCreate).Methods("POST")
	r.HandleFunc("/orders", orders.OrderIndex).Methods("GET")
	r.HandleFunc("/orders/{id}", orders.OrderFind).Methods("GET")
	r.HandleFunc("/orders/{id}", orders.OrderUpdate).Methods("PUT")
	r.HandleFunc("/orders/{id}", orders.OrderDelete).Methods("DELETE")
	
	// itemsorders routes
	r.HandleFunc("/itemsorders", itemsorders.ItemsorderCreate).Methods("POST")
	r.HandleFunc("/itemsorders", itemsorders.ItemsorderIndex).Methods("GET")
	r.HandleFunc("/itemsorders/{id}", itemsorders.ItemsorderFind).Methods("GET")
	r.HandleFunc("/itemsorders/{id}", itemsorders.ItemsorderUpdate).Methods("PUT")
	r.HandleFunc("/itemsorders/{id}", itemsorders.ItemsorderDelete).Methods("DELETE")

	fmt.Println("Cerberus ready to GO!")
	// implement logging middleware, initiate server in port, log fatal in case of error
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r)))

}
