package main

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
	"github.com/stripe/stripe-go/v72/refund"
)

type Card struct {
	Secret string
	Key string
}

func (c *Card) CreateCharge(amount int) (*stripe.Charge, string, error) {
	stripe.Key = c.Secret

	params := &stripe.ChargeParams{
		Amount: stripe.Int64(int64(amount)),
		Currency: stripe.String(string(stripe.CurrencyINR)),
		Source: &stripe.SourceParams{Token: stripe.String("tok_visa")},
		Capture: stripe.Bool(false),
	  }

	ch, err := charge.New(params)

	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = string(stripeErr.Code)
		}
		return nil, msg, err
	}

	return ch, "", nil
}

func (c *Card) CaptureCharge(chargeId string, amount int) (*stripe.Charge, string, error) {
	stripe.Key = c.Secret

	params := &stripe.CaptureParams{
		Amount: stripe.Int64(int64(amount)),
	  }

	ch, err := charge.Capture(chargeId, params)

	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = string(stripeErr.Code)
		}
		return nil, msg, err
	}

	return ch, "", nil
}

func (c *Card) CreateRefund(chargeId string, amount int) (*stripe.Refund, string, error) {
	stripe.Key = c.Secret

	params := &stripe.RefundParams{
		Charge: stripe.String(chargeId),
		Amount: stripe.Int64(int64(amount)),
	  }

	rf, err := refund.New(params)

	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = string(stripeErr.Code)
		}
		return nil, msg, err
	}

	return rf, "", nil
}

func (c *Card) GetAllCharges() ([]*stripe.Charge, string, error) {
	stripe.Key = c.Secret

	params := &stripe.ChargeListParams{}
	params.Filters.AddFilter("limit", "", "100")
	i := charge.List(params)

	ch := []*stripe.Charge{}
	for i.Next() {
		ch = append(ch, i.Charge())
	}

	return ch, "", nil
}