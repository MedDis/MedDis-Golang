package handler

import (
	"encoding/json"
	"gsc_rest/model"
	"gsc_rest/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllDrugs(w http.ResponseWriter, r *http.Request) {
	var result []byte
	var err error

	if r.Method == "GET" {

		drugs := repository.GetAllDrugsData()

		result, err = json.Marshal(model.ResponseTemplate{
			Error:   false,
			Code:    200,
			Message: "Berhasil Mendapatkan data",
			Data:    drugs,
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

func GetAllDrugsByID(w http.ResponseWriter, r *http.Request) {
	var result []byte
	var err error

	if r.Method == "GET" {

		params := mux.Vars(r)

		id, ok := params["drugsid"]

		if !ok {
			result, _ = json.Marshal(model.ResponseError{
				StatusCode: 400,
				Error:      "Bad Request",
				Message:    "Parameter id tidak valid",
			})

			w.WriteHeader(http.StatusBadRequest)

			w.Write(result)
			return
		}

		drugsId, isNum := strconv.Atoi(id)

		if isNum != nil {
			result, _ = json.Marshal(model.ResponseError{
				StatusCode: 400,
				Error:      "Bad Request",
				Message:    "Parameter id tidak valid",
			})

			w.WriteHeader(http.StatusBadRequest)

			w.Write(result)
			return
		}

		drugs, notEmpty := repository.GetDrugsByID(drugsId)

		var drugsList []model.DrugsComposition
		if notEmpty {
			drugsList = append(drugsList, drugs)
		}

		result, err = json.Marshal(model.ResponseTemplate{
			Error:   false,
			Code:    200,
			Message: "Berhasil Mendapatkan Data",
			Data:    drugsList,
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
