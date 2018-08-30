package ex11

import (
	"fmt"
	"os"
	"os/exec"
)

const editorVariable = "EDITOR"

func Edit(file string) error {
	editor, ok := os.LookupEnv(editorVariable)
	if !ok {
		return fmt.Errorf("environment variable %s must be set", editorVariable)
	}
	cmd := exec.Command(editor, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("running %s on %s: %v", editor, file, err)
	}
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("waiting for %s on %s: %v", editor, file, err)
	}
	return nil
}
