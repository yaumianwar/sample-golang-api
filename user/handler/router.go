package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yaumianwar/sample-golang-api/user"
)

// Blueprint register package router to main package
func Blueprint(routerGroup *gin.RouterGroup, db *gorm.DB) {
	svc := user.NewService(db)

	// user group router
	userGroup := routerGroup.Group("/user")
	userGroup.GET("", GetAllUsers(&svc))
	userGroup.POST("", CreateUser(&svc))
	userGroup.GET("/:id", GetSingleUser(&svc))
	userGroup.POST("/:id", UpdateUser(&svc))
	userGroup.DELETE("/:id", DeleteUser(&svc))
	userGroup.GET("/:id/skincare", GetAllSkincareByUser(&svc))
	userGroup.POST("/:id/skincare", CreateSkincare(&svc))
	userGroup.GET("/:id/skincare/:skincareId", GetSingleSkincare(&svc))
	userGroup.POST("/:id/skincare/:skincareId", UpdateSkincare(&svc))
	userGroup.DELETE("/:id/skincare/:skincareId", DeleteSkincare(&svc))
}
