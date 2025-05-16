package main

// We will build a payment system using interfaces

// Create an interface
type Paymenter interface {
	Pay(amount float64)
	Refund(amount float32, transactionID string, accountNo int32, reasone string)
}

type Payment struct {
	gateway Paymenter
}

type Paypal struct{}

func (p Paypal) Pay(amount float64) {
	// PayPal payment logic
	println("Paying with PayPal: ", amount)
}

func (p Paypal) Refund(amount float32, transactionID string, accountNo int32, reason string) {
	// PayPal refund logic
	println("Refunding with PayPal: ", amount, " Transaction ID: ", transactionID, " Account No: ", accountNo, " Reason: ", reason)
}

type Bkash struct{}

func (b Bkash) Pay(amount float64) {
	// Bkash payment logic
	println("Paying with Bkash: ", amount)
}

func (b Bkash) Refund(amount float32, transactionID string, accountNo int32, reason string) {
	// Bkash refund logic
	println("Refunding with Bkash: ", amount, " Transaction ID: ", transactionID, " Account No: ", accountNo, " Reason: ", reason)
}

func main() {
	// Create a new paypal payment instance
	// paypal := Paypal{}

	// Create a new bkash payment instance
	bkash := Bkash{}

	newTransaction := Payment{
		gateway: bkash,
	}

	// Call the Pay method
	newTransaction.gateway.Pay(100)
	// Call the Refund method
	newTransaction.gateway.Refund(100, "12345", 12345678, "Product not received")

}
