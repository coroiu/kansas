package main

import (
	"log"

	"github.com/tucnak/store"
)

type KansasConfiguration struct {
	Users                 []UserConfiguration   `json:"users"`
	Domains               []DomainConfiguration `json:"domains"`
	DynamicPortForwarding bool                  `json:"dynamicPortForwarding"`
}

type UserConfiguration struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DomainConfiguration struct {
	Domain                 string `json:"domain"`
	Target                 string `json:"target"`
	Port                   int    `json:"port"`
	RequiresAuthentication bool   `json:"requiresAuthentication"`
}

var configuration KansasConfiguration

func initStore() {
	store.Init("coroiu/kansas")
}

func loadConfiguration() {
	if err := store.Load("config.json", &configuration); err != nil {
		log.Println("failed to load configuration:", err)
	}
}

func saveConfiguration() {
	if err := store.Save("config.json", &configuration); err != nil {
		log.Println("failed to save configuration:", err)
	}
}
