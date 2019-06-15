package server

import "github.com/gin-gonic/gin"

type router func(*gin.Engine) *gin.Engine

// Engine returns gin.Engine
func Engine(r router, runMode string) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	if runMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine = r(engine)

	return engine
}
