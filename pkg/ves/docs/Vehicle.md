# Vehicle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrationNumber** | **string** | Registration number of the vehicle | 
**TaxStatus** | Pointer to **string** | Tax status of the vehicle | [optional] 
**TaxDueDate** | Pointer to **string** | Date of tax liablity, Used in calculating licence information presented to user | [optional] 
**ArtEndDate** | Pointer to **string** | Additional Rate of Tax End Date, format: YYYY-MM-DD | [optional] 
**MotStatus** | Pointer to **string** | MOT Status of the vehicle | [optional] 
**MotExpiryDate** | Pointer to **string** | Mot Expiry Date | [optional] 
**Make** | Pointer to **string** | Vehicle make | [optional] 
**MonthOfFirstDvlaRegistration** | Pointer to **string** | Month of First DVLA Registration | [optional] 
**MonthOfFirstRegistration** | Pointer to **string** | Month of First Registration | [optional] 
**YearOfManufacture** | Pointer to **int32** | Year of Manufacture | [optional] 
**EngineCapacity** | Pointer to **int32** | Engine capacity in cubic centimetres | [optional] 
**Co2Emissions** | Pointer to **int32** | Carbon Dioxide emissions in grams per kilometre | [optional] 
**FuelType** | Pointer to **string** | Fuel type (Method of Propulsion) | [optional] 
**MarkedForExport** | Pointer to **bool** | True only if vehicle has been export marked | [optional] 
**Colour** | Pointer to **string** | Vehicle colour | [optional] 
**TypeApproval** | Pointer to **string** | Vehicle Type Approval Category | [optional] 
**Wheelplan** | Pointer to **string** | Vehicle wheel plan | [optional] 
**RevenueWeight** | Pointer to **int32** | Revenue weight in kilograms | [optional] 
**RealDrivingEmissions** | Pointer to **string** | Real Driving Emissions value | [optional] 
**DateOfLastV5CIssued** | Pointer to **string** | Date of last V5C issued | [optional] 
**EuroStatus** | Pointer to **string** | Euro Status (Dealer / Customer Provided (new vehicles)) | [optional] 
**AutomatedVehicle** | Pointer to **bool** | Automated Vehicle (AV) | [optional] 

## Methods

### NewVehicle

`func NewVehicle(registrationNumber string, ) *Vehicle`

NewVehicle instantiates a new Vehicle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleWithDefaults

`func NewVehicleWithDefaults() *Vehicle`

NewVehicleWithDefaults instantiates a new Vehicle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrationNumber

