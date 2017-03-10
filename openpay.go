package openpay

import "errors"

// standard errors
var (
	ErrNilClient    = errors.New("nil client")
	ErrNilMerchant  = errors.New("this resource doesn't point to a merchant")
	ErrNoResourceID = errors.New("this resource has no ID")
)
