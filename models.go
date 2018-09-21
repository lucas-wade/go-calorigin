package calorigin

type Client struct {
	ClientID string
	URL      string
}

type CreatePlantingRequest struct {
	LicenseSiteID     int   `json:"licenseSiteId"`
	StrainName        string  `json:"strainName"`
	LocationInSite    string  `json:"locationInSite"`
	InternalBatchID   string  `json:"internalBatchId"`
	LightTypeID       int   `json:"lightTypeId"`
	LightSubType      string  `json:"lightSubType,omitempty"`
	Pesticides        string  `json:"pesticides,omitempty"`
	PlantingDate      string  `json:"plantingDate"`
	PlantQty          int   `json:"plantQty"`
	PlannedYieldDate  string  `json:"plannedYieldDate"`
	EstimatedYieldLbs float64 `json:"estimatedYieldLbs"`
	Actual            bool    `json:"actual,omitempty"`
}

type UpdatePlantingRequest struct {
	PlantingID        int   `json:"plantingId"`
	LicenseSiteID     int   `json:"licenseSiteId"`
	StrainName        string  `json:"strainName"`
	LocationInSite    string  `json:"locationInSite"`
	InternalBatchID   string  `json:"internalBatchId"`
	LightTypeID       int   `json:"lightTypeId"`
	LightSubType      string  `json:"lightSubType,omitempty"`
	Pesticides        string  `json:"pesticides,omitempty"`
	PlantingDate      string  `json:"plantingDate"`
	PlantQty          int   `json:"plantQty"`
	PlannedYieldDate  string  `json:"plannedYieldDate"`
	EstimatedYieldLbs float64 `json:"estimatedYieldLbs"`
	Actual            bool    `json:"actual,omitempty"`
	RevisionReason    string  `json:"revisionReason"`
}

type DeletePlantingRequest struct {
	PlantingID    int  `json:"plantingId"`
	LicenseSiteID int  `json:"licenseSiteId"`
	StrainName    string `json:"strainName"`
	PlantingDate  string `json:"plantingDate"`
	PlantQty      string `json:"plantQty"`
}

type CreateHarvestRequest struct {
	PlantingID     int   `json:"plantingId"`
	HarvestTypeID  int   `json:"harvestTypeId"`
	YieldDate      string  `json:"yieldDate"`
	WetYieldWeight float64 `json:"wetYieldWeight"`
	DryYieldWeight float64 `json:"dryYieldWeight,omitempty"`
	Location       string  `json:"location,omitempty"`
	Comments       string  `json:"comments,omitempty"`
}

type UpdateHarvestRequest struct {
	HarvestID      int     `json:"harvestId"`
	PlantingID     int     `json:"plantingId"`
	HarvestTypeID  int     `json:"harvestTypeId"`
	YieldDate      string  `json:"yieldDate"`
	WetYieldWeight float64 `json:"wetYieldWeight"`
	DryYieldWeight float64 `json:"dryYieldWeight,omitempty"`
	Location       string  `json:"location,omitempty"`
	Comments       string  `json:"comments,omitempty"`
}

type DeleteHarvestRequest struct {
	HarvestID  int    `json:"harvestId"`
	PlantingID int    `json:"plantingId"`
	YieldDate  string `json:"yieldDate"`
}

type CreateProductRequest struct {
	LicenseSiteID   int   `json:"licenseSiteId"`
	ProductTypeID   int   `json:"productTypeId"`
	ProductName     string  `json:"productName"`
	Sku             string  `json:"sku"`
	Description     string  `json:"description"`
	Packaging       string  `json:"packaging"`
	PlantingQty     int   `json:"plantingQty,omitempty"`
	NetWeightInGram float64 `json:"netWeightInGram,omitempty"`
	ThcPct          float64 `json:"thcPct,omitempty"`
}

type UpdateProductRequest struct {
	ProductCatalogID int     `json:"productCatalogId"`
	LicenseSiteID    int     `json:"licenseSiteId"`
	ProductTypeID    int     `json:"productTypeId"`
	ProductName      string  `json:"productName"`
	Sku              string  `json:"sku,omitempty"`
	Description      string  `json:"description"`
	Packaging        string  `json:"packaging"`
	PlantQty         int     `json:"plantQty,omitempty"`
	NetWeightInGram  float64 `json:"netWeightInGram,omitempty"`
	ThcPct           float64 `json:"thcPct,omitempty"`
}

type DeleteProductRequest struct {
	ProductCatalogID int    `json:"productCatalogId"`
	LicenseSiteID    int    `json:"licenseSiteId"`
	ProductName      string `json:"productName"`
}

