package handler

import (
	"encoding/json"
	"gsc_rest/model"
	"net/http"

	"github.com/gorilla/mux"
)

type M map[string]interface{}

func RouteInit() {
	r := mux.NewRouter()
	r.HandleFunc("/", basePath)
	r.HandleFunc("/email/request/verify/", VerifyEmail)
	r.HandleFunc("/email/request/send/", SendEmailCode)
	r.HandleFunc("/drugs/all/products", GetAllGrugsProduct)
	r.HandleFunc("/drugs/{productid}/products", GetAllProductsByID)
	r.HandleFunc("/drugs/all/composition", GetAllDrugs)
	r.HandleFunc("/drugs/{drugsid}/composition", GetAllDrugsByID)
	http.Handle("/", r)
}

func basePath(w http.ResponseWriter, r *http.Request) {
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

// documentas
/*
   url/{userid}/path
params := mux.Vars(r)

id, ok := params["userid"]

if !ok {
	fmt.Println("id is missing in parameters")
}

body request
payload := make(M)
err2 := json.NewDecoder(r.Body).Decode(&payload)
if err2 != nil {
	http.Error(w, err2.Error(), http.StatusBadRequest)
	return
}

query ?key=data
fmt.Println(r.URL.Query())

*/

//Template response
/*
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


*/
