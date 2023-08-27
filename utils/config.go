package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func SetupConfig() error {
	viper.AddConfigPath(".")
	viper.AddConfigPath("..") // for tests
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Error().Msg(fmt.Sprint("Error reading env file", err))
		return err
	}
	return nil
}

func SetupLogger(debug bool, log_file_name string) error {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	os.Remove(log_file_name)
	//if err != nil {
	// ignore error if file already exists
	//}

	file, err := os.OpenFile(log_file_name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	fileLogger := zerolog.New(file).With().Timestamp().Logger()

	consoleLogger := zerolog.ConsoleWriter{Out: os.Stdout}
	writers := io.MultiWriter(consoleLogger, fileLogger)
	log.Logger = log.Output(writers)

	log.Info().Msgf("Starting log '%s'", log_file_name)

	return nil
}
