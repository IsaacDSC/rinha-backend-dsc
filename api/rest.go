package api

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/IsaacDSC/rinha-backend-dsc/api/controllers"
	"github.com/IsaacDSC/rinha-backend-dsc/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartAPI() {
	// defer config.CloseConnectionRedis()
	// defer config.DbClose()
	env, err := filepath.Abs("./.env")
	if err != nil {
		log.Panic(err.Error())
	}
	config.StartEnv(env)
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	controllers.HealthcheckController(router)
	new(controllers.PersonController).Start(router)
	fmt.Println("[ * ] Starting server in port 3000")
	if err := http.ListenAndServe(":3000", router); err == nil {
		log.Fatal(err)
	}
}
