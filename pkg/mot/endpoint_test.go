package mot

import (
	"os"
	"testing"
)

func TestsEndpointReturnsData(t *testing.T) {
	client := NewClient(os.Getenv("DVLA_MOT_API_KEY"))
	car, err := client.GetCar("RJ06JOU")
	if err != nil {
		t.Fatal()
	}

}
