package tess

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"golang.design/x/clipboard"
)

func TestNoArgsInput(t *testing.T) {
	log.SetFlags(0)

	var buf bytes.Buffer
	log.SetOutput(&buf)

	defer func() {
		log.SetOutput(os.Stderr)
	}()

	bytes, err := os.ReadFile("test.png")
	if err != nil {
		t.Error(err)
	}

	clipboard.Write(clipboard.FmtImage, bytes)

	tess := &Tess{}
	tess.Run([]string{})
	res := strings.TrimSpace(buf.String())
	if res != "Hello, world!" {
		t.Errorf("unexpected output: '%s'", res)
	}
}

func TestHelp(t *testing.T) {
	log.SetFlags(0)

	var buf bytes.Buffer
	log.SetOutput(&buf)

	defer func() {
		log.SetOutput(os.Stderr)
	}()

	tess := &Tess{}
	tess.Run([]string{"--help"})
	res := buf.String()
	if !strings.HasPrefix(res, "Usage:") {
		t.Errorf("unexpected output: '%s'", res)
	}
}

func TestFileInput(t *testing.T) {
	log.SetFlags(0)

	var buf bytes.Buffer
	log.SetOutput(&buf)

	defer func() {
		log.SetOutput(os.Stderr)
	}()

	tess := &Tess{}
	tess.Run([]string{"test.png", "stdout"})
	res := strings.TrimSpace(buf.String())
	if res != "Hello, world!" {
		t.Errorf("unexpected output: '%s'", res)
	}
}
