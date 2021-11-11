package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
)

type stripePayload struct {
	Amount string `json:"amount"`
}

type jsonResponse struct {
	OK bool `json:"ok"`
	Message string `json:"message,omitempty"`
	Content string `json:"content,omitempty"`
	ID int `json:"id,omitempty"`
}


func (app *application) createCharge(w http.ResponseWriter, r *http.Request){

	var payload stripePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}

	amount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}

	cards := Card {
		Secret: app.config.stripe.secret,
		Key: app.config.stripe.key,
	}

	success := true

	ch, msg, err := cards.CreateCharge(amount)
	if err != nil {
		success = false
	}

	if success {
		out, err := json.MarshalIndent(ch, "", "  ")
		if err != nil {
			app.errorLogger.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	} else {
		j := jsonResponse {
			OK: false,
			Message: msg,
			Content: "",
		}
	
		out, err := json.MarshalIndent(j, "", "  ")
		if err != nil {
			app.errorLogger.Println(err)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	}

}

func (app *application) captureCharge(w http.ResponseWriter, r *http.Request){

	var payload stripePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}

	amount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}

	params := mux.Vars(r) 

	chargeId := params["chargeId"]

	success := true

	cards := Card {
		Secret: app.config.stripe.secret,
		Key: app.config.stripe.key,
	}

	ch, msg, err := cards.CaptureCharge(chargeId, amount)
	if err != nil {
		success = false
	}
	
	if success {
		out, err := json.MarshalIndent(ch, "", "  ")
		if err != nil {
			app.errorLogger.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	} else {
		j := jsonResponse {
			OK: false,
			Message: msg,
			Content: "",
		}
	
		out, err := json.MarshalIndent(j, "", "  ")
		if err != nil {
			app.errorLogger.Println(err)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	}

}

func (app *application) createRefund(w http.ResponseWriter, r *http.Request){

	var payload stripePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}

	amount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}

	params := mux.Vars(r) 

	chargeId := params["chargeId"]

	success := true

	cards := Card {
		Secret: app.config.stripe.secret,
		Key: app.config.stripe.key,
	}

	ch, msg, err := cards.CreateRefund(chargeId, amount)
	if err != nil {
		success = false
	}
	
	if success {
		out, err := json.MarshalIndent(ch, "", "  ")
		if err != nil {
			app.errorLogger.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	} else {
		j := jsonResponse {
			OK: false,
			Message: msg,
			Content: "",
		}
	
		out, err := json.MarshalIndent(j, "", "  ")
		if err != nil {
			app.errorLogger.Println(err)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	}

}

func (app *application) getAllCharges(w http.ResponseWriter, r *http.Request){

	success := true

	cards := Card {
		Secret: app.config.stripe.secret,
		Key: app.config.stripe.key,
	}

	charges, msg, err := cards.GetAllCharges()
	if err != nil {
		success = false
	}
	
	if success {
		out, err := json.MarshalIndent(charges, "", "  ")
		if err != nil {
			app.errorLogger.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	} else {
		j := jsonResponse {
			OK: false,
			Message: msg,
			Content: "",
		}
	
		out, err := json.MarshalIndent(j, "", "  ")
		if err != nil {
			app.errorLogger.Println(err)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	}

}

func (app *application) routes() http.Handler {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	
	router.HandleFunc("/api/v1/create_charge", app.createCharge).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/capture_charge/{chargeId}", app.captureCharge).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/create_refund/{chargeId}", app.createRefund).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/get_charges", app.getAllCharges).Methods(http.MethodGet)

	return router
}
