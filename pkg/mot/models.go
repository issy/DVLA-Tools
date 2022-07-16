package mot

type ErrorResponse struct {
	HttpStatus   string `json:"httpStatus"`
	ErrorMessage string `json:"errorMessage"`
	AwsRequestId string `json:"awsRequestId"`
}

type MotTest struct {
	CompletedDate  string `json:"completedDate"`
	TestResult     string `json:"testResult"`
	ExpiryDate     string `json:"expiryDate"`
	OdometerValue  string `json:"odometerValue"`
	OdometerUnit   string `json:"odometerUnit"`
	MotTestNumber  string `json:"motTestNumber"`
	RfrAndComments []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"rfrAndComments"`
}

type Car struct {
	Registration  string    `json:"registration"`
	Make          string    `json:"make"`
	Model         string    `json:"model"`
	FirstUsedDate string    `json:"firstUsedDate"`
	FuelType      string    `json:"fuelType"`
	PrimaryColour string    `json:"primaryColour"`
	MotTests      []MotTest `json:"motTests"`
}

type CarResponse []Car
