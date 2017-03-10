package openpay

import (
	"fmt"
	"time"
)

// Customer is an Openpay customer.
type Customer struct {
	ID           string         `json:"id,omitempty"`
	CreationDate time.Time      `json:"creation_date"`
	Name         string         `json:"name"`
	LastName     string         `json:"last_name"`
	Email        string         `json:"email"`
	PhoneNumber  string         `json:"phone_number"`
	Status       string         `json:"status"`
	Balance      float32        `json:"balance"`
	CLABE        string         `json:"clabe"`
	Address      Address        `json:"address"`
	Store        StoreReference `json:"store"`

	// Merchant will be set automatically if obtained from through an API call.
	// Otherwise, you must set it yourself.
	Merchant *Merchant
}

// Address represents a customer's address.
type Address struct {
	Line1       string `json:"line1"`
	Line2       string `json:"line2"`
	Line3       string `json:"line3"`
	PostalCode  string `json:"postal_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	CountryCode string `json:"country_code"`
}

// StoreReference represents a customer's store reference.
type StoreReference struct {
	Reference        string `json:"reference"`
	BarcodeURL       string `json:"barcode_url"`
	PaybinReference  string `json:"paybin_reference"`
	BarcodePaybinURL string `json:"barcode_paybin_url"`
}

// CustomerArgs is the object sent to the Openpay API when a new customer is
// created.
type CustomerArgs struct {
	ExternalID      string  `json:"external_id,omitempty"`
	Name            string  `json:"name"`
	LastName        string  `json:"last_name,omitempty"`
	Email           string  `json:"email"`
	RequiresAccount bool    `json:"requires_acount"`
	PhoneNumber     string  `json:"phone_number,omitempty"`
	Address         Address `json:"address,omitempty"`
}

// AddCustomer creates a new customer on the Openpay API.
func (m *Merchant) AddCustomer(args *CustomerArgs) (*Customer, error) {
	req, err := m.client.newRequest("POST", "customers", &args)
	if err != nil {
		return nil, err
	}
	var customer Customer
	err = m.client.perform(req, &customer)
	if err != nil {
		return nil, err
	}
	customer.Merchant = m
	return &customer, nil
}

// GetCustomers lists all available customers.
func (m *Merchant) GetCustomers() ([]Customer, error) {
	req, err := m.client.newRequest("GET", "customers", nil)
	if err != nil {
		return nil, err
	}
	var customers []Customer
	if err = m.client.perform(req, &customers); err != nil {
		return nil, err
	}
	for i := range customers {
		customers[i].Merchant = m
	}
	return customers, nil
}

// GetCustomer gets an Openpay customer.
func (m *Merchant) GetCustomer(id string) (*Customer, error) {
	var customer Customer
	if err := m.performCustomerOperation("GET", id, nil, &customer); err != nil {
		return nil, err
	}
	customer.Merchant = m
	return &customer, nil
}

// UpdateCustomer updates an existing Openpay customer.
func (m *Merchant) UpdateCustomer(id string, data *Customer) (*Customer, error) {
	var customer Customer
	if err := m.performCustomerOperation("PUT", id, data, &customer); err != nil {
		return nil, err
	}
	customer.Merchant = m
	return &customer, nil
}

// DeleteCustomer deletes an Openpay customer.
func (m *Merchant) DeleteCustomer(id string) error {
	return m.performCustomerOperation("DELETE", id, nil, nil)
}

func (m *Merchant) performCustomerOperation(verb, id string, data, dst interface{}) error {
	client := m.client
	req, err := client.newRequest(verb, fmt.Sprintf("customers/%s", id), data)
	if err != nil {
		return err
	}
	if err = client.perform(req, dst); err != nil {
		return err
	}
	return nil
}
