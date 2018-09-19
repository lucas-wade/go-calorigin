package calorigin

// Sites
func (client *Client) GetSites() {}

func (client *Client) GetSiteBySiteID() {}

func (client *Client) GetSiteByLicense() {}

func (client *Client) GetCustomers() {}

func (client *Client) GetCustomersBySiteID() {}

func (client *Client) GetDisposalSites() {}

func (client *Client) GetTestLabs() {}

func (client *Client) GetDeliveryFromSites() {}

// Plantings
func (client *Client) CreatePlanting(planting Planting) {}

func (client *Client) GetPlantings() {}

func (client *Client) GetPlantingByID() {}

func (client *Client) UpdatePlanting() {}

func (client *Client) DeletePlanting() {}

// Harvests
func (client *Client) CreateHarvest(harvest Harvest) {}

func (client *Client) GetHarvests() {}

func (client *Client) GetHarvestByID() {}

func (client *Client) UpdateHarvest() {}

func (client *Client) DeleteHarvest() {}

// Products
func (client *Client) CreateProduct(product Product) {}

func (client *Client) GetProducts() {}

func (client *Client) GetProductByID() {}

func (client *Client) GetProductTypes() {}

func (client *Client) UpdateProduct() {}

func (client *Client) DeleteProduct() {}

// Activations
func (client *Client) CreateActivation(activation Activation) {}

func (client *Client) GetActivationBatches() {}

func (client *Client) GetActivationBatchByID() {}

// Inventory
func (client *Client) GetProductInventory() {}

func (client *Client) GetProductDetail() {}

func (client *Client) GetStampRollInventory() {}

func (client *Client) GetStampsByStampRoll() {}

// Transfers
func (client *Client) CreateTransfer(transfer Transfer) {}

func (client *Client) GetTransfersIncoming() {}

func (client *Client) GetTransferIncomingByTransferID() {}

func (client *Client) GetTransfersOutgoing() {}

func (client *Client) GetTransfersOutgoingByTransferID() {}

func (client *Client) UpdateTransfer() {}

func (client *Client) ConfirmTransfer() {}

func (client *Client) ReturnTransfer() {}

// Trace
func (client *Client) TraceStamp() {}
