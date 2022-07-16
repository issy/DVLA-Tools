package ves

import (
	"context"
	"os"
	"testing"
)

func TestReturnsData(t *testing.T) {
	registration := "RJ06JOU"
	vehicleRequest := VehicleRequest{RegistrationNumber: &registration}
	xCorrelationId := "xCorrelationId_example"

	xApiKey := os.Getenv("DVLA_VES_API_KEY")
	configuration := NewConfiguration()
	apiClient := NewAPIClient(configuration)
	_, _, err := apiClient.VehicleApi.GetVehicleDetailsByRegistrationNumber(context.Background()).XApiKey(xApiKey).VehicleRequest(vehicleRequest).XCorrelationId(xCorrelationId).Execute()
	if err != nil {
		t.Fatal()
	}
}
