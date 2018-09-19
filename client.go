package calorigin

func NewCalOriginClient(clientID string) *Client {
	return &Client{
		ClientID: clientID,
	}
}
