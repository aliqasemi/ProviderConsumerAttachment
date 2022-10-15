package main

import (
	"fmt"
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/db"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	DB := db.ConnectPostgres(os.Getenv("DATABASE_DSN"))
	fmt.Println(DB)
}
