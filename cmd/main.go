package main

import (
	"log"
	"net/http"

	"github.com/VH288/music-catalogue/internal/configs"
	membershipsHandler "github.com/VH288/music-catalogue/internal/handler/memberships"
	tracksHandler "github.com/VH288/music-catalogue/internal/handler/tracks"
	"github.com/VH288/music-catalogue/internal/models/memberships"
	"github.com/VH288/music-catalogue/internal/models/trackactivities"
	membershipsRepo "github.com/VH288/music-catalogue/internal/repository/memberships"
	"github.com/VH288/music-catalogue/internal/repository/spotify"
	trackactivitiesRepo "github.com/VH288/music-catalogue/internal/repository/trackactivities"
	membershipsSvc "github.com/VH288/music-catalogue/internal/service/memberships"
	"github.com/VH288/music-catalogue/internal/service/tracks"
	"github.com/VH288/music-catalogue/pkg/httpclient"
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
	db.AutoMigrate(&trackactivities.TrackActivity{})

	r := gin.Default()
	httpclient := httpclient.NewClient(&http.Client{})

	membershipRepo := membershipsRepo.NewRepository(db)
	membershipSvc := membershipsSvc.NewService(cfg, membershipRepo)
	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoutes()

	spotifyOutbound := spotify.NewSpotifyOutbound(cfg, httpclient)
	trackActivityRepo := trackactivitiesRepo.NewRepository(db)
	trackSvc := tracks.NewService(spotifyOutbound, trackActivityRepo)
	tracksHandler := tracksHandler.NewHandler(r, trackSvc)
	tracksHandler.RegisterRoutes()

	r.Run(cfg.Service.Port)
	_ = db
}
