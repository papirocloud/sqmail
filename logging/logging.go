package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

func GetLogger(logFormat string) Logger {
	zerolog.DurationFieldUnit = time.Millisecond
	zerolog.TimeFieldFormat = time.RFC3339

	var multi zerolog.LevelWriter

	var wr diode.Writer

	switch logFormat {
	case "color", "text":
		wr = diode.NewWriter(consoleWriter(logFormat), 1000, 10*time.Millisecond, func(missed int) {
			fmt.Printf("Dropped %d messages", missed)
		})
	default:
		wr = diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
			fmt.Printf("Dropped %d messages", missed)
		})
	}

	multi = zerolog.MultiLevelWriter(wr)

	logger := zerolog.New(multi).With().Timestamp().Logger()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	return Logger{logger}
}

func consoleWriter(logFormat string) zerolog.ConsoleWriter {
	var color bool

	if logFormat == "color" {
		color = true
	}

	writer := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05", NoColor: !color}

	writer.PartsOrder = []string{
		zerolog.TimestampFieldName,
		zerolog.CallerFieldName,
		zerolog.LevelFieldName,
		zerolog.MessageFieldName,
	}

	writer.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("| %-60s|", i)
	}

	return writer
}
