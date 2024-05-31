package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Hostname             string
	Url                  string
	Port                 string
	BasicAuthCredentials map[string]string = make(map[string]string)
)

func init() {
	godotenv.Load(".env")

	Hostname = os.Getenv("HOSTNAME")
	Url = fmt.Sprintf("https://%v", Hostname)
	Port = os.Getenv("PORT")
	BasicAuthCredentials[os.Getenv("BASIC_AUTH_USERNAME")] = os.Getenv("BASIC_AUTH_PASSWORD")
}
