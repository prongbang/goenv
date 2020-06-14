package goenv_test

import (
	"os"
	"testing"

	"github.com/prongbang/goenv"
)

func TestLoadEnvDefault(t *testing.T) {
	err := goenv.LoadEnv()
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	if err != nil {
		t.Error("Read .env file wrong.")
	}

	if dbUser == "" && dbPass == "" {
		t.Error("Read .env file wrong.")
	}
}

func TestLoadEnvBasicAuthen(t *testing.T) {
	err := goenv.LoadEnv()
	basic := os.Getenv("BASIC_AUTHORIZATION")

	if err != nil {
		t.Error("Read .env file wrong.")
	}

	if basic != "Basic dXNlcjpwYXNz==" {
		t.Error("Read .env file wrong.")
	}
}

func TestLoadEnvByFile(t *testing.T) {
	err := goenv.LoadEnv(".testenv")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	if err != nil {
		t.Error("Read .testenv file wrong.")
	}

	if dbUser == "" && dbPass == "" {
		t.Error("Read .testenv file wrong.")
	}
}

func TestLoadEnvByFileNotFound(t *testing.T) {
	err := goenv.LoadEnv("7777")

	if err == nil {
		t.Error("Read 7777 file.")
	}
}
