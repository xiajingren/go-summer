package main

import (
	"github.com/xiajingren/go-summer/conf"
	"github.com/xiajingren/go-summer/internal/api"
	"github.com/xiajingren/go-summer/store"
)

func main() {
	store.InitDatabase()
	conf.InitConfig()
	api.Run()
}
