package utils

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func SetupConfig() error {
	viper.AddConfigPath("..")
	viper.AddConfigPath("../../") // for tests
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Error().Msg(fmt.Sprint("Error reading env file", err))
		return err
	}
	return nil
}

func SetupLogger(debug bool) error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log_file_name := "core.log"
	os.Remove(log_file_name) // ignore error if file already exists

	file, err := os.OpenFile(log_file_name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	log.Logger = zerolog.New(file).With().Logger()
	log.Info().Msg("Starting core")

	return nil
}