type CreateActivationRequest struct {
	ActivationSiteID    int   `json:"activationSiteId"`
	StampNoStart        string  `json:"stampNoStart"`
	StampNoEnd          string  `json:"stampNoEnd"`
	StampQty            int   `json:"stampQty"`
	ProductCatalogID    int   `json:"productCatalogId"`
	PlantingID          int   `json:"plantingId"`
	PlantQty            int   `json:"plantQty"`
	HarvestID           int   `json:"harvestId"`
	NetWtg              float64 `json:"netWtg"`
	Thc                 float64 `json:"thc"`
	SourceLicenseSiteID int   `json:"sourceLicenseSiteId"`
	SourceStampQty      int   `json:"sourceStampQty"`
	SourceStamps        []struct {
		StampNo    string  `json:"stampNo"`
		UsedNetWtg float64 `json:"usedNetWtg"`
	} `json:"sourceStamps"`
}

type CreateTransferRequest struct {
	TransferFromSiteID int    `json:"transferFromSiteId"`
	TransferToSiteID   int    `json:"transferToSiteId"`
	TransportAgent     string   `json:"transportAgent,omitempty"`
	LicensePlate       string   `json:"licensePlate,omitempty"`
	DriverName         string   `json:"driverName,omitempty"`
	DriverLicenseNo    string   `json:"driverLicenseNo,omitempty"`
	VehicleMake        string   `json:"vehicleMake,omitempty"`
	VehicleModel       string   `json:"vehicleModel,omitempty"`
	PhoneNumber        string   `json:"phoneNumber,omitempty"`
	StampQty           int    `json:"stampQty"`
	DepartureDate      string   `json:"departureDate"`
	DepartureTime      string   `json:"departureTime"`
	ArrivalTime        string   `json:"arrivalTime"`
	AdditionalInfo     string   `json:"additionalInfo"`
	StampCodes         []string `json:"stampCodes"`
}

type UpdateTransferRequest struct {
	TransferNo         int    `json:"transferNo"`
	TransferFromSiteID int    `json:"transferFromSiteId"`
	TransferToSiteID   int    `json:"transferToSiteId"`
	StampQty           int    `json:"stampQty"`
	TransportAgent     string `json:"transportAgent,omitempty"`
	LicensePlate       string `json:"licensePlate,omitempty"`
	DriverName         string `json:"driverName,omitempty"`
	DriverLicenseNo    string `json:"driverLicenseNo,omitempty"`
	VehicleMake        string `json:"vehicleMake,omitempty"`
	VehicleModel       string `json:"vehicleModel,omitempty"`
	PhoneNumber        string `json:"phoneNumber,omitempty"`
	DepartureDate      string `json:"departureDate"`
	DepartureTime      string `json:"departureTime"`
	ArrivalDate        string `json:"arrivalDate"`
	ArrivalTime        string `json:"arrivalTime"`
	AdditionalInfo     string `json:"additionalInfo,omitempty"`
}

type ConfirmTransferRequest struct {
	TransferNo         int    `json:"transferNo"`
	TransferFromSiteID int    `json:"transferFromSiteId"`
	TransferToSiteID   int    `json:"transferToSiteId"`
	StampQty           int    `json:"stampQty"`
	ArrivalDate        string `json:"arrivalDate"`
	ArrivalTime        string `json:"arrivalTime"`
}

type ReturnTransferRequest struct {
	ReturnTransferNo   int      `json:"returnTransferNo"`
	TransferFromSiteID int      `json:"transferFromSiteId"`
	TransferToSiteID   int      `json:"transferToSiteId"`
	TransportAgent     string   `json:"transportAgent,omitempty"`
	LicensePlate       string   `json:"licensePlate,omitempty"`
	DriverName         string   `json:"driverName,omitempty"`
	DriverLicenseNo    string   `json:"driverLicenseNo,omitempty"`
	VehicleMake        string   `json:"vehicleMake,omitempty"`
	VehicleModel       string   `json:"vehicleModel,omitempty"`
	PhoneNumber        string   `json:"phoneNumber,omitempty"`
	StampQty           int      `json:"stampQty"`
	DepartureDate      string   `json:"departureDate"`
	DepartureTime      string   `json:"departureTime"`
	ArrivalDate        string   `json:"arrivalDate"`
	ArrivalTime        string   `json:"arrivalTime"`
	AdditionalInfo     string   `json:"additionalInfo,omitempty"`
	StampCodes         []string `json:"stampCodes"`
}
