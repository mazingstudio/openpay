package openpay

import (
	"fmt"
	"time"
)

// Card represents a credit or debit card.
type Card struct {
	CardArgs
	CreationDate  time.Time `json:"creation_date,omitempty"`
	AllowsCharges bool      `json:"allows_charges"`
	AllowsPayouts bool      `json:"allows_payouts"`
	Brand         string    `json:"brand,omitempty"`
	Type          string    `json:"type"`
	BankName      string    `json:"bank_name"`
	BankCode      string    `json:"bank_code"`
	CustomerID    string    `json:"customer_id"`
	PointsCard    bool      `json:"points_card"`
}

// CardArgs represents the arguments sent to the Openpay API to create a new
// card.
type CardArgs struct {
	ID              string  `json:"id"`
	HolderName      string  `json:"holder_name"`
	CardNumber      string  `json:"card_number"`
	CVV2            string  `json:"cvv2,omitempty"`
	ExpirationMonth string  `json:"expiration_month"`
	ExpirationYear  string  `json:"expiration_year"`
	Address         Address `json:"address,omitempty"`
}

// CardTokenArgs represents the data sent to the Openpay API to create a new
// card from previously tokenized information.
type CardTokenArgs struct {
	TokenID         string `json:"token_id"`
	DeviceSessionID string `json:"device_session_id"`
}

// AddCard saves a card to your merchant account.
func (m *Merchant) AddCard(args *CardArgs) (*Card, error) {
	return m.client.addCardToResourcePath("cards", args)
}

// AddCardWithToken saves a card to your merchant account, using a card token
// generated with the JavaScript library.
func (m *Merchant) AddCardWithToken(args *CardTokenArgs) (*Card, error) {
	return m.client.addCardToResourcePath("cards", args)
}

// AddCard saves a card to the customer's account.
func (c *Customer) AddCard(args *CardArgs) (*Card, error) {
	if c.ID == "" {
		return nil, ErrNoResourceID
	}
	if c.Merchant == nil {
		return nil, ErrNilMerchant
	}
	path := fmt.Sprintf("customers/%s/cards", c.ID)
	return c.Merchant.client.addCardToResourcePath(path, args)
}

// AddCardWithToken saves a card to the customer's account, using a card token
// generated with the JavaScript library.
func (c *Customer) AddCardWithToken(args *CardTokenArgs) (*Card, error) {
	if c.ID == "" {
		return nil, ErrNoResourceID
	}
	if c.Merchant == nil {
		return nil, ErrNilMerchant
	}
	path := fmt.Sprintf("customers/%s/cards", c.ID)
	return c.Merchant.client.addCardToResourcePath(path, args)
}

func (c *client) addCardToResourcePath(path string, data interface{}) (*Card, error) {
	var card Card
	if err := c.performCardOperation("POST", path, data, &card); err != nil {
		return nil, err
	}
	return &card, nil
}

// GetCard gets a card saved on your merchant account.
func (m *Merchant) GetCard(id string) (*Card, error) {
	return m.client.getCardFromResourcePath("cards", id)
}

// GetCard gets a card saved on the customer's account.
func (c *Customer) GetCard(id string) (*Card, error) {
	if c.ID == "" {
		return nil, ErrNoResourceID
	}
	if c.Merchant == nil {
		return nil, ErrNilMerchant
	}
	path := fmt.Sprintf("customers/%s/cards", c.ID)
	return c.Merchant.client.getCardFromResourcePath(path, id)
}

func (c *client) getCardFromResourcePath(path, id string) (*Card, error) {
	var card Card
	if err := c.performCardOperation("GET", path+"/id", nil, &card); err != nil {
		return nil, err
	}
	return &card, nil
}

// DeleteCard removes a card saved on your merchant account.
func (m *Merchant) DeleteCard(id string) error {
	return m.client.performCardOperation("DELETE", "cards/"+id, nil, nil)
}

// DeleteCard removes a card saved on the customer's account
func (c *Customer) DeleteCard(id string) error {
	path := fmt.Sprintf("customers/%s/cards", c.ID)
	return c.Merchant.client.performCardOperation("DELETE", path, nil, nil)
}

func (c *client) performCardOperation(method, path string, data, dst interface{}) error {
	req, err := c.newRequest(method, path, data)
	if err != nil {
		return err
	}
	if err = c.perform(req, dst); err != nil {
		return err
	}
	return nil
}

// TODO: refactor
