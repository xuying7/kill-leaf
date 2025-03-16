package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/xuying7/kill-leaf/internal/config"
)

const (
	googleAuthURL  = "https://accounts.google.com/o/oauth2/v2/auth"
	googleTokenURL = "https://oauth2.googleapis.com/token"
	googleUserInfo = "https://www.googleapis.com/oauth2/v2/userinfo"
)

// GetGoogleLoginURL 生成 Google 登录URL
func GetGoogleLoginURL(state string) string {
	v := url.Values{}
	v.Set("client_id", config.EnvVar("GOOGLE_CLIENT_ID", ""))
	v.Set("redirect_uri", config.EnvVar("GOOGLE_REDIRECT_URI", "http://localhost:8080/auth/google/callback"))
	v.Set("response_type", "code")
	v.Set("scope", "openid email profile")
	v.Set("state", state)

	return fmt.Sprintf("%s?%s", googleAuthURL, v.Encode())
}

// ExchangeCodeForToken 用授权码向 Google 换取令牌
func ExchangeCodeForToken(code string) (map[string]interface{}, error) {
	data := url.Values{}
	data.Set("client_id", config.EnvVar("GOOGLE_CLIENT_ID", ""))
	data.Set("client_secret", config.EnvVar("GOOGLE_CLIENT_SECRET", ""))
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", config.EnvVar("GOOGLE_REDIRECT_URI", "http://localhost:8080/auth/google/callback"))
	data.Set("code", code)

	resp, err := http.PostForm(googleTokenURL, data)
	if err != nil {
		return nil, fmt.Errorf("error exchanging code for token: %v", err)
	}
	defer resp.Body.Close()

	var tokenRes map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&tokenRes); err != nil {
		return nil, fmt.Errorf("error decoding token response: %v", err)
	}
	return tokenRes, nil
}

// GetGoogleUserInfo 用 access_token 获取用户信息
func GetGoogleUserInfo(accessToken string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", googleUserInfo, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting user info: %v", err)
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("error decoding user info: %v", err)
	}
	return userInfo, nil
}
