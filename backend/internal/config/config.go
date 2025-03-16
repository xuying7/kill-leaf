package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv 加载.env文件(如果存在)，否则读取系统环境变量
func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found. Relying on system environment variables.")
	}
	return nil
}

// EnvVar 从环境变量读取key，若不存在则用fallback
func EnvVar(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
