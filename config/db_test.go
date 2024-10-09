package config

import (
	"os"
	"testing"
)

func TestInitConfig(t *testing.T) {
	err := InitConfig()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if logger == nil {
		t.Fatalf("Expected logger to be initialized")
	}
}

func TestConnectDB(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpassword")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "testdb")

	err := ConnectDB()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if DB == nil {
		t.Fatalf("Expected DB to be initialized")
	}

	// Clean up
	CloseDB()
}

func TestCloseDB(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpassword")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "testdb")

	err := ConnectDB()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	CloseDB()

	if err := DB.Ping(); err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestGetDB(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpassword")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "testdb")

	err := ConnectDB()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	db := GetDB()
	if db == nil {
		t.Fatalf("Expected DB to be initialized")
	}

	// Clean up
	CloseDB()
}

func TestLogger(t *testing.T) {
	err := InitConfig()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	log := Logger()
	if log == nil {
		t.Fatalf("Expected logger to be initialized")
	}
}
