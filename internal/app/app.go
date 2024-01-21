package app

import (
	"log"

	"github.com/blog/config"
	v1 "github.com/blog/internal/controller/http"
	Db "github.com/blog/internal/storage"
	"github.com/blog/pkg/db"

	"github.com/blog/pkg/logger"
)

func Run(cfg config.Config) {
	l := logger.New("debug", "blog")

	connDb, err := db.ConnectToDb(cfg)
	if err != nil {
		log.Println("Error connect db: ", err.Error())
	}

	strg := Db.NewStoragePg(connDb)

	apiServer := v1.New(&v1.Options{
		Cfg:     cfg,
		Storage: strg,
		Log:     l,
	})

	if err := apiServer.Run(":" + cfg.HttpPort); err != nil {
		log.Fatalf("failed to run server: %v", err.Error())
	}
}
