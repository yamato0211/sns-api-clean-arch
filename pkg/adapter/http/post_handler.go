package http

import (
	"clean-sns-api/pkg/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type postHandler struct {
	usecase usecase.IPostUsecase
}

func NewPostHandler(pu usecase.IPostUsecase) *postHandler {
	return &postHandler{
		usecase: pu,
	}
}

func (ph *postHandler) FindAllPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		posts, err := ph.usecase.FindAllPosts(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, posts)
	}
}

func (ph *postHandler) FindPostByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		postID, err := strconv.Atoi(c.Param("post_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		post, err := ph.usecase.FindPostByID(ctx, postID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, post)
	}
}
