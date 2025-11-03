package config

import (
	"log" 
	"os"

	"github.com/joho/godotenv"
)


type DBConnection struct { 
	URI  string 
	Port string 
	Database string
}


func LoadEnv() DBConnection { 
	
	
	if err := godotenv.Load(); err != nil { 
		log.Println("WARNING: Could not find .env file. Reading environment variables directly.")
	}

	conn := DBConnection{
		URI:  os.Getenv("MONGO_URI"),
		Port: os.Getenv("PORT"),
		Database:os.Getenv("DATABASE"),
	}

	return conn
}
