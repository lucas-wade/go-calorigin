package calorigin

func NewCalOriginClient(clientID, url string) *Client {
	return &Client{
		ClientID: clientID,
		URL:      url,
	}
}
