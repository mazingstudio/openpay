package openpay

// Merchant represents your Openpay account
type Merchant struct {
	id         string
	privateKey string
	client     *client
}

// NewMerchant initializes a merchant instance, which is at the same time an
// Openpay API consumer
func NewMerchant(id, privateKey string, sandbox bool) *Merchant {
	return &Merchant{
		id:         id,
		privateKey: privateKey,
		client:     newClient(id, privateKey, sandbox),
	}
}
