package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"shop/pkg/tgbot"
	"shop/repository"
	"shop/service"
)

func main() {
	tg, err := tgbot.NewTelegramAPI("1561350817:AAH5bkKOgg9MRqAJLV-QTRFzIbrSUnjWoK8", -432234189)
	if err != nil {
		log.Fatal("Unable to init telegram bot")
	}

	db := repository.NewMapDB()

	service := service.NewService(tg, db)
	handler := &shopHandler{
		service: service,
		db:      db,
	}

	router := mux.NewRouter()

	router.HandleFunc("/item", handler.createItemHandler).Methods("POST")
	router.HandleFunc("/item/{id}", handler.getItemHandler).Methods("GET")
	router.HandleFunc("/item/{id}", handler.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/item/{id}", handler.updateItemHandler).Methods("PUT")

	router.HandleFunc("/order", handler.createOrderHandler).Methods("POST")
	router.HandleFunc("/order/{id}", handler.getOrderHandler).Methods("GET")

	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
