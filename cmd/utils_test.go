package cmd

import "testing"

func TestsVESAPIKey(t *testing.T) {
	apiKey := readVESApiKeyFromEnvVar()
	if apiKey == "" {
		t.Fatal("API key not set")
	}
}

func TestsMOTAPIKey(t *testing.T) {
	apiKey := readMOTApiKeyFromEnvVar()
	if apiKey == "" {
		t.Fatal("API key not set")
	}
}
