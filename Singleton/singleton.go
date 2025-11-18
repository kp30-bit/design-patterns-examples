package singleton

import (
	"log"
	"os"
	"sync"
)

// Creational Design Pattern
// Only one instance of a class is created and provided a global point of access.
// Used in Loggers, Config Managers, Database Managers
// Used when you want a single shared instance across the system.
// Used to save memory and avoid multiple initializations.
// Thread safety is there, you can call it from multiple go routines.
// Prevents duplication

// cons

// Can become an additional dependency

type Logger struct {
	logger *log.Logger
}

var instance *Logger
var once sync.Once

func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{
			logger: log.New(os.Stdout, "LOG: ", log.LstdFlags),
		}
	})
	return instance
}

func (l *Logger) Log(message string) {
	l.logger.Printf(message)
}
