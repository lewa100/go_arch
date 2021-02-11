package main

import (
	"flag"
	"log"
	"net/http"
	"shop/notification"
	"shop/repository"
	"shop/service"

	"github.com/gorilla/mux"
)

func main() {
	var tokenStr string
	flag.StringVar(&tokenStr, "t", "", "token for telegram api")

	flag.Parse()

	notif, err := notification.NewTelegramBot(tokenStr, 323615875)
	if err != nil {
		log.Fatal(err)
	}

	rep := repository.NewMapDB()
	service := service.NewService(rep, notif)
	s := &server{
		service: service,
		rep:     rep,
	}

	router := mux.NewRouter()

	router.HandleFunc("/items", s.listItemHandler).Methods("GET")
	router.HandleFunc("/items", s.createItemHandler).Methods("POST")
	router.HandleFunc("/items/{id}", s.getItemHandler).Methods("GET")
	router.HandleFunc("/items/{id}", s.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/items/{id}", s.updateItemHandler).Methods("PUT")

	router.HandleFunc("/orders", s.listOrdersHandler).Methods("GET")
	router.HandleFunc("/orders", s.createOrderHandler).Methods("POST")

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	log.Fatal(srv.ListenAndServe())
}
