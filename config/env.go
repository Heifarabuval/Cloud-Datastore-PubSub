package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetEnvConst(key string)  (string, bool) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	response, ok := os.LookupEnv(key)
	return response,ok
}
