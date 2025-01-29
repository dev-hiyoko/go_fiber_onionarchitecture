package main

// exec command
// go run ./cmd/cli/main.go -exec test

import (
	"flag"
	"fmt"
	"path/filepath"

	"hiyoko-fiber/configs"
	"hiyoko-fiber/internal/infrastructure/database"
	"hiyoko-fiber/internal/interactor"
	logger "hiyoko-fiber/pkg/logging/file"
	"hiyoko-fiber/utils"
)

const (
	envRoot = "./cmd/cli"
	logDir  = "./log/cli"

	execTest                  = "test"
	execGenJWTSecretKeyForApp = "genJwtSecretKeyForApp"

	errDefaultMsg = "Failed to exec"
	successfulMsg = "Success exec"
)

var (
	databaseConf database.MysqlConf
	exec         *string
)

func init() {
	exec = flag.String("exec", "test", "exec command")
	flag.Parse()

	logger.SetLogDir(logDir)
	logger.Initialize()
	logger.With("exec", exec)

	err := utils.EnvFile(filepath.Join(envRoot, ".env")).LoadEnv()
	if err != nil {
		logger.Fatal("Failed to load environment variables", "error", err)
	}

	databaseConf = configs.NewMySqlConf()
}

func main() {
	entClient, err := database.NewMySqlConnect(databaseConf)
	if err != nil {
		logger.Fatal("Failed to create dbclient", "error", err)
	}
	defer func(entClient *database.MysqlEntClient) {
		err := entClient.Close()
		if err != nil {
			logger.Fatal("Failed to close dbclient", "error", err)
		}
	}(entClient)

	i := interactor.NewInteractor(entClient)
	h := i.NewCliHandler()

	switch *exec {
	case execTest:
		fmt.Println("Exec Test")
	case execGenJWTSecretKeyForApp:
		err := h.GenJWTSecretKeyForApp()
		if err != nil {
			logger.Fatal(errDefaultMsg, "error", err)
			return
		}
	}
	logger.Info(successfulMsg)
}
