package go_dvla

import (
	"context"
	"github.com/issy/go-dvla/pkg/mot"
	"github.com/issy/go-dvla/pkg/ves"
)

type ApiClient struct {
	motClient *mot.Client
	vesClient *ves.APIClient
}

type Vehicle struct {
	Ves *ves.Vehicle `json:"ves"`
	Mot *mot.Car     `json:"mot"`
}

func NewClient(motApiToken string, vesApiToken string) *ApiClient {
	configuration := ves.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", vesApiToken)
	configuration.AddDefaultHeader("x-correlation-id", "go-dvla client")

	return &ApiClient{
		motClient: mot.NewClient(motApiToken),
		vesClient: ves.NewAPIClient(configuration),
	}
}

func (client *ApiClient) GetVehicle(registration string) *Vehicle {
	vesResponse, err := client.vesRequest(registration)
	if err != nil {
		panic(err)
	}
	motResponse, err := client.motRequest(registration)
	if err != nil {
		panic(err)
	}
	return &Vehicle{Ves: vesResponse, Mot: motResponse}
}

func (client *ApiClient) vesRequest(registration string) (*ves.Vehicle, error) {
	vehicle, _, err := client.vesClient.VehicleApi.GetVehicleDetailsByRegistrationNumber(context.Background()).VehicleRequest(ves.VehicleRequest{RegistrationNumber: &registration}).Execute()
	if err != nil {
		return &ves.Vehicle{}, err
	} else {
		return vehicle, nil
	}
}

func (client *ApiClient) motRequest(registration string) (*mot.Car, error) {
	car, err := client.motClient.GetCar(registration)
	if err != nil {
		return &mot.Car{}, err
	} else {
		return &car, nil
	}
}
