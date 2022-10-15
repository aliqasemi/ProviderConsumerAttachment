package main

import (
	"flag"
	"fmt"
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/db"
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/handlers/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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
	e := echo.New()
	if err = routes.SetRoutes(e); err != nil {
		panic("router has problem")
	}
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
