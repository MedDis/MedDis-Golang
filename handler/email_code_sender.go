package handler

import (
	"encoding/json"
	"fmt"
	"gsc_rest/model"
	"gsc_rest/repository"
	"gsc_rest/service"
	"strconv"
	"time"

	// "gsc_rest/service"
	"net/http"
)

func SendEmailCode(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var result []byte
	var err error

	if r.Method == "POST" {

		payload := make(M)
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if payload["email"] == nil {
			result, _ = json.Marshal(model.ResponseError{
				StatusCode: 400,
				Error:      "Bad Request",
				Message:    "Bad Request",
			})

			w.WriteHeader(http.StatusBadRequest)
			w.Write(result)
			return
		}

		var email string

		if str, ok := payload["email"].(string); !ok {
			result, _ = json.Marshal(model.ResponseError{
				StatusCode: 400,
				Error:      "Bad Request",
				Message:    "Bad Request",
			})

			w.WriteHeader(http.StatusBadRequest)
			w.Write(result)
			return
		} else {
			email = str
		}

		var codeInt int

		code, times := repository.AddNewCodeVerification(email)
		if service.IsStartBeforeTimeEndWithDelayed(time.Now().String()[:19], times, (5 * time.Minute)) {
			if codeInt, error := repository.UpdateCodeVerification(email); !error {
				service.SendCodeVerification(email, codeInt)
			} else {
				fmt.Println(error)
			}
		} else {
			codeInt, _ = strconv.Atoi(code)
			service.SendCodeVerification(email, codeInt)
		}

		result, err = json.Marshal(model.ResponseTemplate{
			Error:   false,
			Code:    200,
			Message: "Email Terkirim, Harap Cek Kotak Masuk Atau spam",
		})

		if err != nil {
			result, _ = json.Marshal(model.ResponseError{
				StatusCode: 503,
				Error:      "Internal Server Error",
				Message:    "Internal Server Error",
			})

			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(result)
		return
	}

	result, _ = json.Marshal(model.ResponseError{
		StatusCode: 404,
		Error:      "Not Found",
		Message:    "Not Found",
	})

	w.WriteHeader(http.StatusNotFound)
	w.Write(result)

}

func SendEmailLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var result []byte
	var err error

	if r.Method == "GET" {

		result, err = json.Marshal(model.ResponseTemplate{
			Error:   false,
			Code:    200,
			Message: "Sukses Terhubung dengan Rest API",
		})

		if err != nil {
			result, _ = json.Marshal(model.ResponseError{
				StatusCode: 503,
				Error:      "Internal Server Error",
				Message:    "Internal Server Error",
			})

			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(result)
		return
	}

	result, _ = json.Marshal(model.ResponseError{
		StatusCode: 404,
		Error:      "Not Found",
		Message:    "Not Found",
	})

	w.WriteHeader(http.StatusNotFound)
	w.Write(result)
}
