package config

var (
	logger *Logger
)

func GetLogger(p string) *Logger {
	// INITIALIZE LOGGER
	logger = NewLogger(p)
	return logger
}
