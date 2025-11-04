package main

import (
	"log"

	"github.com/VarisNithiparkorn/cryptoGraph/config"
	"github.com/VarisNithiparkorn/cryptoGraph/routers"
)


func main() {
	port := config.LoadEnv().Port
	 router := routers.SetUpRouter()
	 err := router.Run(":"+ port) 
     if err != nil {
         log.Fatalf("Server failed to start: %v", err)
     }
	
	
}