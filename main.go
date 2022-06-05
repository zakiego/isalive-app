package main

import (
	"runtime"
	"time"

	"github.com/gonutz/w32/v2"
)

func main() {
	os := runtime.GOOS

	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, w32.SW_HIDE)
		}
	}

	func() {
		for range time.Tick(time.Second * 3) {

			switch os {
			case "windows":
				Windows()
			case "linux":
				Linux()
			}

		}
	}()

}
