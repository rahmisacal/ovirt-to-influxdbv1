package main

import (
	"./infrastructure"
	"os"
)

func main() {
	url := os.Getenv("url")
	token := infrastructure.GetToken(url + "sso/oauth/token")
	infrastructure.MapMetric(url+"api/hosts", token)
}
