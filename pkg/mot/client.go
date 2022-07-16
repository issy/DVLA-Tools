package mot

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const DVLA_URL = "https://beta.check-mot.service.gov.uk/trade/vehicles/mot-tests"

func makeUrl(queryParams map[string]string) string {
	myUrl, _ := url.Parse(DVLA_URL)
	values := myUrl.Query()
	for k, v := range queryParams {
		values.Add(k, v)
	}
	myUrl.RawQuery = values.Encode()
	return myUrl.String()
}

type Client struct {
	token string
}

func NewClient(token string) *Client {
	return &Client{token: token}
}

func (client *Client) setApiKeyHeader(header *http.Header) {
	header.Set("x-api-key", client.token)
}

func parseBody(body []byte) (CarResponse, error) {
	cars := CarResponse{}
	err := json.Unmarshal(body, &cars)
	if err == nil {
		return cars, nil
	}
	errorResponse := ErrorResponse{}
	err = json.Unmarshal(body, &errorResponse)
	if err != nil {
		return nil, err
	}
	return make(CarResponse, 0), nil
}

// Response is 200
func checkResponseIsOk(response *http.Response) bool {
	return response.StatusCode == 200
}

func (client *Client) makeRequest(queryParams map[string]string) CarResponse {
	req, err := http.NewRequest(http.MethodGet, makeUrl(queryParams), nil)
	if err != nil {
		log.Fatal(err)
	}

	client.setApiKeyHeader(&req.Header)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if !checkResponseIsOk(response) {
		log.Println(response.Body)
		log.Fatal("Received response with status not 200 OK")
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	cars, err := parseBody(body)
	if err != nil {
		log.Fatal(err)
	}
	return cars
}

func (client *Client) getCarsRaw(pageNum int) []byte {
	req, _ := http.NewRequest(http.MethodGet, makeUrl(map[string]string{"page": strconv.Itoa(pageNum)}), nil)
	client.setApiKeyHeader(&req.Header)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if checkResponseIsOk(response) == false {
		log.Fatal("Received response with status not 200 OK")
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func (client *Client) GetCar(registration string) (Car, error) {
	carResponse := client.makeRequest(map[string]string{"registration": strings.Replace(registration, " ", "", -1)})
	if len(carResponse) == 0 {
		return Car{}, errors.New("no car found")
	} else {
		return carResponse[0], nil
	}
}

func (client *Client) GetCars(pageNum int) (CarResponse, error) {
	carsBody := client.getCarsRaw(pageNum)
	cars, err := parseBody(carsBody)
	if err != nil {
		return CarResponse{}, err
	}
	return cars, nil
}
