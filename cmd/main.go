package main

import (
	"log"

	"github.com/VH288/music-catalogue/internal/configs"
	"github.com/VH288/music-catalogue/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	var cfg *configs.Config

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiai config", err)
	}

	cfg = configs.Get()
	log.Printf("Configs: %+v", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to databasei, err: &+v\n", err)
	}

	r := gin.Default()
	r.Run(cfg.Service.Port)
	_ = db
}
