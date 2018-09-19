package calorigin

// Sites
func GetSites() {}

func GetSiteBySiteID() {}

func GetSiteByLicense() {}

func GetCustomers() {}

func GetCustomersBySiteID() {}

func GetDisposalSites() {}

func GetTestLabs() {}

func GetDeliveryFromSites() {}

// Plantings
func CreatePlanting() {}

func GetPlantings() {}

func GetPlantingByID() {}

func UpdatePlanting() {}

func DeletePlanting() {}

// Harvests
func CreateHarvest() {}

func GetHarvests() {}

func GetHarvestByID() {}

func UpdateHarvest() {}

func DeleteHarvest() {}

// Products
func CreateProduct() {}

func GetProducts() {}

func GetProductByID() {}

func GetProductTypes() {}

func UpdateProduct() {}

func DeleteProduct() {}

// Activations
func CreateActivation() {}

func GetActivationBatches() {}

func GetActivationBatchByID() {}

// Inventory
func GetProductInventory() {}

func GetProductDetail() {}

func GetStampRollInventory() {}

func GetStampsByStampRoll() {}

// Transfers
func CreateTransfer() {}

func GetTransfersIncoming() {}

func GetTransferIncomingByTransferID() {}

func GetTransfersOutgoing() {}

func GetTransfersOutgoingByTransferID() {}

func UpdateTransfer() {}

func ConfirmTransfer() {}

func ReturnTransfer() {}

// Trace
func TraceStamp() {}