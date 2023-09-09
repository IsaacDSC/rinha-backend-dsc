package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Environment struct {
	AWS_ACCESS_KEY_ID      string
	AWS_SECRET_ACCESS_KEY  string
	DATABASE_URL           string
	TOKEN_API              string
	SECRET_API             string
	GODEV                  string
	DATABASE_URL_ANALYTICS string
	PORT                   string
	WHOAMI_URL_DEV         string
	WHOAMI_URL_PROD        string
	WHOAMI_BASIC_TOKEN     string
}

var datasetEnv *Environment

func GetEnv() *Environment {
	return datasetEnv
}

func StartEnv(path string) *Environment {
	envs, err := godotenv.Read(path)
	if err != nil {
		log.Fatal(err)
	}
	var (
		port string = "3000"
	)
	if len(envs["PORT"]) > 0 {
		port = envs["PORT"]
	}
	// godev, _ := strconv.ParseBool(envs["GODEV"])
	datasetEnv = &Environment{
		AWS_ACCESS_KEY_ID:      envs["AWS_ACCESS_KEY_ID"],
		AWS_SECRET_ACCESS_KEY:  envs["AWS_SECRET_ACCESS_KEY"],
		DATABASE_URL_ANALYTICS: envs["DATABASE_URL_ANALYTICS"],
		DATABASE_URL:           envs["DATABASE_URL"],
		TOKEN_API:              envs["OLI_TOKEN"],
		SECRET_API:             envs["OLI_SECRET"],
		GODEV:                  envs["GODEV"],
		PORT:                   port,
		WHOAMI_URL_DEV:         envs["WHOAMI_URL_DEV"],
		WHOAMI_URL_PROD:        envs["WHOAMI_URL_PROD"],
		WHOAMI_BASIC_TOKEN:     envs["WHOAMI_BASIC_TOKEN"],
	}
	return datasetEnv
}
