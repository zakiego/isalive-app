package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

var skipped = []string{"@!0,0;BDHF", "@!0,1080;BDHF"}

func Linux() {

	// get list run application
	cmd := exec.Command(`/bin/sh`, "-c", `wmctrl -l|awk '{$3=""; $2=""; $1="";  print $0}'`)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// print list run application
	// trim space and join again

	first := false

	var output string
	for _, line := range strings.Split(strings.TrimRight(string(stdout), "\n"), "\n") {
		if contains(strings.TrimSpace(line), skipped) {
			continue
		}

		if first == false {
			output = strings.TrimSpace(line)
			first = true
			continue
		}
		output = output + "; " + strings.TrimSpace(line)
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
