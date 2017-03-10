package openpay

import (
	"fmt"
	"time"
)

// Subscription represents a customer's subscription to a plan
type Subscription struct {
	ID                  string             `json:"id"`
	CreationDate        time.Time          `json:"creation_date"`
	CancelAtPeriodEnd   bool               `json:"cancel_at_period_end"`
	ChargeDate          time.Time          `json:"charge_date"`
	CurrentPeriodNumber int                `json:"current_period_number"`
	PerdiodEndDate      time.Time          `json:"period_end_date"`
	TrialEndDate        time.Time          `json:"trial_end_date"`
	PlanID              int                `json:"plan_id"`
	Status              SubscriptionStatus `json:"status"`
	CustomerID          int                `json:"customer_id"`
	Card                Card               `json:"card"`
}

// SubscriptionStatus is a subscription's current status
type SubscriptionStatus string

// Available status options
const (
	ActiveStatus    SubscriptionStatus = "active"
	TrialStatus                        = "trial"
	PastDueStatus                      = "past_due"
	UnpaidStatus                       = "unpaid"
	CancelledStatus                    = "cancelled"
)

// SubscriptionArgs is the object sent to the Openpay API when a new
// subscription is created
type SubscriptionArgs struct {
	PlanID       int    `json:"plan_id"`
	TrialEndDate string `json:"trial_end_date"`
	SourceID     string `json:"source_id,omitempty"`
	Card         Card   `json:"card,omitempty"`
}

// AddSubscription subscribes a customer to a plan
func (c *Customer) AddSubscription(args *SubscriptionArgs) (*Subscription, error) {
	var subscription Subscription
	if err := c.performSubOperation("POST", "subscriptions", args, &subscription); err != nil {
		return nil, err
	}
	return &subscription, nil
}

// UpdateSubscription updates a subscription with the provided data
func (c *Customer) UpdateSubscription(id string, data *Subscription) (*Subscription, error) {
	var subscription Subscription
	if err := c.performSubOperation("PUT", "subscriptions/"+id, data, &subscription); err != nil {
		return nil, err
	}
	return &subscription, nil
}

// GetSubscription retrieves a subscription with the specified ID
func (c *Customer) GetSubscription(id string) (*Subscription, error) {
	var subscription Subscription
	if err := c.performSubOperation("GET", "subscriptions/"+id, nil, &subscription); err != nil {
		return nil, err
	}
	return &subscription, nil
}

// GetSubscriptions gets all of a customer's subscriptions
func (c *Customer) GetSubscriptions() ([]Subscription, error) {
	var subscriptions []Subscription
	if err := c.performSubOperation("GET", "subscriptions", nil, &subscriptions); err != nil {
		return nil, err
	}
	return subscriptions, nil
}

// DeleteSubscription cancels a subscription immediately
func (c *Customer) DeleteSubscription(id string) error {
	return c.performSubOperation("DELETE", "subscriptions/"+id, nil, nil)
}

func (c *Customer) performSubOperation(verb, path string, data, dst interface{}) error {
	path = fmt.Sprintf("customers/%s/%s", c.ID, path)
	client := c.Merchant.client
	req, err := client.newRequest(verb, path, data)
	if err != nil {
		return err
	}
	if err = client.perform(req, dst); err != nil {
		return err
	}
	return nil
}

// TODO: refactor
