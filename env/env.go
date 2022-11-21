package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	listOfEnvVariables = []string{
		"PROCESS_NAME",
		"API_HOST", "API_PORT", "API_PATH",
		"TIME_LAG",
	}
	Env map[string]string
)

func init() {
	Env = map[string]string{}
	env_vars, _ := godotenv.Read(".env")
	fmt.Println("Environment variables: ")
	for _, envVar := range listOfEnvVariables {
		if val, ok := os.LookupEnv(envVar); ok {
			Env[envVar] = strings.TrimSpace(val)
		} else if val, ok := env_vars[envVar]; ok {
			Env[envVar] = strings.TrimSpace(val)
		} else {
			Env[envVar] = ""
		}
		fmt.Println(envVar, ": ", Env[envVar])
	}
}
