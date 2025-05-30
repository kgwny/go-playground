package main

import (
	"strings"
	"testing"
)

// モックロガー　(呼び出しログを保持)
type MockLogger struct {
	Messages []string
}

func (m *MockLogger) Info(msg string) {
	m.Messages = append(m.Messages, "[INFO] "+msg)
}

func (m *MockLogger) Error(msg string) {
	m.Messages = append(m.Messages, "[ERROR] "+msg)
}

func TestNewUser_Success(t *testing.T) {
	logger := &MockLogger{}

	user, err := NewUser("Alice", 30, logger)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.Name != "Alice" || user.Age != 30 {
		t.Errorf("unexpected user fields: %+v", user)
	}

	if len(logger.Messages) == 0 || !strings.Contains(logger.Messages[0], "creating new user") {
		t.Errorf("expected info log not foud: %v", logger.Messages)
	}
}

func TestNewUser_InvalidAge(t *testing.T) {
	logger := &MockLogger{}

	user, err := NewUser("Bob", -5, logger)
	if err == nil {
		t.Fatal("expected error for negative age, got nil")
	}

	if user != nil {
		t.Errorf("expected nil user, got %+v", user)
	}

	foundErrorLog := false
	for _, msg := range logger.Messages {
		if strings.Contains(msg, "invalid age") {
			foundErrorLog = true
			break
		}
	}

	if !foundErrorLog {
		t.Errorf("expected error log message, got: %v", logger.Messages)
	}
}
