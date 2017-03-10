package openpay

import "net/http"

const (
	productionURL = "https://api.openpay.mx/v1"
	sandboxURL    = "https://sandbox-api.openpay.mx/v1"
)

// client is an Openpay client
type client struct {
	merchantID string
	privateKey string
	apiBase    string
	client     *http.Client
}

// NewClient initializes a client instance with the specified merchant ID and
// private key.
func newClient(merchantID, privateKey string, sandbox bool) *client {
	url := productionURL
	if sandbox {
		url = sandboxURL
	}
	return &client{
		merchantID: merchantID,
		privateKey: privateKey,
		apiBase:    url,
		client:     &http.Client{},
	}
}
