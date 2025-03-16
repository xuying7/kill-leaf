package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuying7/kill-leaf/internal/db"
	"github.com/xuying7/kill-leaf/internal/models"
	"github.com/xuying7/kill-leaf/internal/services"
	"github.com/xuying7/kill-leaf/internal/utils"
)

// GoogleLoginHandler 处理 /auth/google
func GoogleLoginHandler(c *gin.Context) {
	// 1. 生成 state 并存session
	state := utils.GenerateState(16)
	utils.SaveStateInSession(c, state)

	// 2. 重定向至 Google 登录
	loginURL := services.GetGoogleLoginURL(state)
	c.Redirect(http.StatusTemporaryRedirect, loginURL)
}

// GoogleCallbackHandler 处理 /auth/google/callback
func GoogleCallbackHandler(c *gin.Context) {
	// 1. 获取 code, state
	code := c.Query("code")
	returnedState := c.Query("state")

	// 2. 校验 state
	savedState := utils.GetStateFromSession(c)
	if returnedState == "" || returnedState != savedState {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
		return
	}

	// 3. 用 code 换令牌
	tokenRes, err := services.ExchangeCodeForToken(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessToken, ok := tokenRes["access_token"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No access_token"})
		return
	}

	// 4. 获取用户信息
	userInfo, err := services.GetGoogleUserInfo(accessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 一般 Google 返回 userInfo["id"] / userInfo["email"] / userInfo["name"] / userInfo["picture"]
	sub := userInfo["id"].(string)
	email := userInfo["email"].(string)
	name := userInfo["name"].(string)
	pic := userInfo["picture"].(string)

	// 5. 在数据库查找/创建用户
	var user models.User
	db.DB.Where("google_sub = ?", sub).First(&user)
	if user.ID == 0 {
		user = models.User{
			GoogleSub:  sub,
			Email:      email,
			Name:       name,
			PictureURL: pic,
		}
		db.DB.Create(&user)
	} else {
		// 如果已存在，则更新一些字段
		user.Email = email
		user.Name = name
		user.PictureURL = pic
		db.DB.Save(&user)
	}

	// 6. 设置 session
	utils.SetUserSession(c, user.Email)

	// 7. 重定向到受保护页面
	c.Redirect(http.StatusFound, "/protected")
}
