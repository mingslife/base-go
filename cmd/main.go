package main

import (
	"fmt"
	"net/http"

	"base-go/pkg/conf"
	"base-go/pkg/models"
	"base-go/pkg/router"
)

func main() {
	cfg := conf.ParseConfig()

	models.InitDB(cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPwd, cfg.DbName)
	models.Seed()

	router := router.NewRouter(cfg)

	http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), router)
}
