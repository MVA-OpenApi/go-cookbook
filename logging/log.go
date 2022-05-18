package main

import (
		"errors"
		"os"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func main() {
	  // set time format 
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

		// use pretty logging
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

		// enable deubgging
		zerolog.SetGlobalLevel(zerolog.DebugLevel)

    log.Debug().Msg("This message appears only when log level set to Debug")
    log.Warn().Msg("This message appears only when log level set to Warn or above")
    log.Info().Msg("This message appears when log level set to Info or above")

		// logging errors
		err := errors.New("seems we have an error here")
		log.Error().Err(err).Msg("")

}