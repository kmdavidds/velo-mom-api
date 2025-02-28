package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/valentinusdelvin/velo-mom-api/addition/jwt"
	"github.com/valentinusdelvin/velo-mom-api/internal/usecase"
)

type Interface interface {
	Authenticate(c *gin.Context)
}

type Middleware struct {
	usecase *usecase.Usecase
	jwt     jwt.InterJWT
}

func Init(usecase *usecase.Usecase) Interface {
	return &Middleware{
		usecase: usecase,
		jwt:     jwt.NewJWT(),
	}
}
