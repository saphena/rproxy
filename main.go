package main

import (
	"fmt"
	"log"
	"time"

	"github.com/emersion/go-smtp"
)

const appversion = "RProxy v1.0"

func main() {
	fmt.Printf("%v   Copyright (c) 2023 Bob Stammers\n", appversion)
	be := &Backend{}

	s := smtp.NewServer(be)

	s.Addr = ":1025"
	s.Domain = "localhost"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
