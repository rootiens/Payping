package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func PostRequest[T PaymentRequest | VerifyPaymentRequest](payload *PaypingRequest[T]) (*http.Response, error) {
	body, err := json.Marshal(payload.Body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", payload.Url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+payload.Token)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

