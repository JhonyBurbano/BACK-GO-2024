package main

import (
	"github.com/jnates/smartOshApi/infrastructure"
	"github.com/jnates/smartOshApi/infrastructure/kit/enum"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting API CMD")
	infrastructure.InitLogger()

	port := os.Getenv(enum.APIPort)
	infrastructure.Start(port)
}
