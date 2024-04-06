package screen_cleaner

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
)

var handlers = map[string]func(){
	"linux":   clearScreenInPosix,
	"windows": clearScreenInWindows,
	"darwin":  clearScreenInPosix,
}

func ClearScreen() error {
	currentOS := runtime.GOOS
	handler, handlerFound := handlers[currentOS]
	if !handlerFound {
		return errors.New("Unsupported OS: " + currentOS + ". Can't clear screen")
	}

	handler()
	return nil
}

func clearScreenInPosix() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func clearScreenInWindows() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
