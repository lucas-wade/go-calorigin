package calorigin

type Client struct {
	ClientID string
	URL      string
}

type Request struct {
}

type Planting struct {
	LicenseSiteID     int64   `json:"licenseSiteId"`
	StrainName        string  `json:"strainName"`
	LocationInSite    string  `json:"locationInSite"`
	InternalBatchID   string  `json:"internalBatchId"`
	LightTypeID       int     `json:"lightTypeId"`
	LightSubType      string  `json:"lightSubType,omitempty"`
	Pesticides        string  `json:"pesticides,omitempty"`
	PlantingDate      string  `json:"plantingDate"`
	PlantQty          int     `json:"plantQty"`
	PlannedYieldDate  string  `json:"plannedYieldDate"`
	EstimatedYieldLbs float64 `json:"estimatedYieldLbs"`
	Actual            bool    `json:"actual,omitempty"`
}

type Harvest struct {
	PlantingID     int     `json:"plantingId"`
	HarvestTypeID  int     `json:"harvestTypeId"`
	YieldDate      string  `json:"yieldDate"`
	WetYieldWeight float64 `json:"wetYieldWeight"`
	DryYieldWeight float64 `json:"dryYieldWeight,omitempty"`
	Location       string  `json:"location,omitempty"`
	Comments       string  `json:"comments,omitempty"`
}

type Product struct {
	LicenseSiteID   int     `json:"licenseSiteId"`
	ProductTypeID   int     `json:"productTypeId"`
	ProductName     string  `json:"productName"`
	Sku             string  `json:"sku"`
	Description     string  `json:"description"`
	Packaging       string  `json:"packaging"`
	PlantingQty     int     `json:"plantingQty,omitempty"`
	NetWeightInGram float64 `json:"netWeightInGram,omitempty"`
	ThcPct          float64 `json:"thcPct,omitempty"`
}

type Activation struct {
}

type Transfer struct {
}
