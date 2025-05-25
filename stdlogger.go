package logger

import (
	"log"
)

type StdLogger struct{}

func NewStdLogger() Logger {
	return &StdLogger{}
}

func (s *StdLogger) Info(msg string, args ...any) {
	log.Printf("[INFO] "+msg, args...)
}

func (s *StdLogger) Error(msg string, args ...any) {
	log.Printf("[ERROR] "+msg, args...)
}

func (s *StdLogger) Debug(msg string, args ...any) {
	log.Printf("[DEBUG] "+msg, args...)
}

func (s *StdLogger) Warn(msg string, args ...any) {
	log.Printf("[WARN] "+msg, args...)
}

func (s *StdLogger) Fatal(msg string, args ...any) {
	log.Fatalf("[FATAL] "+msg, args...)
}
