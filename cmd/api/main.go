package main

import (
	"github.com/xiajingren/go-summer/conf"
	"github.com/xiajingren/go-summer/internal/api"
	"github.com/xiajingren/go-summer/store"
)

func main() {
	conf.InitConfig()
	store.InitDatabase()
	api.Run()
}
