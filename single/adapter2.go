package main

import (
	"fmt"
	"errors"
)

type Paypal struct {}
func (self *Paypal) sendMoney(from,to string, amount float64) error {
	fmt.Println("PAYPAL: " + from + " " + to)
	if(from == "") {
		return errors.New("dari paypal")
	}
	return nil
}

type Bank struct {}
func (self *Bank) delivery(from,to string, amount float64) error {
	fmt.Println("BANK: " + from + " " + to)
	return nil
}

/* Adapter */

type PaypalAdapter struct {
	o *Paypal
}

func (self *PaypalAdapter) Pay(from,to string,amount float64) error {
	self.o.sendMoney(from,to,amount)
	return nil
}

type BankAdapter struct {
	o *Bank
}

func (self *BankAdapter) Pay(from,to string,amount float64) error {
	self.o.delivery(from,to,amount)
	return nil
}


/* Requester */
type Payment interface {
	Pay(from,to string,amount float64) error
}

type Shopping struct {
	PaymentMethod Payment
}

func (self *Shopping) Checkout(from,to string, amount float64) error {
	self.PaymentMethod.Pay(from,to,amount)
	return nil
} 

func main() {
	var s Shopping
	s.PaymentMethod = &PaypalAdapter{};
	s.Checkout("kambing","gunung",20000);
}