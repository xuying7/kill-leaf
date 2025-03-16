package main

import (
	"log"

	"github.com/xuying7/kill-leaf/internal/config"
	"github.com/xuying7/kill-leaf/internal/db"
	"github.com/xuying7/kill-leaf/internal/handlers"
	"github.com/xuying7/kill-leaf/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载环境变量
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}

	// 2. 初始化 MySQL 数据库
	err = db.InitDB()
	if err != nil {
		log.Fatalf("Failed to init DB: %v", err)
	}

	// 3. 创建 Gin 引擎
	r := gin.Default()

	// 4. 启用 session 中间件
	utils.SetupSession(r)

	// 5. 路由: Google OAuth
	r.GET("/auth/google", handlers.GoogleLoginHandler)
	r.GET("/auth/google/callback", handlers.GoogleCallbackHandler)

	// 6. 测试受保护路由
	r.GET("/protected", func(c *gin.Context) {
		// session.go 里可自动检查是否登录
		userEmail := utils.GetUserSession(c)
		if userEmail == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.JSON(200, gin.H{
			"message": "You are logged in with Google!",
			"email":   userEmail,
		})
	})

	// 7. 启动服务
	port := config.EnvVar("PORT", "8080")
	r.Run(":" + port)
}
