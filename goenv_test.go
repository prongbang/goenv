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
