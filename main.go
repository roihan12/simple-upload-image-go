package main

import (
	"fmt"
	// "os"

	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/controllers"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/database"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/helpers"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/repository"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/routers"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/services"
)

func main() {
	db := database.DBConnection()
	// validate := validator.New()

	authService := helpers.NewAuthService()

	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	photoRepository := repository.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepository)

	controller := &routers.Controllers{
		UserController:  controllers.NewUserController(userService, authService),
		PhotoController: controllers.NewPhotoController(photoService, authService),
	}

	middleware := &routers.AuthMiddleware{
		AuthService: authService,
		UserService: userService,
	}

	r := routers.NewRouter(controller, middleware)

	err := r.Run(":3000")
	if err != nil {
		fmt.Println("Error on the route run")
	}
	// fmt.Println("Running on " + os.Getenv("POSTGRES_HOST") + " : " + os.Getenv("PORT"))
}
