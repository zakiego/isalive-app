package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	cron "github.com/robfig/cron/v3"
)

func main() {
	os := runtime.GOOS

	switch os {
	case "windows":
		CronWindows()
	case "linux":
		ubuntu()
	}
}

func CronWindows() {
	c := cron.New()
	// i := 1
	c.AddFunc("* * * * *", func() {
		Windows()
	})
	c.Start()
	time.Sleep(time.Minute * 5)
}

func Windows() {

	file, _ := os.Create(`F:\Code\go\isalive-app\log.txt`)

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
		output = output + "\n" + window
	}

	file.WriteString(output)
	file.Close()

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

func ubuntu() {

	// get list run application
	cmd := exec.Command(`/bin/sh`, "-c", `wmctrl -l|awk '{$3=""; $2=""; $1="";  print $0}'`)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// print list run application
	// trim space and join again
	var output string
	for id, line := range strings.Split(strings.TrimRight(string(stdout), "\n"), "\n") {
		if id == 0 {
			output = strings.TrimSpace(line)
			continue
		}
		output = output + "\n" + strings.TrimSpace(line)
	}

	fmt.Println(output)

	link := "https://api-isalive.zakiego.workers.dev/?device=laptop-ubuntu&status="
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
