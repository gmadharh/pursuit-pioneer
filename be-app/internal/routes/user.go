package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gmadharh/hack-united/be-app/internal/handlers"
)

type UserRouter struct {
	UserHandler *handlers.UserHandler
}

func NewUserRouter(userHandler *handlers.UserHandler) *UserRouter {
	return &UserRouter{UserHandler: userHandler}
}

func (userRouter *UserRouter) InitializeRouter(route *gin.Engine) {
	route.GET("/users", userRouter.UserHandler.GetAllUsers)
	route.GET("/users/:id", userRouter.UserHandler.GetUserByID)
	route.GET("/users/email", userRouter.UserHandler.GetUserByEmail)
	route.POST("/users/new", userRouter.UserHandler.CreateUser)
	route.POST("users/login", userRouter.UserHandler.AuthenticateUser)
	route.PUT("/users/:id", userRouter.UserHandler.UpdateUser)
}
