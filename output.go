package acres

import (
	"io"
	"os"
)

type OutputSettings struct {
	Writer io.Writer
	Prefix string
	Muted  bool
}

var DefaultOutputSettings = OutputSettings{
	Writer: os.Stdout,
	Prefix: "Result code:",
	Muted:  false,
}
