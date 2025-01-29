package main

import (
	"fmt"
	"path/filepath"

	"hiyoko-fiber/configs"
	"hiyoko-fiber/internal/infrastructure/database"
	"hiyoko-fiber/internal/interactor"
	"hiyoko-fiber/internal/presentation/http/app/middleware"
	"hiyoko-fiber/internal/presentation/http/app/router"
	"hiyoko-fiber/pkg/logging/file"
	"hiyoko-fiber/utils"

	"github.com/gofiber/fiber/v2"
)

const (
	envRoot = "./cmd/app"
	logDir  = "./log/app"
)

func init() {
	log.SetLogDir(logDir)
	log.Initialize()

	envFile := utils.EnvFile(filepath.Join(envRoot, ".env"))
	err := envFile.LoadEnv()
	if err != nil {
		log.Fatal("Failed to load environment variables", "error", err)
	}
	err = envFile.CheckMustEnv(configs.GetMustEnvItemsForApp())
	if err != nil {
		log.Fatal("Must envs were not set", "error", err)
	}

	utils.LoadTimezone(utils.Env("TZ").GetString())
}

func main() {
	f := fiber.New(configs.NewServerConf())
	entClient, err := database.NewMySqlConnect(configs.NewMySqlConf())
	if err != nil {
		log.Fatal("Failed to create dbclient", "error", err)
	}
	defer func(entClient *database.MysqlEntClient) {
		err := entClient.Close()
		if err != nil {
			log.Fatal("Failed to close dbclient", "error", err)
		}
	}(entClient)

	i := interactor.NewInteractor(entClient)
	h := i.NewAppHandler()

	middleware.NewMiddleware(f)
	router.NewRouter(f, h)
	if err := f.Listen(fmt.Sprintf(":%d", utils.Env("SERVER_PORT").GetInt(8080))); err != nil {
		log.Fatal("Failed to start server", "error", err)
	}

	log.Fatal(fmt.Sprintf("Server started on port: %d", utils.Env("SERVER_PORT").GetInt(8080)))
}
