package logs

type LogEvent struct {
	message string
	logType LogType
}

const EventName = "log"

type LogType string

const (
	LogTypeError   LogType = "ERROR"
	LogTypeInfo    LogType = "INFO"
	LogTypeWarning LogType = "WARN"
)

func NewLogEvent(message string, logType LogType) LogEvent {
	return LogEvent{
		message: message,
		logType: logType,
	}
}
