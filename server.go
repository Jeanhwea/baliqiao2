package main

import (
	"github.com/Jeanhwea/baliqiao2/bootstrap"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	bootstrap.RootApp.Execute()
}
