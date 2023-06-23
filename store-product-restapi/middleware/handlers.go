package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"store-product-restapi/controllers"
	"store-product-restapi/models"
	_ "github.com/lib/pq"
)

type response struct {
	ID      string 	`json:"id,omitempty"`
	Message string  `json:"message,omitempty"`
}

func CreateStore(w http.ResponseWriter, r *http.Request) {
	var store models.Store
	err := json.NewDecoder(r.Body).Decode(&store)

	if err != nil {
		log.Fatalf("Unable to decode the http request body.(JSON) %v", err)
	}

	addID := controllers.Addstore(store)

	res := response{
		ID:      addID,
		Message: "Store created successfully",
	}

	json.NewEncoder(w).Encode(res)

}

func GetStores(w http.ResponseWriter, r *http.Request) {

}

func GetStoreById(w http.ResponseWriter, r *http.Request) {

}

func GetStoreByName(w http.ResponseWriter, r *http.Request) {

}

func UpdateStore(w http.ResponseWriter, r *http.Request) {

}

func DeleteStore(w http.ResponseWriter, r *http.Request) {

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	var metadata models.Metadata
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		log.Fatalf("Unable to decode the http request body.(JSON) %v", err)
	}

	addID := controllers.Addproduct(product, metadata)

	res := response{
		ID:      addID,
		Message: "Product created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {

}
func GetProductById(w http.ResponseWriter, r *http.Request) {

}
func GetProductByName(w http.ResponseWriter, r *http.Request) {

}
func GetProductByDate(w http.ResponseWriter, r *http.Request) {

}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {

}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {

}
