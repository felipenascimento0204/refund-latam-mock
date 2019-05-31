package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/felipenascimento0204/refund-latam-mock/business"

	"github.com/felipenascimento0204/refund-latam-mock/transport"

	"github.com/gorilla/mux"
)

//http://localhost:8080/refund-latam/api/v1/DisplayIdentifyDocumentsRefund/document/12345678900/pnr/ABCDEF
func displayRefundHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	displayRefundResponse := business.DisplayDocumentRefund(vars["document"], vars["pnr"])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(displayRefundResponse); err != nil {
		panic(err)
	}
}

//http://localhost:8080/refund-latam/api/v1/request-refund
func requestRefundHandler(w http.ResponseWriter, r *http.Request) {
	requestRefundResponse := business.RequestRefund(transport.RefundRequest{})
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(requestRefundResponse); err != nil {
		panic(err)
	}
}

func main() {

	port := "8080"

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	fmt.Printf("Listining port... %v", port)

	r := mux.NewRouter()
	s := r.PathPrefix("/refund-latam/api/v1").Subrouter()
	s.HandleFunc("/DisplayIdentifyDocumentsRefund/document/{document}/pnr/{pnr}", displayRefundHandler).Methods("GET")
	s.HandleFunc("/request-refund", requestRefundHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+port, r))

}
