package calorigin

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Sites
func (client *Client) GetSites() () {
}

func (client *Client) GetSiteBySiteID() {}

func (client *Client) GetSiteByLicense() {}

func (client *Client) GetCustomers() {}

func (client *Client) GetCustomersBySiteID() {}

func (client *Client) GetDisposalSites() {}

func (client *Client) GetTestLabs() {}

func (client *Client) GetDeliveryFromSites() {}

// Plantings
func (client *Client) CreatePlanting(token string, planting Planting) (*Planting, error) {
	url := client.URL + "/plantings/create"

	j, err := json.Marshal(planting)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	byteData, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	var data Planting
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (client *Client) GetPlantings() {}

func (client *Client) GetPlantingByID() {}

func (client *Client) UpdatePlanting() {}

func (client *Client) DeletePlanting() {}

// Harvests
func (client *Client) CreateHarvest(token string, harvest Harvest) (*Harvest, error) {
	url := client.URL + "/harvests/create"

	j, err := json.Marshal(harvest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	byteData, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	var data Harvest
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (client *Client) GetHarvests() {}

func (client *Client) GetHarvestByID() {}

func (client *Client) UpdateHarvest() {}

func (client *Client) DeleteHarvest() {}

// Products
func (client *Client) CreateProduct(token string, product Product) (*Product, error) {
	url := client.URL + "/products/create"

	j, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	byteData, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	var data Product
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (client *Client) GetProducts() {}

func (client *Client) GetProductByID() {}

func (client *Client) GetProductTypes() {}

func (client *Client) UpdateProduct() {}

func (client *Client) DeleteProduct() {}

// Activations
func (client *Client) CreateActivation(token string, activation Activation) (*Activation, error) {
	url := client.URL + "/activation"

	j, err := json.Marshal(activation)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	byteData, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	var data Activation
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (client *Client) GetActivationBatches() {}

func (client *Client) GetActivationBatchByID() {}

// Inventory
func (client *Client) GetProductInventory() {}

func (client *Client) GetProductDetail() {}

func (client *Client) GetStampRollInventory() {}

func (client *Client) GetStampsByStampRoll() {}

// Transfers
func (client *Client) CreateTransfer(token string, transfer Transfer) (*Transfer, error) {
	url := client.URL + "/transfers/create"

	j, err := json.Marshal(transfer)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	byteData, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	var data Transfer
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (client *Client) GetTransfersIncoming() {}

func (client *Client) GetTransferIncomingByTransferID() {}

func (client *Client) GetTransfersOutgoing() {}

func (client *Client) GetTransfersOutgoingByTransferID() {}

func (client *Client) UpdateTransfer() {}

func (client *Client) ConfirmTransfer() {}

func (client *Client) ReturnTransfer() {}

// Trace
func (client *Client) TraceStamp() {}
