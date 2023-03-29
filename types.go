package main

type PaymentRequest struct {
	Amount      int32  `json:"amount"`
	PayerID     string `json:"payerIdentity"`
	PayerName   string `json:"payerName"`
	Description string `json:"description"`
	ReturnURL   string `json:"returnUrl"`
	ClientRefID string `json:"clientRefId"`
}

type PaypingRequest[T PaymentRequest | VerifyPaymentRequest] struct {
	Url   string
	Token string
	Body  T
}

type PaymentResponse struct {
	Code string `json:"code"`
}

type InvoiceResponse struct {
	Code  string
	ClientRefID string
	Link  string
}

type VerifyPaymentRequest struct {
	Amount      int32  `json:"amount"`
	RefID string `json:"RefId"`
}

type VerifyPaymentResponse struct {
	Amount      int32  `json:"amount"`
	CardNumber  string `json:"cardNumber"`
	CardHashPan string `json:"cardHashPan"`
}
