package main

import (
	"flag"
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
	if err = db.ConnectPostgres(os.Getenv("DATABASE_DSN")); err == nil {
		fmt.Println(db.GetDataBase())
	}
	str1 := flag.String("m", "", "set migration to 'all' value")
	flag.Parse()
	if *str1 == "all" {
		db.MigratePostgres()
	}
}
