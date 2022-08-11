package http

import (
	"fmt"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/adaptor/mysql"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/infra"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase"
	"github.com/gin-gonic/gin"
)

const (
	apiVersion      = "/v1"
	healthCheckRoot = "/health_check"
	usersAPIRoot    = apiVersion + "/users"
	userIDParam     = "user_id"
	roomIDParam     = "room_id"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// TODO:
	// 別ファイルにした方がいいかも
	mySQLConn := infra.NewMySQLConnector()

	userRepository := mysql.NewUserRepository(mySQLConn.Conn)
	userUsecase := usecase.NewUserUsecase(userRepository)

	healthCheckGroup := router.Group(healthCheckRoot)
	{
		relativePath := ""
		healthCheckGroup.GET(relativePath, healthCheck())
	}

	usersGroup := router.Group(usersAPIRoot)
	{
		relativePath := ""
		userHandler := NewUserHandler(userUsecase)

		// 確認用API
		// /users
		usersGroup.GET(relativePath, userHandler.findUsers())
		// /users/{user_id}
		relativePath = fmt.Sprintf("/:%s", userIDParam)
		usersGroup.GET(relativePath, userHandler.findUserByUserID())

		// /users/{user_id}/rooms
		relativePath = fmt.Sprintf("/:%s/rooms", userIDParam)
		usersGroup.GET(relativePath, userHandler.findRooms())

		// /users/{user_id}/rooms/{room_id}
		relativePath = fmt.Sprintf("/:%s/rooms/:%s", userIDParam, roomIDParam)
		usersGroup.GET(relativePath, userHandler.findMessages())
	}

	return router
}
