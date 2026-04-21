package hosts

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const lineToAdd = "10.96.16.67\talmapps.online"

func hostsPath() string {
	windir := os.Getenv("SystemRoot")
	if windir == "" {
		windir = `C:\Windows`
	}
	return filepath.Join(windir, "System32", "drivers", "etc", "hosts")
}

func IsLinePresent() (bool, error) {
	data, err := ioutil.ReadFile(hostsPath())
	if err != nil {
		return false, err
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == lineToAdd {
			return true, nil
		}
		if strings.Contains(line, "10.96.16.67") && strings.Contains(line, "almapps.online") {
			return true, nil
		}
	}
	return false, nil
}

func AddLineIfMissing() error {
	present, err := IsLinePresent()
	if err != nil {
		return err
	}
	if present {
		return nil
	}

	data, err := ioutil.ReadFile(hostsPath())
	if err != nil {
		return err
	}
	lines := strings.Split(string(data), "\n")

	// Insert after the last localhost line
	idx := -1
	for i, l := range lines {
		if strings.Contains(strings.ToLower(l), "localhost") {
			idx = i
		}
	}
	insertIndex := len(lines)
	if idx >= 0 {
		insertIndex = idx + 1
	}

	lines = append(lines, "")
	copy(lines[insertIndex+1:], lines[insertIndex:])
	lines[insertIndex] = lineToAdd
	out := strings.Join(lines, "\r\n")
	return ioutil.WriteFile(hostsPath(), []byte(out), 0644)
}
