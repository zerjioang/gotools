package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "",
	})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// generic logging function
func Debug(args ...interface{}) {
	log.Debug(args)
}

func Info(args ...interface{}) {
	log.Info(args)
}

// generic logging function
func Error(args ...interface{}) {
	log.Error(args)
}

func Enabled(status bool) {

}
