package main

import (
	"log"
)

const (
	INFO  = "info"
	DEBUG = "debug"
	WARN  = "warn"
	ERROR = "error"
	TRACE = "trace"
)

type Logger struct {
	Level string
}

func (logger Logger) Debug(message string) {
	if logger.Level == DEBUG || logger.Level == TRACE {
		log.Printf("\x1b[1;32m [%s] \x1b[0m : %s \n", "DEBUG", message)
	}
}

func (logger Logger) Error(message string) {
	log.Printf("\x1b[1;31m [%s] \x1b[0m : %s \n", "ERROR", message)
}

func (logger Logger) Warn(message string) {
	if logger.Level == WARN || logger.Level == DEBUG || logger.Level == INFO || logger.Level == TRACE {
		log.Printf("\x1b[1;33m [%s] \x1b[0m  : %s \n", "WARN", message)
	}
}

func (logger Logger) Trace(message string) {
	if logger.Level == TRACE {
		log.Printf("\x1b[1;36m [%s] \x1b[0m : %s \n", "TRACE", message)
	}
}

func (logger Logger) Info(message string) {
	if logger.Level == INFO || logger.Level == DEBUG || logger.Level == TRACE {
		log.Printf("\x1b[1;34m [%s] \x1b[0m  : %s \n", "INFO", message)
	}
}
