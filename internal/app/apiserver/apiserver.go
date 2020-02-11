package apiserver

import (
"fmt"
"net/http"
"encoding/json"
	"log"
	// "strconv"

	"github.com/gorilla/mux"
)

type ReqCount struct {
	Offer_Id int `json:"offer_id"`
	Land_Url string `json:"land_url"`
	Count_Id int `json:"count_Id"`
}

// GetCount ...
type GetCount struct {
	Offer_Id int `json:"offer_id"`
	Offer_Name string `json:"offer_name"`
	Land_Url string `json:"land_url"`
}

func router()  {
	router := mux.NewRouter()
	router.HandleFunc("/api", GetApi).Methods("POST")

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8081", nil)
}

func GetApi(w http.ResponseWriter, r *http.Request)  {
	getCount := GetCount{}

	err := json.NewDecoder(r.Body).Decode(&getCount)
	log.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"`+err.Error()+`"`))
		return
	}

	//labelName := getCount.Offer_Name + " " + strconv.Itoa(getCount.Offer_Id)

	reqest := ReqCount{}
	reqest.Land_Url = getCount.Land_Url
	reqest.Offer_Id = getCount.Offer_Id
	reqest.Count_Id = 555

	payload, err := json.Marshal(&reqest)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}