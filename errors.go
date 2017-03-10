package openpay

import "fmt"

// APIError represents an error returned by the Openpay API
type APIError struct {
	Category    string   `json:"category"`
	ErrorCode   int      `json:"error_code"`
	Description string   `json:"description"`
	HTTPCode    int      `json:"http_code"`
	RequestID   string   `json:"request_id"`
	FraudRules  []string `json:"fraud_rules"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	return fmt.Sprintf("http status code: %d; error %d: %s",
		e.HTTPCode, e.ErrorCode, e.Description)
}
