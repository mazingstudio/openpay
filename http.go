package openpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

func (c *client) newRequest(method, resource string, data interface{}) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s/%s", c.apiBase, c.merchantID, resource)

	var body bytes.Buffer
	enc := json.NewEncoder(&body)
	if err := enc.Encode(data); err != nil {
		return nil, errors.Wrap(err, "error encoding JSON")
	}

	req, err := http.NewRequest(method, url, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.privateKey, "")

	return req, nil
}

func (c *client) perform(req *http.Request, dst interface{}) error {
	res, err := c.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "error performing request")
	}
	dec := json.NewDecoder(res.Body)
	if res.StatusCode >= 400 {
		var apiErr APIError
		if err = dec.Decode(&apiErr); err != nil && err != io.EOF {
			return errors.Wrap(err, "error decoding api error from JSON")
		}
		apiErr.HTTPCode = res.StatusCode
		return &apiErr
	}
	if dst != nil {
		if err = dec.Decode(&dst); err != nil {
			return errors.Wrap(err, "error decoding JSON")
		}
	}
	return nil
}
