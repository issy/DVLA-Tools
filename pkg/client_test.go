package go_dvla

import (
	"os"
	"testing"
)

func TestReturnsData(t *testing.T) {
	vesApiKey := os.Getenv("DVLA_VES_API_KEY")
	motApiKey := os.Getenv("DVLA_MOT_API_KEY")
	client := NewClient(motApiKey, vesApiKey)
	registration := "RJ06JOU"
	client.GetVehicle(registration)
}
