
# Payping 

Payping gateway in Golang 



## Methods

- [Create Invoice](https://github.com/rootiens/Payping#create-invoice)
- [Verify Payment](https://github.com/rootiens/Payping#verify-payment)


## Installation


```bash
go get -u github.com/rootiens/payping
```
    

## Usage/Examples

### Create Invoice

```go
package main

import (
    "fmt"
    "github.com/rootiens/payping"
)

func main() {
	token := "payping token"
	payment := PaymentRequest{
		Amount:      100, // Int32
		PayerID:     "09123456789",
		PayerName:   "Rootiens",
		Description: "Desc",
		ReturnURL:   "https://example.com",
		ClientRefID: "", // Package automatically generates an UUID in case you don't specify any ClientRefID or Its length is < 6
	}

	response, err := CreateInvoice(payment, token)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.ClientRefID) // ClientRefID
	fmt.Println(response.Link)        // Gateway link for redirection
	fmt.Println(response.Code)        // Gateway payment code
}
```

### Verify Payment

```go
package main

import (
     "fmt"
     "github.com/rootiens/payping"
)

func main() {
	verify := VerifyPaymentRequest{
		Amount: 100, // Int32
		RefID:  refID, // After going back to ReturnURL, Payping gives RefID to verify payment 
	}

	responseVer, err := VerifyPayment(verify, token)
	if err != nil {
		panic(err)
	}

	fmt.Println(responseVer.Amount)
	fmt.Println(responseVer.CardNumber)
	fmt.Println(responseVer.CardHashPan)
}
```




