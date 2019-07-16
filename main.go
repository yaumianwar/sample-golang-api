package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"net/http"

	user "github.com/yaumianwar/sample-golang-api/user/handler"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	_ = viper.ReadInConfig()
}

func main() {

	// connect to mysql databse
	db, err := gorm.Open("mysql", viper.GetString("datasource"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.LogMode(viper.GetBool("app.show-sql"))

	// router config
	r := gin.Default()
	if viper.GetBool("app.production") {
		gin.SetMode(gin.ReleaseMode)
	}

	// err message for invalid route
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": gin.H{"message": "resource not found"}})
	})

	// register router blueprint
	user.Blueprint(r.Group("/v1/user"), db)

	// run
	if err := r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port"))); err != nil {
		log.Fatal(err)
	}

}
