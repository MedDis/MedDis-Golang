package handler

import (
	"encoding/json"
	"gsc_rest/model"
	"gsc_rest/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllGrugsProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var result []byte
	var err error

	if r.Method == "GET" {

		if len(r.URL.Query()["page"]) != 1 {
			result, err = json.Marshal(model.ResponseError{
				StatusCode: 400,
				Error:      "Bad Request",
				Message:    "Parameter id tidak valid",
			})
		} else {
			page, _ := strconv.Atoi(r.URL.Query()["page"][0])
			var dataResult = repository.GetAllDrugProductsData(page)

			result, err = json.Marshal(model.ResponseTemplate{
				Error:   false,
				Code:    200,
				Message: "Get Drugs Product complete",
				Data:    dataResult,
			})
		}

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

func GetAllProductsByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var result []byte
	var err error

	if r.Method == "GET" {

		params := mux.Vars(r)

		id, ok := params["productid"]

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

		drugs, notEmpty := repository.GetDrugProductsDataById(drugsId)

		var drugsList []model.DrugsProduct
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
