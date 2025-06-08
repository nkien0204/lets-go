package greeting

import (
	"encoding/json"
	"net/http"
)

func (delivery *delivery) Greeting() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		greetingResponse, err := delivery.usecase.Greeting()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(greetingResponse.Msg)
	}
}
