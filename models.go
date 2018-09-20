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
	ActivationSiteID    int     `json:"activationSiteId"`
	StampNoStart        string  `json:"stampNoStart"`
	StampNoEnd          string  `json:"stampNoEnd"`
	StampQty            int     `json:"stampQty"`
	ProductCatalogID    int     `json:"productCatalogId"`
	PlantingID          int     `json:"plantingId"`
	PlantQty            int     `json:"plantQty"`
	HarvestID           int     `json:"harvestId"`
	NetWtg              float64 `json:"netWtg"`
	Thc                 float64 `json:"thc"`
	SourceLicenseSiteID int     `json:"sourceLicenseSiteId"`
	SourceStampQty      int     `json:"sourceStampQty"`
	SourceStamps        []struct {
		StampNo    string  `json:"stampNo"`
		UsedNetWtg float64 `json:"usedNetWtg"`
	} `json:"sourceStamps"`
}

type Transfer struct {
}
