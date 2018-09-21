package calorigin

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (client *Client) GetSites(token string) ([]byte, error) {
	url := client.URL + "/sites"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetSiteBySiteID(token string, siteID int) ([]byte, error)  {
	url := client.URL + "/sites/" + string(siteID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetSiteByLicense(token string, licenseNumber string) ([]byte, error)  {
	url := client.URL + "/sites/license/" + licenseNumber

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetCustomers(token string) ([]byte, error)  {
	url := client.URL + "/sites/customers"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetCustomersBySiteID(token string, siteID int) ([]byte, error)  {
	url := client.URL + "/sites/" + string(siteID) + "/customers"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetDisposalSites(token string) ([]byte, error)  {
	url := client.URL + "/sites/disposal"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetTestLabs(token string) ([]byte, error)  {
	url := client.URL + "/sites/testlabs"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetDeliveryFromSites(token string, siteID int) ([]byte, error)  {
	url := client.URL + "/sites/" + string(siteID) + "/from"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// CreatePlanting will create a planting and return the response.
func (client *Client) CreatePlanting(token string, planting CreatePlantingRequest) ([]byte, error) {
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

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetPlantings(token string) ([]byte, error)  {
	url := client.URL + "/plantings"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetPlantingByID(token string, plantingID int) ([]byte, error)  {
	url := client.URL + "/planting/" + string(plantingID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetPlantingLightTypes(token string) ([]byte, error) {
	url := client.URL + "/plantings/light_types"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) UpdatePlanting(token string, planting UpdatePlantingRequest) ([]byte, error)  {
	url := client.URL + "/plantings/edit"

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

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) DeletePlanting(token string, planting DeletePlantingRequest) ([]byte, error)  {
	url := client.URL + "/plantings/delete"

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

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// CreateHarvest will create a harvest and return the response.
func (client *Client) CreateHarvest(token string, harvest CreateHarvestRequest) ([]byte, error) {
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

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetHarvests(token string) ([]byte, error)  {
	url := client.URL + "/harvests"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetHarvestByID(token string, harvestID int) ([]byte, error)  {
	url := client.URL + "/harvests/" + string(harvestID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) UpdateHarvest(token string, harvest UpdateHarvestRequest) ([]byte, error)  {
	url := client.URL + "/harvests/edit"

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

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) DeleteHarvest(token string, harvest DeleteHarvestRequest) ([]byte, error)  {
	url := client.URL + "/plantings/delete"

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

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// CreateProduct will create a product and return the response
func (client *Client) CreateProduct(token string, product CreateProductRequest) ([]byte, error) {
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

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetProducts(token string) ([]byte, error)  {
	url := client.URL + "/products"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetProductByID(token string, productID int) ([]byte, error)  {
	url := client.URL + "/products/" + string(productID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetProductTypes(token string) ([]byte, error)  {
	url := client.URL + "/products/types"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) UpdateProduct(token string, product UpdateProductRequest) ([]byte, error)  {
	url := client.URL + "/products/edit"

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

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) DeleteProduct(token string, product DeleteProductRequest) ([]byte, error)  {
	url := client.URL + "/products/delete"

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

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) CreateActivation(token string, activation CreateActivationRequest) ([]byte, error) {
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

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetActivationBatches(token string) ([]byte, error)  {
	url := client.URL + "/activation/batches"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetActivationBatchByID(token string, batchID int) ([]byte, error)  {
	url := client.URL + "/activations/batches/" + string(batchID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetProductInventory(token string) ([]byte, error)  {
	url := client.URL + "/inventory/product"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetProductDetail(token string, siteID, productID int) ([]byte, error)  {
	url := client.URL + "/inventory/product/" + string(siteID) + "/" + string(productID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetStampRollInventory(token string) ([]byte, error)  {
	url := client.URL + "/inventory/stamp"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetStampsByStampRoll(token string, rollNumber int) ([]byte, error)  {
	url := client.URL + "/inventory/stamp/" + string(rollNumber)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) CreateTransfer(token string, transfer CreateTransferRequest) ([]byte, error) {
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

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetTransfersIncoming(token string) ([]byte, error)  {
	url := client.URL + "/transfers/incoming"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetTransferIncomingByTransferID(token string, transferNumber int) ([]byte, error) {
	url := client.URL + "/transfers/incoming/" + string(transferNumber)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetTransfersOutgoing(token string) ([]byte, error)  {
	url := client.URL + "/transfers/outgoing"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) GetTransfersOutgoingByTransferID(token string, transferNumber int) ([]byte, error)  {
	url := client.URL + "/transfers/outgoing/" + string(transferNumber)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) UpdateTransfer(token string, transfer UpdateTransferRequest) ([]byte, error)  {
	url := client.URL + "/transfers/edit"

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

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) ConfirmTransfer(token string, transfer ConfirmTransferRequest) ([]byte, error)  {
	url := client.URL + "/transfers/confirm"

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

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) ReturnTransfer(token string, transfer ReturnTransferRequest) ([]byte, error)  {
	url := client.URL + "/transfers/return"

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

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *Client) TraceStamp(token string, stampNumber string) ([]byte, error)  {
	url := client.URL + "/trace/" + string(stampNumber)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-ibm-client-id", client.ClientID)
	req.Header.Add("authorization", token)

	data, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}
