package main

import (
	"log"

	"github.com/VarisNithiparkorn/cryptoGraph/backend/controllers"
	"github.com/VarisNithiparkorn/cryptoGraph/backend/routers"
)


func main() {

	 router := routers.SetUpRouter()
	 go controllers.RunWebSocketClient()
	 err := router.Run(":8080") 
     if err != nil {
         log.Fatalf("Server failed to start: %v", err)
     }
	
	
}