package main

import (
	"github.com/keisuke-333/go_practice/db"
	"github.com/keisuke-333/go_practice/server"
)

func main() {
	db.Init()
	server.Init()
}
