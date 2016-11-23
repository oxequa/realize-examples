package sub

import (
	"io"
	"os"
)

type SpeakWriter struct {
	w io.Writer
}

func (s *SpeakWriter) Say(message string) {
	io.WriteString(os.Stdout, "writes a string ")
}
