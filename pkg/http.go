package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func MakeRequest(
	ctx context.Context,
	method,
	url string,
	modifyRequest func(*http.Request),
	body interface{},
	response interface{},
	errorResponse interface{},
) error {
	// Convert body to []byte
	var reqBody []byte
	if body != nil {
		var ok bool
		reqBody, ok = body.([]byte)
		if !ok {
			// Convert body to JSON
			jsonBody, err := json.Marshal(body)
			if err != nil {
				return err
			}
			reqBody = jsonBody
		}
	}

	// Create a new request object
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	// Set headers
	if modifyRequest != nil {
		modifyRequest(req)
	}

	client := &http.Client{}
	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()

	}()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// If the response status code is not in the 200-299 range, unmarshal the error onto the provided struct
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err = json.Unmarshal(respBody, &errorResponse)
		return fmt.Errorf(
			"request failed with status %d: %s error: %w",
			resp.StatusCode,
			errorResponse,
			err,
		)
	}

	// Unmarshal the response onto the provided struct
	if response != nil {
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return err
		}
	}
	// Return nil if there were no errors
	return nil
}
