package main

import (
	"encoding/json"
	"errors"
	"io"
)

const (
	CreateInvoiceURL = "https://api.payping.ir/v2/pay"
	PurchaseURL      = "https://api.payping.ir/v2/pay/gotoipg/"
	VerifyURL        = "https://api.payping.ir/v2/pay/verify"
)

func CreateInvoice(payment PaymentRequest, token string) (InvoiceResponse, error) {

	if urlok := ReturnURLValidation(payment.ReturnURL); !urlok {
		return InvoiceResponse{}, errors.New("Invalid ReturnURL")
	}

    if payment.PayerName == "" {
        return InvoiceResponse{}, errors.New("PayerName shouldn't be empty")
    }

    if payment.PayerID == "" {
        return InvoiceResponse{}, errors.New("PayerID shouldn't be empty")
    }

    if payment.Description == "" {
        return InvoiceResponse{}, errors.New("Description shouldn't be empty")
    }

	payment.ClientRefID = RefIDValidation(payment.ClientRefID)

	payload := &PaypingRequest[PaymentRequest]{
		Url:   CreateInvoiceURL,
		Token: token,
		Body:  payment,
	}

	response, err := PostRequest(payload)

	defer response.Body.Close()

	if err != nil {
		return InvoiceResponse{}, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return InvoiceResponse{}, err
	}

	if response.StatusCode != 200 {
		var response map[string]string

		if err := json.Unmarshal(body, &response); err != nil {
			return InvoiceResponse{}, err
		}

		for key, value := range response {
			return InvoiceResponse{}, errors.New(key + ": " + value)
		}
	}

	var payRes PaymentResponse
	if err := json.Unmarshal(body, &payRes); err != nil {
		return InvoiceResponse{}, err
	}

	Invoice := &InvoiceResponse{
		Code:        payRes.Code,
		ClientRefID: payment.ClientRefID,
		Link:        PurchaseURL + payRes.Code,
	}

	return *Invoice, nil
}

func VerifyPayment(req VerifyPaymentRequest, token string) (VerifyPaymentResponse, error) {
	payload := &PaypingRequest[VerifyPaymentRequest]{
		Url:   VerifyURL,
		Token: token,
		Body:  req,
	}
	response, err := PostRequest(payload)

	defer response.Body.Close()

	if err != nil {
		return VerifyPaymentResponse{}, err
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return VerifyPaymentResponse{}, err
	}

	if response.StatusCode != 200 {
		var response map[string]string

		if err := json.Unmarshal(body, &response); err != nil {
			return VerifyPaymentResponse{}, err
		}

		for key, value := range response {
			return VerifyPaymentResponse{}, errors.New(key + ": " + value)
		}
	}

	var verifyRes VerifyPaymentResponse
	if err := json.Unmarshal(body, &verifyRes); err != nil {
		return VerifyPaymentResponse{}, err
	}

	return verifyRes, nil
}
