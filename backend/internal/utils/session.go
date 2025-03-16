package utils

import (
	"math/rand"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SetupSession 在 Gin 中启用 session 中间件
func SetupSession(r *gin.Engine) {
	store := cookie.NewStore([]byte("secret1234")) // 替换成更安全的随机密钥
	r.Use(sessions.Sessions("myapp_session", store))
}

// GenerateState 生成指定长度的随机字符串
func GenerateState(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	state := make([]byte, length)
	for i := 0; i < length; i++ {
		state[i] = chars[rand.Intn(len(chars))]
	}
	return string(state)
}

// SaveStateInSession 把 state 存入 session
func SaveStateInSession(c *gin.Context, state string) {
	session := sessions.Default(c)
	session.Set("oauthState", state)
	session.Save()
}

// GetStateFromSession 取出 state
func GetStateFromSession(c *gin.Context) string {
	session := sessions.Default(c)
	val := session.Get("oauthState")
	if val == nil {
		return ""
	}
	return val.(string)
}

// SetUserSession 把用户 Email 存入 session
func SetUserSession(c *gin.Context, email string) {
	session := sessions.Default(c)
	session.Set("user_email", email)
	session.Save()
}

// GetUserSession 从 session 拿到用户 Email
func GetUserSession(c *gin.Context) string {
	session := sessions.Default(c)
	val := session.Get("user_email")
	if val == nil {
		return ""
	}
	return val.(string)
}
