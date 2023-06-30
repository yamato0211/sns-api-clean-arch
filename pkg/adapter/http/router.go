package http

import (
	"clean-sns-api/pkg/domain/service"
	"clean-sns-api/pkg/infra"
	"clean-sns-api/pkg/infra/postgres"
	"clean-sns-api/pkg/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	apiVersion   = "/v1"
	postsAPIRoot = apiVersion + "/posts"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())
	postgresConn := infra.NewPostgresConnector()
	postRepository := postgres.NewPostRepository(postgresConn.Conn)
	postService := service.NewPostService(postRepository)
	postUsecase := usecase.NewPostUsecase(postService)

	postGroup := r.Group(postsAPIRoot)
	{
		handler := NewPostHandler(postUsecase)
		postGroup.GET("", handler.FindAllPosts())
		postGroup.GET("/:post_id", handler.FindPostByID())
	}

	return r
}
