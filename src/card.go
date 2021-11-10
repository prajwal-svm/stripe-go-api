package main

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/stripe/stripe-go/v72/charge"
	"github.com/stripe/stripe-go/v72/paymentmethod"
)

type Card struct {
	Secret string
	Key string
	Currency string
}

type Transaction struct {
	TransactionStatusId int
	Amount int
	Currency string
	BankResponseCode string
}

func (c *Card) GetPaymentMethodId(number string, month string, year string, cvc string) (*stripe.PaymentMethod, string, error) {
	stripe.Key = c.Secret

	methodParams := &stripe.PaymentMethodParams{
		Card: &stripe.PaymentMethodCardParams{
		  Number: stripe.String(number),
		  ExpMonth: stripe.String(month),
		  ExpYear: stripe.String(year),
		  CVC: stripe.String(cvc),
		},
		Type: stripe.String("card"),
	  }
	  pm, err := paymentmethod.New(methodParams)

	  if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = string(stripeErr.Code)
		}
		return nil, msg, err
	}

	return pm, "", nil
}


func (c *Card) CreateCharge(currency string, amount int, paymentMethodId string, customerId string) (*stripe.PaymentIntent, string, error) {
	stripe.Key = c.Secret

	// _params := &stripe.PaymentMethodAttachParams{
	// 	Customer: stripe.String(customerId),
	// }
	
	// paymentmethod.Attach(
	// 	paymentMethodId,
	// 	_params,
	// )

	params := &stripe.PaymentIntentParams {
		Amount: stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
		PaymentMethod: stripe.String(paymentMethodId),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		  }),
		CaptureMethod: stripe.String(string(stripe.PaymentIntentCaptureMethodManual)),
		Confirm: stripe.Bool(true),
		
	}

	pi, err := paymentintent.New(params)

	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = string(stripeErr.Code)
		}
		return nil, msg, err
	}

	return pi, "", nil
}

func (c *Card) CaptureCharge(chargeId string, amount int) (*stripe.PaymentIntent, string, error) {
	stripe.Key = c.Secret

	params := &stripe.PaymentIntentCaptureParams{
		AmountToCapture: stripe.Int64(int64(amount)),
	  }

	paymentintent.Confirm(chargeId, nil)

	pi, err := paymentintent.Capture(chargeId, params)

	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = string(stripeErr.Code)
		}
		return nil, msg, err
	}

	return pi, "", nil
}

func (c *Card) GetAllCharges() ([]*stripe.Charge, string, error) {
	stripe.Key = c.Secret

	params := &stripe.ChargeListParams{}
	params.Filters.AddFilter("limit", "", "100")
	i := charge.List(params)

	a := []*stripe.Charge{}
	for i.Next() {
		a = append(a, i.Charge())
	}

	return a, "", nil
}