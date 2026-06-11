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
	userController := handler.NewUserController(server, userService)

	userRouter.POST("/signup", userController.RegisterUser)
	userRouter.POST("/signin", userController.LoginUser)
	userRouter.GET("/info", userController.GetUserInfo)

}
