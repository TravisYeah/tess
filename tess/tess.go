package tess

import (
	"log"
	"os"
	"os/exec"

	"golang.design/x/clipboard"
)

type Tess struct{}

func (t *Tess) Run(args []string) {
	if len(args) == 0 {
		args = []string{"clip", "clip"}
	}

	if args[0] == "clip" {
		f, err := os.CreateTemp("", "tesseract")
		if err != nil {
			log.Fatal(err)
		}

		err = clipboard.Init()
		if err != nil {
			panic(err)
		}

		bytes := clipboard.Read(clipboard.FmtImage)
		_, err = f.Write(bytes)
		if err != nil {
			log.Fatal(err)
		}

		args[0] = f.Name()

		defer os.Remove(f.Name())
	}

	isClipOut := false
	if len(args) > 1 && args[1] == "clip" {
		isClipOut = true
		args[1] = "stdout"
	}

	cmd := exec.Command("tesseract", args...)
	bytes, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	if isClipOut {
		clipboard.Write(clipboard.FmtText, bytes)
	}

	log.Print(string(bytes))
}
