package main

import (
	"fmt"
	"log"

	"github.com/tucnak/store"
)

type KansasConfiguration struct {
	Users                 []UserConfiguration  `json:"users"`
	Domains               DomainConfigurations `json:"domains"`
	DynamicPortForwarding bool                 `json:"dynamicPortForwarding"`
}

type UserConfiguration struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type DomainConfigurations []DomainConfiguration
type DomainConfiguration struct {
	Id                     string `json:"id"`
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

	fmt.Println(configuration)
}

func saveConfiguration() {
	if err := store.Save("config.json", &configuration); err != nil {
		log.Println("failed to save configuration:", err)
	}
}

func (domains DomainConfigurations) findIndex(id string) int {
	for i, domain := range domains {
		if domain.Id == id {
			return i
		}
	}
	return -1
}
