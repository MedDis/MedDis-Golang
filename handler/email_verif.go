package handler

import (
	"encoding/json"
	"gsc_rest/model"
	"gsc_rest/repository"
	"gsc_rest/service"
	"net/http"
	"time"
)

func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var result []byte
	var err error

	if r.Method == "POST" {

		payload := make(M)
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if payload["email"] == nil || payload["code"] == nil {
			result, _ = json.Marshal(model.ResponseError{
				StatusCode: 400,
				Error:      "Bad Request",
				Message:    "Bad Request",
			})

			w.WriteHeader(http.StatusBadRequest)
			w.Write(result)
			return
		}

		var codeVerif int
		var email string

		str, ok := payload["code"].(float64)
		str2, ok2 := payload["email"].(string)

		if !ok && !ok2 {
			result, _ = json.Marshal(model.ResponseError{
				StatusCode: 400,
				Error:      "Bad Request",
				Message:    "Bad Request",
			})

			w.WriteHeader(http.StatusBadRequest)
			w.Write(result)
			return
		} else {
			codeVerif = int(str)
			email = str2
		}

		codeResult, times, _ := repository.SelectCodeVerification(email)

		if !service.IsStartBeforeTimeEndWithDelayed(time.Now().String()[:19], times, (5 * time.Minute)) {
			if codeResult == codeVerif {
				repository.RemoveCodeVerification(email)
				result, _ = json.Marshal(model.ResponseTemplate{
					Error:   false,
					Code:    200,
					Message: "Validasi Kode Berhasil",
					Data: map[string]any{
						"success": true,
					},
				})
			} else {
				result, _ = json.Marshal(model.ResponseTemplate{
					Error:   false,
					Code:    200,
					Message: "Validasi Gagal",
					Data: map[string]any{
						"success": false,
					},
				})
			}
		} else {
			result, err = json.Marshal(model.ResponseTemplate{
				Error:   false,
				Code:    200,
				Message: "Validasi Gagal, code verifikasi telah expire",
				Data: map[string]any{
					"success": false,
				},
			})
		}

		if err != nil {
			result, _ = json.Marshal(model.ResponseError{
				StatusCode: 503,
				Error:      err.Error(),
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
