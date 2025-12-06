package main

import (
	"log"

	"github.com/VH288/music-catalogue/internal/configs"
	membershipsHandler "github.com/VH288/music-catalogue/internal/handler/memberships"
	"github.com/VH288/music-catalogue/internal/models/memberships"
	membershipsRepo "github.com/VH288/music-catalogue/internal/repository/memberships"
	membershipsSvc "github.com/VH288/music-catalogue/internal/service/memberships"
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

	db.AutoMigrate(&memberships.User{})

	membershipRepo := membershipsRepo.NewRepository(db)

	membershipSvc := membershipsSvc.NewService(cfg, membershipRepo)

	r := gin.Default()

	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoutes()

	r.Run(cfg.Service.Port)
	_ = db
}
