package main

import (
	"fmt"
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/db"
	"os"
)

func init() {
	pos := db.ConnectPostgres(os.Getenv("DATABASE_DSN"))
	fmt.Println(pos)
}
