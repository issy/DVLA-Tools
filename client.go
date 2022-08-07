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

type VehicleResponse struct {
	Registration      string        `json:"registration"`
	YearOfManufacture int32         `json:"yearOfManufacture"`
	Make              string        `json:"make"`
	Model             string        `json:"model"`
	Colour            string        `json:"colour"`
	FuelType          string        `json:"fuelType"`
	EngineCapacity    int32         `json:"engineCapacity"`
	MotStatus         string        `json:"motStatus"`
	MotExpiryDate     string        `json:"motExpiryDate"`
	MotTests          []mot.MotTest `json:"motTests"`
	TaxStatus         string        `json:"taxStatus"`
	TaxDueDate        string        `json:"taxDueDate"`
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

func makeVehicleResponse(motResponse *mot.Car, vesResponse *ves.Vehicle) *VehicleResponse {
	return &VehicleResponse{
		Registration:      motResponse.Registration,
		YearOfManufacture: vesResponse.GetYearOfManufacture(),
		Make:              vesResponse.GetMake(),
		Model:             motResponse.Model,
		Colour:            vesResponse.GetColour(),
		FuelType:          vesResponse.GetFuelType(),
		EngineCapacity:    vesResponse.GetEngineCapacity(),
		MotStatus:         vesResponse.GetMotStatus(),
		MotExpiryDate:     vesResponse.GetMotExpiryDate(),
		MotTests:          motResponse.MotTests,
		TaxStatus:         vesResponse.GetTaxStatus(),
		TaxDueDate:        vesResponse.GetTaxDueDate(),
	}
}

func (client *ApiClient) GetVehicle(registration string) any {
	vesResponse, err := client.vesRequest(registration)
	if err != nil {
		panic(err)
	}
	motResponse, err := client.motRequest(registration)
	if err != nil {
		panic(err)
	}
	return makeVehicleResponse(motResponse, vesResponse)
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
