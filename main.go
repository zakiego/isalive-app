package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)



func main() {
	os := runtime.GOOS

	switch os {
	case "windows":
		fmt.Println("Windows")
	case "linux":
		ubuntu()
	}
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
