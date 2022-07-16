package mot

import (
	"os"
	"testing"
)

func TestEndpointReturnsData(t *testing.T) {
	client := NewClient(os.Getenv("DVLA_MOT_API_KEY"))
	_, err := client.GetCar("RJ06JOU")
	if err != nil {
		t.Fatal()
	}

}
