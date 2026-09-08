package logs

import (
	"log/slog"
	"storage-api/shared/events"
)

func Log(event events.Event) {
	logEvent, ok := event.(LogEvent)
	if !ok {
		return
	}
	message := logEvent.message
	if message == "" {
		message = "Unknown"
	}
	switch logEvent.logType {
	case LogTypeInfo:
		slog.Info(logEvent.message)
	case LogTypeWarning:
		slog.Warn(logEvent.message)
	case LogTypeError:
		slog.Error(logEvent.message)
	}
}
