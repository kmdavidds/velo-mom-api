package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/valentinusdelvin/velo-mom-api/initializers"
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/internal/rest"
	"github.com/valentinusdelvin/velo-mom-api/internal/usecase"
	"github.com/valentinusdelvin/velo-mom-api/utils/bcrypt"
	"github.com/valentinusdelvin/velo-mom-api/utils/jwt"
	"github.com/valentinusdelvin/velo-mom-api/utils/middleware"
	"github.com/valentinusdelvin/velo-mom-api/utils/supabase"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	bcrypt := bcrypt.Init()
	jwt := jwt.NewJWT()
	db := initializers.ConnectToDB()
	supabase := supabase.Init()
	initializers.AutoMigrate(db)

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(usecase.InitializersParam{
		Repository: repository,
		Bcrypt:     &bcrypt,
		JWT:        &jwt,
		Supabase:   &supabase,
	})
	middleware := middleware.Init(usecase)

	rest := rest.NewRest(usecase, middleware)
	rest.FinalCheck()

	//Janlup Buat apus ini tiap Jumat
	r := gin.Default()
	for _, route := range r.Routes() {
		fmt.Println(route.Method, route.Path)
	}
	rest.Run()
}
