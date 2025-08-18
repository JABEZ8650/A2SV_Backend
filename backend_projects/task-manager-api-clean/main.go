package main

import (
	"task-manager-api-clean/api/router"
	"task-manager-api-clean/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	env, _ := config.Load()
	db, _ := config.GetClient(env.DatabaseURL, env.DatabaseName)
	router.Setup(env, db, r)
	r.Run("localhost:" + env.Port)
}
