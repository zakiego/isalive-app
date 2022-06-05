package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

var skippedWindows = []string{"ApplicationFrameHost - Settings", "TextInputHost - Windows Input Experience", "SystemSettings - Settings"}

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
	firstWindows := true

	var output string
	for id, line := range strings.Split(strings.TrimRight(string(stdout), "\n"), "\n") {
		if id < 3 {
			continue
		}

		words := strings.Fields(line)
		window := words[1] + " - " + strings.Join(words[2:], " ")

		if contains(window, skippedWindows) {
			continue
		}

		if firstWindows {
			output = window
			firstWindows = false
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
