package main

import (
	"flag"
	"gb_go_arch/lesson-2/shop_new/notification"
	"gb_go_arch/lesson-2/shop_new/repository"
	"gb_go_arch/lesson-2/shop_new/service"
	"log"
	"net/http"

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

	from := "service@list.ru"
	password := "passd"
	host := "smtp.list.ru"
	port := "587"

	smtpBot := notification.NewSMTPBot(from, host, port, password)

	rep := repository.NewMapDB()
	service := service.NewService(rep, notif, *smtpBot)
	s := &server{
		service: service,
		rep:     rep,
		smtpBot: *smtpBot,
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
