package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"rest-api/middleware"
	ginitem "rest-api/modules/item/transport/gin"
	"rest-api/modules/media"
	"time"
)

func main() {
	postgresDsn := "host=localhost user=postgres password=postgres dbname=todos port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	postgresDb, postgresErr := gorm.Open(postgres.Open(postgresDsn), &gorm.Config{})
	if postgresErr != nil {
		log.Fatal(postgresErr)
	}

	dsn := os.Getenv("DB_CONNECTION_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	r := gin.Default()

	r.Use(middleware.RecoveryGin()) // Cách 1: sử dụng với tất cả api

	//api := r.Group("/api", middleware.Recovery(), ...n middleware) // Cách 2 chỉ apply middleware cho 1 group api nào đó
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.PUT("medias", media.UploadHandler(db))

			items := v1.Group("/items")
			{
				items.GET("/pos", ginitem.GetPosItems(postgresDb))
				items.GET("", ginitem.GetItems(db))
				items.POST("", ginitem.CreateItem(db))
				items.POST("/pos", ginitem.CreatePosItem(postgresDb))
				items.GET("/:id", ginitem.GetItem(db))
				//items.GET("/:id", middleware.Recovery(), ginitem.GetItem(db)) // Chỉ apply middleware cho 1 api duy nhất
				items.PUT("/:id", ginitem.UpdateItem(db))
				items.DELETE("/:id", ginitem.DeleteItem(db))
			}
		}
	}
	r.GET("/ping", func(c *gin.Context) {
		//fmt.Println([]int{}[1]) // Điểm gây panic (err crash ứng dụng)

		// Vì khi chạy goroutine k cùng stack với recovery middleware nên ta cần khai phải chạy recover ngay trong gorountine đó
		// Handle recover gorountine cho ứng dụng k crash khi sử dụng goroutine có panic
		//go func() {
		//	defer common.Recovery()
		//	fmt.Println([]int{}[1])
		//}()

		c.JSON(http.StatusOK, gin.H{
			"data": "pong",
		})
	})
	r.Run()
}
