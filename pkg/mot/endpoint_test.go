package mot

import "testing"

func TestsEndpointReturnsData(t *testing.T) {
	client := NewClient(readMOTApiKeyFromEnvVar())
	car, err := client.GetCar(registration)
	if err != nil {
		t.Fatal()
	}

}
