package main

import (
	"gopher-bank/controller"
	"gopher-bank/repository"
	"gopher-bank/router"
	"gopher-bank/service"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
	userRepository repository.UserRepository = repository.NewUserRepository()
	userService service.UserService = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)
)

func main() {
	const port string = "8080"

	httpRouter.POST("/users", userController.AddUser)
	httpRouter.GET("/users", userController.FindAllUsers)

	httpRouter.Serve(port)
}