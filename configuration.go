package main

import (
	"log"

	"github.com/tucnak/store"
)

type KansasConfiguration struct {
	users   []UserConfiguration
	domains []DomainConfiguration
}

type UserConfiguration struct {
	username string
	password string
}

type DomainConfiguration struct {
	domain string
	target string
}

var configuration KansasConfiguration

func loadConfiguration() {
	store.Init("coroiu/kansas")
	if err := store.Load("config.yaml", &configuration); err != nil {
		log.Println("failed to load configuration:", err)
	}
}

func saveConfiguration() {
	if err := store.Save("config.yaml", &configuration); err != nil {
		log.Println("failed to save configuration:", err)
	}
}
