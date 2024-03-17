package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func SetupTestConfig() error {
	return SetupConfigPath("..")
}

func SetupConfigPath(rootPath string) error {
	viper.SetConfigFile(filepath.Join(rootPath, ".env"))
	viper.SetConfigType("env")
	if err := viper.MergeInConfig(); err != nil {
		log.Error().Msg(fmt.Sprint("Error reading env file in ", rootPath, err))
		return err
	}
	return nil
}

func SetupLogger(debug bool, logFileName string) error {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	os.Remove(logFileName)
	//if err != nil {
	// ignore error if file already exists
	//}

	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	fileLogger := zerolog.New(file).With().Timestamp().Logger()

	consoleLogger := zerolog.ConsoleWriter{Out: os.Stdout,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("| %s |", i)
		},
		// FormatCaller: func(i interface{}) string {
		// 	return filepath.Base(fmt.Sprintf("%s", i))
		// },
		PartsExclude: []string{
			zerolog.TimestampFieldName,
		}}
	writers := io.MultiWriter(consoleLogger, fileLogger)
	log.Logger = log.Output(writers)

	log.Info().Msgf("Starting log '%s'", logFileName)

	return nil
}
