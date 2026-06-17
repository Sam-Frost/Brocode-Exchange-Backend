package route

import (
	"github.com/Sam-Frost/web-server/internal/handler"
	"github.com/Sam-Frost/web-server/internal/repo"
	"github.com/Sam-Frost/web-server/internal/service"
	"github.com/Sam-Frost/web-server/internal/util"
)

func UserRouter(server *util.Server) {
	userRouter := server.Router.Group("/api/v1/user")

	userRepo := repo.NewUserRepo(server)
	userService := service.NewUserService(server, userRepo)
	userHandler := handler.NewUserHandler(server, userService)

	userRouter.POST("/signup", userHandler.RegisterUser)
	userRouter.POST("/signin", userHandler.LoginUser)
	userRouter.GET("/info", userHandler.GetUserInfo)

}
