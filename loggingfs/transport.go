package loggingfs

type Transport interface {
	WriteMessage(f *LogFile, message string)
}
