package api

import (
	"os"
)

type Api struct {
	APIKey string
}

func NewApi() *Api {
	return &Api{
		APIKey: os.Getenv("API_KEY"),
	}
}
