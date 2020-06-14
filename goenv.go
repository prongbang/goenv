package goenv

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Env is the map
type Env map[string]string

func init() {
	LoadEnv()
}

// LoadEnv is the function load .env file
func LoadEnv(filename ...string) error {
	if len(filename) == 0 {
		filename = []string{".env"}
	}

	for _, file := range filename {
		f, e := os.Open(file)
		if e != nil {
			return e
		}
		defer f.Close()

		env := parse(f)
		for key, val := range env {
			Setenv(key, val)
		}
	}
	return nil
}

// Setenv is the function set key:value ton environment
func Setenv(key string, val string) {
	os.Setenv(key, val)
}

func parse(r io.Reader) Env {
	env := make(Env)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "=") {
			i := strings.Index(text, "=")
			if i > -1 {
				key := text[:i]
				val := text[i+1:]
				env[key] = val
			}
		}
	}

	return env
}
