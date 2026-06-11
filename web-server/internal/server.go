package internal

import (
	"fmt"
	"log"
	"os"

	database "github.com/Sam-Frost/db"
	db "github.com/Sam-Frost/db/generated"
	"github.com/Sam-Frost/web-server/internal/middleware"
	"github.com/Sam-Frost/web-server/internal/route"
	"github.com/Sam-Frost/web-server/internal/util"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	util.LoadEnv()

	envVariables, err := util.GetEnvVariables()
	if err != nil {
		log.Fatalf("Trying to read env variables : %v", err)
	}

	dbConnectionPool := database.CreateConnectionPool(envVariables.DatabaseURL)
	defer dbConnectionPool.Close()

	router := gin.Default()
	router.Use(middleware.ErrorHandler())

	query := db.New(dbConnectionPool)

	server := &util.Server{
		Router: router,
		DB:     dbConnectionPool,
		Query:  query,
	}

	route.UserRouter(server)
	// route.OrderRouter(ginEngine, appContext)
	// route.AffiliateRouter(ginEngine, appContext)

	router.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
