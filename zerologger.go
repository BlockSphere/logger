package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	z zerolog.Logger
}

func NewZeroLogger(pretty bool) Logger {
	var zl zerolog.Logger
	if pretty {
		zl = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	} else {
		zl = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

	return &ZeroLogger{z: zl}
}

func (z *ZeroLogger) Info(msg string, args ...any) {
	z.z.Info().Msgf(msg, args...)
}

func (z *ZeroLogger) Error(msg string, args ...any) {
	z.z.Error().Msgf(msg, args...)
}

func (z *ZeroLogger) Debug(msg string, args ...any) {
	z.z.Debug().Msgf(msg, args...)
}

func (z *ZeroLogger) Warn(msg string, args ...any) {
	z.z.Warn().Msgf(msg, args...)
}

func (z *ZeroLogger) Fatal(msg string, args ...any) {
	z.z.Fatal().Msgf(msg, args...)

}
