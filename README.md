# stripe-go-api

Backend APIs for **Stripe Payment Gateway** Integration implemented in **GoLang**

## Supported APIs

```
## Note: Supports only 4 simple APIs which are not production ready!

## Charge Currency: INR

# Create Charge for CC payment
POST /api/v1/create_charge

# Capture the created charge
POST /api/v1/capture_charge/:chargeId

# Create Refund for the created charge
POST /api/v1/create_refund/:chargeId

# Get a list of all charges
GET /api/v1/get_charges
```

## Requirements

- OS : Linux or Mac or Windows with WSL2
- Go: V1.17+
- Make
- Stripe Secret key & Publishable key (Get your keys for free from https://dashboard.stripe.com/register)
- Postman for API testing (Could also use curl for simplicity)

## Installation

1. Clone the Repo   
``` 
→ git clone https://github.com/Prajwal-S-Venkatesh/stripe-go-api.git
```

2. Add the Publishable & Secret keys in the Makefile
3. Start the go server by running
```
→ make serve
```


## API Testing 

Start the Postman application and import the `stripe-go-api-collection.postman_collection.json` file. 

1. Trigger **Create Charge** API with required amount and copy the charge id from the returned json data which resembles `ch_3JugY3SIJJvsfEwy1ZRBRjdr` pattern.
2. Paste the Charge Id in the **Capture Charge** API as URI params and trigger it to capture the charge that was previously created with required amount.
3. Paste the same Charge Id in the **Refund Charge** API as URI params and trigger it to refund the the captured charge with required amount(derfaults to charged amount).
4. Trigger **Get All Charges** API to get a list of all charges and verify their status (amount_captured, amount_refunded, captured, id, status...)