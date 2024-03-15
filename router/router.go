package router

import "github.com/gin-gonic/gin"

type Server interface {
	Option() interface{}
	Engin() *gin.Engine
	Group(relativePath string) *gin.RouterGroup
}

func Setup(s Server) {
}
