package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gonutz/w32"
)

func main() {
	os := runtime.GOOS

	switch os {
	case "windows":
		console := w32.GetConsoleWindow()
		if console != 0 {
			_, consoleProcID := w32.GetWindowThreadProcessId(console)
			if w32.GetCurrentProcessId() == consoleProcID {
				w32.ShowWindowAsync(console, w32.SW_HIDE)
			}
		}
		func() {
			for range time.Tick(time.Minute * 3) {
				Windows()
			}
		}()

	case "linux":
		func() {
			// for range time.Tick(time.Minute * 3) {
			for range time.Tick(time.Second * 3) {
				Linux()
			}
		}()
	}

}

func Windows() {

	// file, _ := os.Create(`F:\Code\go\isalive-app\log.txt`)

	// get list run application
	cmd := exec.Command(`F:\Code\go\isalive-app\wmctrl.exe`, "-c", "-l")

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// print list run application
	// trim space and join again
	var output string
	for id, line := range strings.Split(strings.TrimRight(string(stdout), "\n"), "\n") {
		if id < 3 {
			continue
		}

		words := strings.Fields(line)
		window := words[1] + " - " + strings.Join(words[2:], " ")

		if id == 3 {
			output = window
			continue
		}
		output = output + "; " + window
	}

	fmt.Println(output)
	// file.WriteString(output)
	// file.Close()

	link := "https://api-isalive.zakiego.workers.dev/?device=laptop-windows&status="
	api := fmt.Sprint(link, url.QueryEscape(output))

	// get api
	resp, err := http.Get(string(api))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// read response
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(body))

}
