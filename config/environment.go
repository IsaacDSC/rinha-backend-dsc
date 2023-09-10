package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Environment struct {
	DATABASE_URL string
	GODEV        string
	PORT         string
	REDIS_URL    string
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
		DATABASE_URL: envs["DATABASE_URL"],
		REDIS_URL:    envs["REDIS_URL"],
		GODEV:        envs["GODEV"],
		PORT:         port,
	}
	return datasetEnv
}
