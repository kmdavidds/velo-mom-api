package usecase

import (
	"github.com/valentinusdelvin/velo-mom-api/internal/repository"
	"github.com/valentinusdelvin/velo-mom-api/utils/bcrypt"
	"github.com/valentinusdelvin/velo-mom-api/utils/jwt"
)

type Usecase struct {
	UserUsecase    InterUserUsecase
	ArticleUsecase InterArticleUsecase
}

type InitializersParam struct {
	Repository *repository.Repository
	Bcrypt     *bcrypt.InterBcrypt
	JWT        *jwt.InterJWT
}

func NewUsecase(param InitializersParam) *Usecase {
	UserUsecase := NewUserUsecase(param.Repository.UserRepository, *param.Bcrypt, *param.JWT)
	ArticleUsecase := NewArticleUsecase(param.Repository.ArticleRepository)

	return &Usecase{
		UserUsecase:    UserUsecase,
		ArticleUsecase: ArticleUsecase,
	}
}
