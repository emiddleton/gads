package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

var commands = map[string][]string{
	"windows": []string{"cmd", "/c", "start"},
	"darwin":  []string{"open"},
	"linux":   []string{"xdg-open"},
	"freebsd": []string{"xdg-open"},
	"netbsd":  []string{"xdg-open"},
	"openbsd": []string{"xdg-open"},
}

// Open calls the OS default program for uri
// e.g. Open("http://www.google.com") will open the default browser on www.google.com
func webbrowserOpen(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	// Escape characters not allowed by cmd
	switch runtime.GOOS {
	case "windows":
		uri = strings.Replace(uri, "&", `^&`, -1)
	}

	run = append(run, uri)
	cmd := exec.Command(run[0], run[1:]...)

	return cmd.Start()
}
