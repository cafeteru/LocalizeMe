package tools

import (
	"github.com/joho/godotenv"
	"path/filepath"
)

func LoadEnv() {
	environmentPath := filepath.Join("./", ".env")
	_ = godotenv.Load(environmentPath)
}