`func (o *Vehicle) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *Vehicle) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *Vehicle) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.


### GetTaxStatus

`func (o *Vehicle) GetTaxStatus() string`

GetTaxStatus returns the TaxStatus field if non-nil, zero value otherwise.

### GetTaxStatusOk

`func (o *Vehicle) GetTaxStatusOk() (*string, bool)`

GetTaxStatusOk returns a tuple with the TaxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxStatus

`func (o *Vehicle) SetTaxStatus(v string)`

SetTaxStatus sets TaxStatus field to given value.

### HasTaxStatus

`func (o *Vehicle) HasTaxStatus() bool`

HasTaxStatus returns a boolean if a field has been set.

### GetTaxDueDate

`func (o *Vehicle) GetTaxDueDate() string`

GetTaxDueDate returns the TaxDueDate field if non-nil, zero value otherwise.

### GetTaxDueDateOk

`func (o *Vehicle) GetTaxDueDateOk() (*string, bool)`

GetTaxDueDateOk returns a tuple with the TaxDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDueDate

`func (o *Vehicle) SetTaxDueDate(v string)`

SetTaxDueDate sets TaxDueDate field to given value.

### HasTaxDueDate

`func (o *Vehicle) HasTaxDueDate() bool`

HasTaxDueDate returns a boolean if a field has been set.

### GetArtEndDate

`func (o *Vehicle) GetArtEndDate() string`

GetArtEndDate returns the ArtEndDate field if non-nil, zero value otherwise.

### GetArtEndDateOk

`func (o *Vehicle) GetArtEndDateOk() (*string, bool)`

GetArtEndDateOk returns a tuple with the ArtEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtEndDate

`func (o *Vehicle) SetArtEndDate(v string)`

SetArtEndDate sets ArtEndDate field to given value.

### HasArtEndDate

`func (o *Vehicle) HasArtEndDate() bool`

HasArtEndDate returns a boolean if a field has been set.

### GetMotStatus

`func (o *Vehicle) GetMotStatus() string`

GetMotStatus returns the MotStatus field if non-nil, zero value otherwise.

### GetMotStatusOk

`func (o *Vehicle) GetMotStatusOk() (*string, bool)`

GetMotStatusOk returns a tuple with the MotStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotStatus

`func (o *Vehicle) SetMotStatus(v string)`

SetMotStatus sets MotStatus field to given value.

### HasMotStatus

`func (o *Vehicle) HasMotStatus() bool`

HasMotStatus returns a boolean if a field has been set.

### GetMotExpiryDate

`func (o *Vehicle) GetMotExpiryDate() string`

GetMotExpiryDate returns the MotExpiryDate field if non-nil, zero value otherwise.

### GetMotExpiryDateOk

`func (o *Vehicle) GetMotExpiryDateOk() (*string, bool)`

GetMotExpiryDateOk returns a tuple with the MotExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotExpiryDate

`func (o *Vehicle) SetMotExpiryDate(v string)`

SetMotExpiryDate sets MotExpiryDate field to given value.

### HasMotExpiryDate

`func (o *Vehicle) HasMotExpiryDate() bool`

HasMotExpiryDate returns a boolean if a field has been set.

### GetMake

`func (o *Vehicle) GetMake() string`

GetMake returns the Make field if non-nil, zero value otherwise.

### GetMakeOk

`func (o *Vehicle) GetMakeOk() (*string, bool)`

GetMakeOk returns a tuple with the Make field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMake

`func (o *Vehicle) SetMake(v string)`

SetMake sets Make field to given value.

### HasMake

`func (o *Vehicle) HasMake() bool`

HasMake returns a boolean if a field has been set.

### GetMonthOfFirstDvlaRegistration

`func (o *Vehicle) GetMonthOfFirstDvlaRegistration() string`

GetMonthOfFirstDvlaRegistration returns the MonthOfFirstDvlaRegistration field if non-nil, zero value otherwise.

### GetMonthOfFirstDvlaRegistrationOk

`func (o *Vehicle) GetMonthOfFirstDvlaRegistrationOk() (*string, bool)`

GetMonthOfFirstDvlaRegistrationOk returns a tuple with the MonthOfFirstDvlaRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthOfFirstDvlaRegistration

`func (o *Vehicle) SetMonthOfFirstDvlaRegistration(v string)`

SetMonthOfFirstDvlaRegistration sets MonthOfFirstDvlaRegistration field to given value.

### HasMonthOfFirstDvlaRegistration

`func (o *Vehicle) HasMonthOfFirstDvlaRegistration() bool`

HasMonthOfFirstDvlaRegistration returns a boolean if a field has been set.

### GetMonthOfFirstRegistration

`func (o *Vehicle) GetMonthOfFirstRegistration() string`

GetMonthOfFirstRegistration returns the MonthOfFirstRegistration field if non-nil, zero value otherwise.

### GetMonthOfFirstRegistrationOk

`func (o *Vehicle) GetMonthOfFirstRegistrationOk() (*string, bool)`

GetMonthOfFirstRegistrationOk returns a tuple with the MonthOfFirstRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthOfFirstRegistration

`func (o *Vehicle) SetMonthOfFirstRegistration(v string)`

SetMonthOfFirstRegistration sets MonthOfFirstRegistration field to given value.

### HasMonthOfFirstRegistration

`func (o *Vehicle) HasMonthOfFirstRegistration() bool`

HasMonthOfFirstRegistration returns a boolean if a field has been set.

### GetYearOfManufacture

`func (o *Vehicle) GetYearOfManufacture() int32`

GetYearOfManufacture returns the YearOfManufacture field if non-nil, zero value otherwise.

### GetYearOfManufactureOk

`func (o *Vehicle) GetYearOfManufactureOk() (*int32, bool)`

GetYearOfManufactureOk returns a tuple with the YearOfManufacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearOfManufacture

`func (o *Vehicle) SetYearOfManufacture(v int32)`

SetYearOfManufacture sets YearOfManufacture field to given value.

### HasYearOfManufacture

`func (o *Vehicle) HasYearOfManufacture() bool`

HasYearOfManufacture returns a boolean if a field has been set.

### GetEngineCapacity

`func (o *Vehicle) GetEngineCapacity() int32`

GetEngineCapacity returns the EngineCapacity field if non-nil, zero value otherwise.

### GetEngineCapacityOk

`func (o *Vehicle) GetEngineCapacityOk() (*int32, bool)`

GetEngineCapacityOk returns a tuple with the EngineCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineCapacity

`func (o *Vehicle) SetEngineCapacity(v int32)`

SetEngineCapacity sets EngineCapacity field to given value.

### HasEngineCapacity

`func (o *Vehicle) HasEngineCapacity() bool`

HasEngineCapacity returns a boolean if a field has been set.

### GetCo2Emissions

`func (o *Vehicle) GetCo2Emissions() int32`

GetCo2Emissions returns the Co2Emissions field if non-nil, zero value otherwise.

### GetCo2EmissionsOk

`func (o *Vehicle) GetCo2EmissionsOk() (*int32, bool)`

GetCo2EmissionsOk returns a tuple with the Co2Emissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCo2Emissions

`func (o *Vehicle) SetCo2Emissions(v int32)`

SetCo2Emissions sets Co2Emissions field to given value.

### HasCo2Emissions

`func (o *Vehicle) HasCo2Emissions() bool`

HasCo2Emissions returns a boolean if a field has been set.

### GetFuelType

`func (o *Vehicle) GetFuelType() string`

GetFuelType returns the FuelType field if non-nil, zero value otherwise.

### GetFuelTypeOk

`func (o *Vehicle) GetFuelTypeOk() (*string, bool)`

GetFuelTypeOk returns a tuple with the FuelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelType

`func (o *Vehicle) SetFuelType(v string)`

SetFuelType sets FuelType field to given value.

### HasFuelType

`func (o *Vehicle) HasFuelType() bool`

HasFuelType returns a boolean if a field has been set.

### GetMarkedForExport

`func (o *Vehicle) GetMarkedForExport() bool`

GetMarkedForExport returns the MarkedForExport field if non-nil, zero value otherwise.

### GetMarkedForExportOk

`func (o *Vehicle) GetMarkedForExportOk() (*bool, bool)`

GetMarkedForExportOk returns a tuple with the MarkedForExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForExport

`func (o *Vehicle) SetMarkedForExport(v bool)`

SetMarkedForExport sets MarkedForExport field to given value.

### HasMarkedForExport

`func (o *Vehicle) HasMarkedForExport() bool`

HasMarkedForExport returns a boolean if a field has been set.

### GetColour

`func (o *Vehicle) GetColour() string`

GetColour returns the Colour field if non-nil, zero value otherwise.

### GetColourOk

`func (o *Vehicle) GetColourOk() (*string, bool)`

GetColourOk returns a tuple with the Colour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColour

`func (o *Vehicle) SetColour(v string)`

SetColour sets Colour field to given value.

### HasColour

`func (o *Vehicle) HasColour() bool`

HasColour returns a boolean if a field has been set.

### GetTypeApproval

`func (o *Vehicle) GetTypeApproval() string`

GetTypeApproval returns the TypeApproval field if non-nil, zero value otherwise.

### GetTypeApprovalOk

`func (o *Vehicle) GetTypeApprovalOk() (*string, bool)`

GetTypeApprovalOk returns a tuple with the TypeApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeApproval

`func (o *Vehicle) SetTypeApproval(v string)`

SetTypeApproval sets TypeApproval field to given value.

### HasTypeApproval

`func (o *Vehicle) HasTypeApproval() bool`

HasTypeApproval returns a boolean if a field has been set.

### GetWheelplan

`func (o *Vehicle) GetWheelplan() string`

GetWheelplan returns the Wheelplan field if non-nil, zero value otherwise.

### GetWheelplanOk

`func (o *Vehicle) GetWheelplanOk() (*string, bool)`

GetWheelplanOk returns a tuple with the Wheelplan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWheelplan

`func (o *Vehicle) SetWheelplan(v string)`

SetWheelplan sets Wheelplan field to given value.

### HasWheelplan

`func (o *Vehicle) HasWheelplan() bool`

HasWheelplan returns a boolean if a field has been set.

### GetRevenueWeight

`func (o *Vehicle) GetRevenueWeight() int32`

GetRevenueWeight returns the RevenueWeight field if non-nil, zero value otherwise.

### GetRevenueWeightOk

`func (o *Vehicle) GetRevenueWeightOk() (*int32, bool)`

GetRevenueWeightOk returns a tuple with the RevenueWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueWeight

`func (o *Vehicle) SetRevenueWeight(v int32)`

SetRevenueWeight sets RevenueWeight field to given value.

### HasRevenueWeight

`func (o *Vehicle) HasRevenueWeight() bool`

HasRevenueWeight returns a boolean if a field has been set.

### GetRealDrivingEmissions

`func (o *Vehicle) GetRealDrivingEmissions() string`

GetRealDrivingEmissions returns the RealDrivingEmissions field if non-nil, zero value otherwise.

### GetRealDrivingEmissionsOk

`func (o *Vehicle) GetRealDrivingEmissionsOk() (*string, bool)`

GetRealDrivingEmissionsOk returns a tuple with the RealDrivingEmissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealDrivingEmissions

`func (o *Vehicle) SetRealDrivingEmissions(v string)`

SetRealDrivingEmissions sets RealDrivingEmissions field to given value.

### HasRealDrivingEmissions

`func (o *Vehicle) HasRealDrivingEmissions() bool`

HasRealDrivingEmissions returns a boolean if a field has been set.

### GetDateOfLastV5CIssued

`func (o *Vehicle) GetDateOfLastV5CIssued() string`

GetDateOfLastV5CIssued returns the DateOfLastV5CIssued field if non-nil, zero value otherwise.

### GetDateOfLastV5CIssuedOk

`func (o *Vehicle) GetDateOfLastV5CIssuedOk() (*string, bool)`

GetDateOfLastV5CIssuedOk returns a tuple with the DateOfLastV5CIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfLastV5CIssued

`func (o *Vehicle) SetDateOfLastV5CIssued(v string)`

SetDateOfLastV5CIssued sets DateOfLastV5CIssued field to given value.

### HasDateOfLastV5CIssued

`func (o *Vehicle) HasDateOfLastV5CIssued() bool`

HasDateOfLastV5CIssued returns a boolean if a field has been set.

### GetEuroStatus

`func (o *Vehicle) GetEuroStatus() string`

GetEuroStatus returns the EuroStatus field if non-nil, zero value otherwise.

### GetEuroStatusOk

`func (o *Vehicle) GetEuroStatusOk() (*string, bool)`

GetEuroStatusOk returns a tuple with the EuroStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEuroStatus

`func (o *Vehicle) SetEuroStatus(v string)`

SetEuroStatus sets EuroStatus field to given value.

### HasEuroStatus

`func (o *Vehicle) HasEuroStatus() bool`

HasEuroStatus returns a boolean if a field has been set.

### GetAutomatedVehicle

`func (o *Vehicle) GetAutomatedVehicle() bool`

GetAutomatedVehicle returns the AutomatedVehicle field if non-nil, zero value otherwise.

### GetAutomatedVehicleOk

`func (o *Vehicle) GetAutomatedVehicleOk() (*bool, bool)`

GetAutomatedVehicleOk returns a tuple with the AutomatedVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedVehicle

`func (o *Vehicle) SetAutomatedVehicle(v bool)`

SetAutomatedVehicle sets AutomatedVehicle field to given value.

### HasAutomatedVehicle

`func (o *Vehicle) HasAutomatedVehicle() bool`

HasAutomatedVehicle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


