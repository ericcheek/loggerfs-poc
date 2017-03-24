package loggingfs

import (
	"fmt"
)

type StdoutTransport struct {
	transport *Transport
}

func WriteMessage(f *LogFile, message string) {
	fmt.Printf("%s", message)
}
