package main

// type student struct {
// 	ID    string
// 	Name  string
// 	Grade int
// }

// var data = []student{
// 	{"E001", "ethan", 21},
// 	{"W001", "wick", 22},
// 	{"B001", "bourne", 23},
// 	{"B002", "bond", 23},
// 	{"B002", "test", 24},
// }

// func users(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == "GET" {
// 		var result, err = json.Marshal(data)

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Write(result)
// 		return
// 	}

// 	http.Error(w, "", http.StatusBadRequest)
// }

// func test(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var result []byte
// 	var err error

// 	if r.Method == "GET" {

// 		var buffer = responseTemplate{
// 			Error:   false,
// 			Code:    200,
// 			Message: "Sukses",
// 			Data:    data,
// 		}

// 		result, err = json.Marshal(buffer)

// 		fmt.Println(result)
// 		fmt.Println(buffer)

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		w.Write(result)
// 		return
// 	}

// 	http.Error(w, "", http.StatusBadRequest)
// }

// func user(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == "GET" {
// 		var id = r.FormValue("id")
// 		var result []byte
// 		var err error

// 		for _, each := range data {
// 			if each.ID == id {
// 				result, err = json.Marshal(each)
// 				fmt.Println(each)

// 				if err != nil {
// 					http.Error(w, err.Error(), http.StatusInternalServerError)
// 					return
// 				}

// 				w.Write(result)
// 				return
// 			}
// 		}

// 		http.Error(w, "User not found", http.StatusNotFound)
// 		return
// 	}

// 	http.Error(w, "", http.StatusBadRequest)
// }
